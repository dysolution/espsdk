package espsdk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

// A Result provides an overview of a completed API request and
// its result, including timing and HTTP status codes.
type Result struct {
	*request
	*VerboseResult
}

// Marshal serializes a FulfilledRequest into a byte stream.
func (r *Result) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// MarshalIndent serializes a FulfilledRequest into indented JSON.
func (r *Result) MarshalIndent() ([]byte, error) {
	return json.MarshalIndent(r, "", "    ")
}

// Stats returns fields that logrus can parse.
func (r *Result) Stats() log.Fields {
	return log.Fields{
		"method":        r.Verb,
		"path":          r.Path,
		"response_time": r.Duration * time.Millisecond,
		"status":        r.Response.Status,
		"status_code":   r.Response.StatusCode,
	}
}

func (r *Result) Log() *log.Entry {
	return log.WithFields(log.Fields{
		"response_time": r.Duration * time.Millisecond,
		"status_code":   r.Response.StatusCode,
	})
}

func (r *Result) GetStatusCode() int {
	return r.VerboseResult.GetStatusCode()
}

// Private

// A Response contains the HTTP status code and text that represent the API's
// response to a request.
type response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}

func (r *response) GetStatusCode() int {
	return r.StatusCode
}

// A VerboseResult contains information relative to a completed request,
// including the time elapsed to fulfill the request and any errors.
type VerboseResult struct {
	Response *response     `json:"response"`
	Payload  []byte        `json:"-"`
	Duration time.Duration `json:"response_time"`
	Err      error         `json:"-"`
}

func (v *VerboseResult) GetStatusCode() int {
	return v.Response.GetStatusCode()
}

func getResult(c *http.Client, req *http.Request) (*VerboseResult, error) {
	start := time.Now()
	resp, err := c.Do(req)
	duration := time.Since(start) / time.Millisecond
	if err != nil {
		log.Error(err)
		return buildResult(resp, nil, duration, err), err
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return buildResult(resp, payload, duration, err), nil
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.WithFields(log.Fields{
			"object":      "response",
			"status_code": resp.StatusCode,
			"status":      resp.Status,
		}).Debug("getResult")
	}
	return buildResult(resp, payload, duration, nil), nil
}

func buildResult(resp *http.Response, payload []byte, duration time.Duration, err error) *VerboseResult {
	return &VerboseResult{
		&response{
			resp.StatusCode,
			resp.Status,
		},
		payload, duration, err}
}
