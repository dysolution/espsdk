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

// request performs a request using the provided HTTP verb and returns
// the response as a JSON payload. If the verb is POST, the optional
// serialized object will become the body of the HTTP request.
func (espClient Client) Request(verb string, path string, token Token, object []byte) ([]byte, error) {
	uri := endpoint + path
	log.Debug(uri)

	if verb == "POST" && object != nil {
		log.Debugf("Received serialized object: %s", object)
	}
	req, err := http.NewRequest(verb, uri, bytes.NewBuffer(object))
	c := insecureClient()

	payload, err := getJSON(c, req, token, espClient.Credentials.APIKey)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return payload, nil
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

// insecureClient returns an HTTP client that will not verify the validity
// of an SSL certificate when performing a request.
func insecureClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
