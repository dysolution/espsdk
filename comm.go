package espsdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

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

type FulfilledRequest struct {
	*RequestParams
	*Result
}

func (r *FulfilledRequest) Marshal() ([]byte, error) { return json.Marshal(r) }

type RequestParams struct {
	Verb   string `json:"method"`
	Path   string `json:"path"`
	Token  Token  `json:"-"`
	Object []byte `json:"-"`
}

type Response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"-"`
}

type Result struct {
	Response *Response     `json:"response"`
	Payload  []byte        `json:"-"`
	Duration time.Duration `json:"duration"`
	Err      error         `json:"-"`
}

// Request performs a request using the provided HTTP verb and returns
// the response as a JSON payload. If the verb is POST, the optional
// serialized object will become the body of the HTTP request.
func (c Client) Request(p *RequestParams) FulfilledRequest {
	uri := endpoint + p.Path

	if (p.Verb == "POST" || p.Verb == "PUT") && p.Object != nil {
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
		return FulfilledRequest{
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
	return FulfilledRequest{p, &result}
}

// Private

func getJSON(c *http.Client, req *http.Request, token Token, apiKey string) Result {
	httpCommand := req.Method + " " + string(req.URL.Path)
	start := start(httpCommand)
	req.Header.Set("Authorization", "Token token="+string(token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", apiKey)

	resp, err := c.Do(req)
	duration := elapsed(httpCommand, start)
	if err != nil {
		log.Fatal(err)
		return Result{
			&Response{
				resp.StatusCode,
				resp.Status,
			},
			nil, duration, err}
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return Result{
			&Response{
				resp.StatusCode,
				resp.Status,
			},
			nil, duration, err}
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Warnf("HTTP %s", resp.Status)
	}
	return Result{
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
