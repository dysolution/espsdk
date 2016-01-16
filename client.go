package espsdk

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/Sirupsen/logrus"
)

var pool *x509.CertPool

// Serializable objects can be Marshaled into JSON.
type Serializable interface {
	Marshal() ([]byte, error)
}

// Findable objects can report the URL where they can be found.
type Findable interface {
	Path() string
}

// A RESTObject has a canonical API endpoint URL and can be serialized to JSON.
type RESTObject interface {
	Serializable
	Findable
}

// GetClient returns a Client that can be used to send requests to the ESP API.
func GetClient(key, secret, username, password, uploadBucket string) *Client {
	creds := credentials{
		APIKey:      key,
		APISecret:   secret,
		ESPUsername: username,
		ESPPassword: password,
	}
	token := getToken(&creds)
	return &Client{creds, token, uploadBucket}
}

// A Client is able to request an access token and submit HTTP requests to
// the ESP API.
type Client struct {
	credentials
	Token        Token
	UploadBucket string
}

// getToken submits the provided credentials to Getty's OAuth2 endpoint
// and returns a token that can be used to authenticate HTTP requests to the
// ESP API.
func getToken(credentials *credentials) Token {
	if credentials.areInvalid() {
		log.Fatal("Not all required credentials were supplied.")
	}

	uri := oauthEndpoint
	log.Debugf("%s", uri)
	formValues := formValues(credentials)
	log.Debugf("%s", formValues.Encode())

	resp, err := http.PostForm(uri, formValues)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	log.Debugf("HTTP %d", resp.StatusCode)
	log.Debugf("%s", payload)
	return tokenFrom(payload)
}

// GetKeywords requests suggestions from the Getty controlled vocabulary
// for the keywords provided.
//
// TODO: not implemented (keywords and personalities need a new struct type)
func (c *Client) GetKeywords() []byte { return c.get(Keywords) }

// GetPersonalities requests suggestions from the Getty controlled vocabulary
// for the famous personalities provided.
//
// TODO: not implemented (keywords and personalities need a new struct type)
func (c *Client) GetPersonalities() []byte { return c.get(Personalities) }

// GetControlledValues returns complete lists of values and descriptions for
// fields with controlled vocabularies, grouped by submission type.
//
// TODO: not implemented (needs new struct type)
func (c *Client) GetControlledValues() []byte { return c.get(ControlledValues) }

// GetTranscoderMappings lists acceptable transcoder mapping values
// for Getty and iStock video.
func (c *Client) GetTranscoderMappings() TranscoderMappingList {
	return TranscoderMappingList{}.Unmarshal(c.get(TranscoderMappings))
}

// GetTermList lists all possible values for the given controlled vocabulary.
func (c *Client) GetTermList(endpoint string) TermList {
	return TermList{}.Unmarshal(c.get(endpoint))
}

// DeleteLastBatch looks up the newest Batch and deletes it.
func (c *Client) DeleteLastBatch() (Result, error) {
	lastBatch := c.Index(Batches).Last()
	return c.verboseDelete(lastBatch.Path())

}

// Index requests a list of all Batches owned by the user.
func (c *Client) Index(path string) *DeserializedObject {
	var obj *DeserializedObject
	return Deserialize(c.get(path), obj)
}

// VerboseCreate uses the provided metadata to create and object
// and returns it along with metadata about the HTTP request, including
// response time.
func (c *Client) VerboseCreate(object Findable) (Result, error) {
	result, err := c.verbosePost(object)
	if err != nil {
		log.Errorf("Client.VerboseCreate: %v", err)
		return Result{}, err
	}
	return result, nil
}

// Update changes metadata for an existing Batch.
func (c *Client) Update(object RESTObject) DeserializedObject {
	return Unmarshal(c.put(object))
}

// VerboseUpdate uses the provided metadata to update an object and returns
// metadata about the HTTP request, including response time.
func (c *Client) VerboseUpdate(object Findable) (Result, error) {
	result, err := c.verbosePut(object)
	if err != nil {
		log.Errorf("Client.VerboseUpdate: %v", err)
		return Result{}, err
	}
	return result, nil
}

// Delete destroys the object at the provided path.
func (c *Client) Delete(path string) DeserializedObject {
	bytes := c._delete(path)
	if len(bytes) > 0 {
		return Unmarshal(bytes)
	}
	// successful deletion usually returns a 204 without a payload/body
	return DeserializedObject{}
}

// VerboseDelete destroys the object described by the provided object,
// as long as enough data is provided to unambiguously identify it to the API.
func (c *Client) VerboseDelete(object Findable) (Result, error) {
	result, err := c.verboseDelete(object.Path())
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Errorf("Client.VerboseDelete: %v", err)
		return Result{}, err
	}
	return result, nil
}

func (c *Client) verboseDelete(path string) (Result, error) {
	result, err := c.performRequest(newRequest("DELETE", path, c.Token, nil))
	if err != nil {
		return Result{}, err
	}
	return result, nil
}

