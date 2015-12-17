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
	uri := endpoint + path
	log.Debug(uri)
	c := insecureClient()

	req, err := http.NewRequest("GET", uri, nil)
	payload, err := getJSON(c, req, token, espClient.Credentials.APIKey)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return payload, nil
}

func (espClient Client) Post(o []byte, token Token, path string) ([]byte, error) {
	log.Debugf("Received serialized object: %s", o)

	c := insecureClient()

	uri := endpoint + path
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(o))
	payload, err := getJSON(c, req, token, espClient.Credentials.APIKey)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return payload, nil
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
