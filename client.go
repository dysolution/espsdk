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

// A Client is able to request an access token and submit HTTP requests to
// the ESP API.
type Client struct {
	Credentials
	UploadBucket string
}

// GetToken submits the provided credentials to Getty's OAuth2 endpoint
// and returns a token that can be used to authenticate HTTP requests to the
// ESP API.
func (c Client) GetToken() Token {
	if c.Credentials.areInvalid() {
		log.Fatal("Not all required credentials were supplied.")
	}

	uri := oauthEndpoint
	log.Debugf("%s", uri)
	formValues := c.formValues()
	log.Debugf("%s", formValues.Encode())

	resp, err := http.PostForm(uri, formValues)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	log.Debugf("HTTP %d", resp.StatusCode)
	log.Debugf("%s", payload)
	return c.tokenFrom(payload)
}

func (c Client) tokenFrom(payload []byte) Token {
	var response map[string]string
	json.Unmarshal(payload, &response)
	return Token(response["access_token"])
}

// PerformRequest performs a request using the given parameters and
// returns a struct that contains the HTTP status code and payload from
// the server's response as well as metadata such as the response time.
func (c Client) PerformRequest(p *request) *FulfilledRequest {
	uri := ESPAPIRoot + p.Path

	if p.requiresAnObject() && p.Object != nil {
		log.Debugf("Received serialized object: %s", p.Object)
	}
	req, err := http.NewRequest(p.Verb, uri, bytes.NewBuffer(p.Object))
	if err != nil {
		log.Fatal(err)
	}
	p.httpRequest = req

	p.addHeaders(p.Token, c.APIKey)

	result := getResult(insecureClient(), req)
	if result.Err != nil {
		log.Fatal(result.Err)
	}
	return &FulfilledRequest{p, result}
}

func (c *Client) get(path string) []byte {
	request := newRequest("GET", path, c.GetToken(), nil)
	result := c.PerformRequest(request)
	if result.Err != nil {
		log.Fatal(result.Err)
	}
	stats, err := result.Marshal()
	if err != nil {
		log.Fatal(result.Err)
	}
	log.Info(string(stats))
	log.Debugf("%s\n", result.Payload)
	return result.Payload
}

func (c *Client) post(object Serializable, path string) []byte {
	serializedObject, err := object.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	request := newRequest("POST", path, c.GetToken(), serializedObject)
	result := c.PerformRequest(request)
	if result.Err != nil {
		log.Fatal(result.Err)
	}

	stats, err := result.Marshal()
	if err != nil {
		log.Fatal(result.Err)
	}
	log.Info(string(stats))
	log.Debugf("%s\n", result.Payload)
	return result.Payload
}

func (c *Client) put(object Serializable, path string) []byte {
	serializedObject, err := object.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	request := newRequest("PUT", path, c.GetToken(), serializedObject)
	result := c.PerformRequest(request)
	if result.Err != nil {
		log.Fatal(result.Err)
	}

	stats, err := result.Marshal()
	if err != nil {
		log.Fatal(result.Err)
	}
	log.Info(string(stats))
	log.Debugf("%s\n", result.Payload)
	return result.Payload
}

func (c *Client) _delete(path string) {
	request := newRequest("DELETE", path, c.GetToken(), nil)
	result := c.PerformRequest(request)
	if result.Err != nil {
		log.Fatal(result.Err)
	}

	stats, err := result.Marshal()
	if err != nil {
		log.Fatal(result.Err)
	}
	log.Info(string(stats))
	log.Debugf("%s\n", result.Payload)
}

// insecureClient returns an HTTP client that will not verify the validity
// of an SSL certificate when performing a request.
func insecureClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
