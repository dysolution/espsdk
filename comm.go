package espsdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
)

const (
	endpoint      = "https://esp-sandbox.api.gettyimages.com/esp"
	oauthEndpoint = "https://api.gettyimages.com/oauth2/token"
)

type Token string

type Credentials struct {
	APIKey      string
	APISecret   string
	ESPUsername string
	ESPPassword string
}

func (c *Credentials) areInvalid() bool {
	if len(c.APIKey) < 1 || len(c.APISecret) < 1 || len(c.ESPUsername) < 1 || len(c.ESPPassword) < 1 {
		return true
	}
	return false
}

func (c *Credentials) formValues() url.Values {
	v := url.Values{}
	v.Set("client_id", c.APIKey)
	v.Set("client_secret", c.APISecret)
	v.Set("username", c.ESPUsername)
	v.Set("password", c.ESPPassword)
	v.Set("grant_type", "password")
	return v
}

type Client struct {
	Credentials
	UploadBucket string
}

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

// request performs a request using the provided HTTP verb and returns
// the response as a JSON payload. If the verb is POST, the optional
// serialized object will become the body of the HTTP request.
func (c Client) Request(verb string, path string, token Token, object []byte) ([]byte, error) {
	uri := endpoint + path
	log.Debug(uri)

	if (verb == "POST" || verb == "PUT") && object != nil {
		log.Debugf("Received serialized object: %s", object)
	}
	req, err := http.NewRequest(verb, uri, bytes.NewBuffer(object))
	httpClient := insecureClient()

	payload, err := getJSON(httpClient, req, token, c.APIKey)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return payload, nil
}

// Private

func getJSON(c *http.Client, req *http.Request, token Token, apiKey string) ([]byte, error) {
	req.Header.Set("Authorization", "Token token="+string(token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", apiKey)

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

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
