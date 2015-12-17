package espapi

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

const endpoint = "https://esp-sandbox.api.gettyimages.com/esp"

type APIClient interface {
	PostBatch(SubmissionBatch) error
	PostRelease(Release) error
}

type Credentials struct {
	APIKey      string
	APISecret   string
	ESPUsername string
	ESPPassword string
}

type Client struct {
	Credentials
	UploadBucket string
}

type Token string

func (espClient Client) Get(path string, token Token) ([]byte, error) {
	payload, err := espClient.request("GET", path, token, nil)
	return payload, err
}

func (espClient Client) Post(o []byte, token Token, path string) ([]byte, error) {
	payload, err := espClient.request("POST", path, token, o)
	return payload, err
}

func (client Client) PostRelease(r []byte) {
	log.Infof("Received serialized release: %s", r)
	client.Call()
}

func (client Client) PostContribution(c []byte) {
	log.Infof("Received serialized contribution: %s", c)
	client.Call()
}

func (c Client) Call() {
}

// Private

func getJSON(c *http.Client, req *http.Request, token Token, apiKey string) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Token token=%s", token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", apiKey)

	resp, err := c.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Infof("HTTP %s", resp.Status)
	return payload, nil
}

func insecureClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

// request performs a request using the provided HTTP verb and returns
// the response as a JSON payload. If the verb is POST, the optional
// serialized object will become the body of the HTTP request.
func (espClient Client) request(verb string, path string, token Token, object []byte) ([]byte, error) {
	var body *bytes.Buffer
	if verb == "POST" && object != nil {
		log.Debugf("Received serialized object: %s", object)
		body = bytes.NewBuffer(object)
	}
	uri := endpoint + path
	log.Debug(uri)
	c := insecureClient()

	req, err := http.NewRequest(verb, uri, body)
	payload, err := getJSON(c, req, token, espClient.Credentials.APIKey)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return payload, nil
}
