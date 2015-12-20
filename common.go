/*
The ESP SDK provides a Go interface to the JSON API of Getty Images'
Enterprise Submission Portal (ESP).

You will need a username and password that allows you
to log in to https://sandbox.espaws.com/ as well as
an API Key and and API Secret, both of which are accessible
at https://developer.gettyimages.com/apps/mykeys/.
*/
package espsdk

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	endpoint      = "https://esp-sandbox.api.gettyimages.com/esp"
	oauthEndpoint = "https://api.gettyimages.com/oauth2/token"
	jsonIndent    = "\t"
)

// A Token is a string representation of an OAuth2 token. It grants a user
// access to the ESP API for a limited time.
type Token string

// A Serializable object can be serialized to a byte stream such as JSON.
type serializable interface {
	Marshal() ([]byte, error)
}

func prettyPrint(obj serializable) string {
	prettyOutput, err := obj.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return string(prettyOutput)
}

// A FulfilledRequest provides an overview of a completed API request and
// its result, including timing and HTTP status codes.
type FulfilledRequest struct {
	*RequestParams
	*Result
}

// Marshal serializes a FulfilledRequest into a byte stream.
func (r *FulfilledRequest) Marshal() ([]byte, error) { return json.Marshal(r) }

// RequestParams are provided to a Request to indicate the specific API
// endpoint and action to take. The Object is optional and applies only to
// endpoints that create or update items (POST and PUT).
type RequestParams struct {
	Verb   string `json:"method"`
	Path   string `json:"path"`
	Token  Token  `json:"-"`
	Object []byte `json:"-"`
}

func (p *RequestParams) requiresAnObject() bool {
	if p.Verb == "POST" || p.Verb == "PUT" || p.Verb == "DELETE" {
		return true
	}
	return false
}

// A Response contains the HTTP status code and text that represent the API's
// response to a request.
type Response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"-"`
}

// A Result contains information relative to a completed request, including
// the time elapsed to fulfill the request and any errors.
type Result struct {
	Response *Response     `json:"response"`
	Payload  []byte        `json:"-"`
	Duration time.Duration `json:"response_ms"`
	Err      error         `json:"-"`
}

// Private

func getJSON(c *http.Client, req *http.Request, token Token, apiKey string) *Result {
	httpCommand := req.Method + " " + string(req.URL.Path)
	start := start(httpCommand)
	req.Header.Set("Authorization", "Token token="+string(token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", apiKey)

	resp, err := c.Do(req)
	duration := elapsed(httpCommand, start) / time.Millisecond
	if err != nil {
		log.Fatal(err)
		return getResult(resp, nil, duration)
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return getResult(resp, payload, duration)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Warnf("HTTP %s", resp.Status)
	}
	return getResult(resp, payload, duration)
}

func getResult(resp *http.Response, payload []byte, duration time.Duration) *Result {
	return &Result{
		&Response{
			resp.StatusCode,
			resp.Status,
		},
		payload, duration, nil}
}

// insecureClient returns an HTTP client that will not verify the validity
// of an SSL certificate when performing a request.
func insecureClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func start(s string) time.Time { return time.Now() }

func elapsed(s string, startTime time.Time) time.Duration {
	duration := time.Now().Sub(startTime)
	return duration
}

func indentedJSON(obj interface{}) ([]byte, error) {
	return json.MarshalIndent(obj, "", "\t")
}

func get(path string, token Token) []byte {
	params := RequestParams{"GET", path, token, nil}
	result := Client{}.Request(&params)
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
