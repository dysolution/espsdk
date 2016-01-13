package espsdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

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
func GetClient(key, secret, username, password, uploadBucket string) Client {
	creds := credentials{
		APIKey:      key,
		APISecret:   secret,
		ESPUsername: username,
		ESPPassword: password,
	}
	token := getToken(&creds)
	return Client{creds, token, uploadBucket}
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

// Index requests a list of all Batches owned by the user.
func (c *Client) Index(path string) *DeserializedObject {
	var obj *DeserializedObject
	return Deserialize(c.get(path), obj)
}

// Create uses the provided path and data to ask the API to create a new
// object and returns the deserialized response.
func (c *Client) Create(object RESTObject) DeserializedObject {
	marshaledObject := c.post(object)
	return Unmarshal(marshaledObject)
}

// Update changes metadata for an existing Batch.
func (c *Client) Update(object RESTObject) DeserializedObject {
	return Unmarshal(c.put(object))
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
func (c *Client) Get(path string) DeserializedObject {
	return Unmarshal(c.get(path))
}

// GetFromObject requests the metadata for the provided object, as long as
// enough data is provided to unambiguously identify it to the API.
func (c *Client) GetFromObject(object RESTObject) DeserializedObject {
	return Unmarshal(c.get(object.Path()))
}

func (c *Client) get(path string) []byte {
	request := newRequest("GET", path, c.Token, nil)
	result, err := c.performRequest(request)
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(result.Stats()).Info()
	log.Debugf("%s\n", result.Payload)
	return result.Payload
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

	log.WithFields(result.Stats()).Info()
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

	log.WithFields(result.Stats()).Info()
	log.Debugf("%s\n", result.Payload)
	return result.Payload
}

func (c *Client) _delete(path string) []byte {
	request := newRequest("DELETE", path, c.Token, nil)
	result, err := c.performRequest(request)
	if err != nil {
		log.Fatal(err)
	}

	log.WithFields(result.Stats()).Info()
	log.Debugf("response payload: %s\n", result.Payload)
	return result.Payload
}

// insecureClient returns an HTTP client that will not verify the validity
// of an SSL certificate when performing a request.
func insecureClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

// performRequest performs a request using the given parameters and
// returns a struct that contains the HTTP status code and payload from
// the server's response as well as metadata such as the response time.
func (c Client) performRequest(p *request) (*fulfilledRequest, error) {
	uri := ESPAPIRoot + p.Path

	if p.requiresAnObject() && p.Object != nil {
		log.Debugf("Received serialized object: %s", p.Object)
	}
	req, err := http.NewRequest(p.Verb, uri, bytes.NewBuffer(p.Object))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	p.httpRequest = req

	p.addHeaders(p.Token, c.APIKey)

	result, err := getResult(insecureClient(), req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &fulfilledRequest{p, result}, nil
}

func tokenFrom(payload []byte) Token {
	var response map[string]string
	json.Unmarshal(payload, &response)
	return Token(response["access_token"])
}
