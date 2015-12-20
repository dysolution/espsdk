package espsdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

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

// Request performs a request using the provided HTTP verb and returns
// the response as a JSON payload. If the verb is POST, the optional
// serialized object will become the body of the HTTP request.
func (c Client) Request(p *RequestParams) *FulfilledRequest {
	uri := endpoint + p.Path

	if p.requiresAnObject() && p.Object != nil {
		log.Debugf("Received serialized object: %s", p.Object)
	}
	req, err := http.NewRequest(p.Verb, uri, bytes.NewBuffer(p.Object))
	if err != nil {
		log.Fatal(err)
	}
	httpClient := insecureClient()

	result := getJSON(httpClient, req, p.Token, c.APIKey)
	if result.Err != nil {
		log.Fatal(result.Err)
		return &FulfilledRequest{
			p,
			&Result{
				&Response{
					result.Response.StatusCode,
					result.Response.Status,
				},
				nil,
				result.Duration,
				result.Err,
			},
		}
	}
	return &FulfilledRequest{p, result}
}