// DeleteFromObject destroys the object described by the provided object,
// as long as enough data is provided to unambiguously identify it to the API.
func (c *Client) DeleteFromObject(object RESTObject) DeserializedObject {
	bytes := c._delete(object.Path())
	if len(bytes) > 0 {
		return Unmarshal(bytes)
	}
	// successful deletion usually returns a 204 without a payload/body
	return DeserializedObject{}
}

// Get requests the metadata for the object at the provided path.
func (c *Client) DeprecatedGet(path string) DeserializedObject {
	pc, _, _, _ := runtime.Caller(0)
	callerPC, _, _, _ := runtime.Caller(1)
	caller := runtime.FuncForPC(callerPC).Name()
	myself := runtime.FuncForPC(pc).Name()
	logPrefix := fmt.Sprintf("%v %v)", caller, myself)
	result, err := c.verboseGet(path)
	if err != nil {
		result.Log().Error("Client.Get")
	}
	log.WithFields(result.Stats()).Info(logPrefix)
	return Unmarshal(result.VerboseResult.Payload)
}

// GetFromObject requests the metadata for the provided object, as long as
// enough data is provided to unambiguously identify it to the API.
func (c *Client) GetFromObject(object RESTObject) DeserializedObject {
	return Unmarshal(c.get(object.Path()))
}

// VerboseGet uses the provided metadata to request an object from the API
// and returns it along with metadata about the HTTP request, including
// response time.
func (c *Client) VerboseGet(object Findable) (Result, error) {
	result, err := c.verboseGet(object.Path())
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Errorf("Client.VerboseGet: %v", err)
		return Result{}, err
	}
	return result, nil
}

func (c *Client) verboseGet(path string) (Result, error) {
	result, err := c.performRequest(newRequest("GET", path, c.Token, nil))
	if err != nil {
		return Result{}, err
	}
	return result, nil
}

func (c *Client) get(path string) []byte {
	request := newRequest("GET", path, c.Token, nil)
	result, err := c.performRequest(request)
	if err != nil {
		log.Fatal(err)
	}
	result.Log().Debug("Client.get")
	log.Debugf("%s\n", result.Payload)
	return result.Payload
}

func (c *Client) verbosePost(object Findable) (Result, error) {
	serializedObject, err := Marshal(object)
	if err != nil {
		return Result{}, err
	}

	request := newRequest("POST", object.Path(), c.Token, serializedObject)
	result, err := c.performRequest(request)
	if err != nil {
		return Result{}, err
	}
	result.Log().Debug("Client.verbosePost")
	return result, nil
}

func (c *Client) post(object RESTObject) []byte {
	serializedObject, err := Marshal(object)
	if err != nil {
		log.Fatal(err)
	}
	request := newRequest("POST", object.Path(), c.Token, serializedObject)
	result, err := c.performRequest(request)
	if err != nil {
		log.Fatal(err)
	}

	result.Log().Debug()
	log.Debugf("%s\n", result.Payload)
	return result.Payload
}

func (c *Client) put(object RESTObject) []byte {
	serializedObject, err := object.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	request := newRequest("PUT", object.Path(), c.Token, serializedObject)
	result, err := c.performRequest(request)
	if err != nil {
		log.Fatal(err)
	}

	result.Log().Debug()
	log.Debugf("%s\n", result.Payload)
	return result.Payload
}

func (c *Client) verbosePut(object Findable) (Result, error) {
	serializedObject, err := Marshal(object)
	if err != nil {
		log.Errorf("Client.verbosePut: %v", err)
		return Result{}, err
	}
	request := newRequest("PUT", object.Path(), c.Token, serializedObject)
	result, err := c.performRequest(request)
	if err != nil {
		log.Errorf("Client.verbosePut: %v", err)
		return Result{}, err
	}
	return result, nil
}

func (c *Client) _delete(path string) []byte {
	request := newRequest("DELETE", path, c.Token, nil)
	result, err := c.performRequest(request)
	if err != nil {
		log.Fatal(err)
	}

	result.Log().Debugf("response payload: %s\n", result.Payload)
	return result.Payload
}

// insecureClient returns an HTTP client that will not verify the validity
// of an SSL certificate when performing a request.
func insecureClient() *http.Client {
	// pool = x509.NewCertPool()
	// pool.AppendCertsFromPEM(pemCerts)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		// RootCAs:            pool},
	}
	return &http.Client{Transport: tr}
}

// performRequest performs a request using the given parameters and
// returns a struct that contains the HTTP status code and payload from
// the server's response as well as metadata such as the response time.
func (c Client) performRequest(p request) (Result, error) {
	uri := ESPAPIRoot + p.Path

	if p.requiresAnObject() && p.Object != nil {
		log.Debugf("Received serialized object: %s", p.Object)
	}
	req, err := http.NewRequest(p.Verb, uri, bytes.NewBuffer(p.Object))
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Debug("Client.performRequest")
		return Result{}, err
	}
	p.httpRequest = req

	p.addHeaders(p.Token, c.APIKey)

	result, err := getResult(insecureClient(), req)
	if err != nil {
		log.Error(err)
		return Result{}, err
	}
	return Result{p, result}, nil
}

func tokenFrom(payload []byte) Token {
	var response map[string]string
	json.Unmarshal(payload, &response)
	return Token(response["access_token"])
}
