package espapi

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
)

const endpoint = "https://esp-sandbox.api.gettyimages.com/esp"

type ApiClient interface {
	PostBatch(SubmissionBatch) error
	PostRelease(Release) error
}

type Credentials struct {
	ApiKey      string
	ApiSecret   string
	EspUsername string
	EspPassword string
}

type Client struct {
	Credentials
	UploadBucket string
}

type Token string

func (espClient Client) Post(o []byte, token Token, path string) ([]byte, error) {
	log.Debugf("Received serialized object: %s", o)

	v := url.Values{}
	v.Set("Authorization", fmt.Sprintf("Token token=%s", token))
	v.Set("Content-Type", "application/json")

	uri := endpoint + path
	log.Debug(uri)
	log.Debug(v.Encode())
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(o))
	req.Header.Set("Authorization", fmt.Sprintf("Token token=%s", token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", espClient.Credentials.ApiKey)

	resp, err := c.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	payload, err := ioutil.ReadAll(resp.Body)
	log.Infof("HTTP %s", resp.Status)
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
