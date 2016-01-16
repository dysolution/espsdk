package espsdk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
)

// A Result provides an overview of a completed API request and
// its result, including timing and HTTP status codes.
type Result struct {
	request
	VerboseResult
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
func (r *Result) stats() logrus.Fields {
	return logrus.Fields{
		"method":        r.Verb,
		"path":          r.Path,
		"response_time": r.Duration * time.Millisecond,
		"status_code":   r.Response.StatusCode,
	}
}

// Log provides a convenient way to output information about an HTTP request
// the user is likely to want.
func (r *Result) Log() *logrus.Entry {
	return Log.WithFields(r.stats())
}

// LogBrief logs only an HTTP request's status code and response time.
func (r *Result) LogBrief() *logrus.Entry {
	return Log.WithFields(logrus.Fields{
		"response_time": r.Duration * time.Millisecond,
		"status_code":   r.Response.StatusCode,
	})
}

// A Response contains the HTTP status code and text that represent the API's
// response to a request.
type Response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}

// A VerboseResult contains information relative to a completed request,
// including the time elapsed to fulfill the request and any errors.
type VerboseResult struct {
	Response `json:"response"`
	Payload  []byte        `json:"-"`
	Duration time.Duration `json:"response_time"`
}

func getResult(c *http.Client, req *http.Request) (VerboseResult, error) {
	start := time.Now()
	resp, err := c.Do(req)
	duration := time.Since(start) / time.Millisecond
	if err != nil {
		Log.Error(err)
		return buildResult(resp, nil, duration), err
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.Error(err)
		return buildResult(resp, payload, duration), nil
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		Log.WithFields(logrus.Fields{
			"object":      "response",
			"status_code": resp.StatusCode,
			"status":      resp.Status,
		}).Debug("getResult")
	}
	return buildResult(resp, payload, duration), nil
}

func buildResult(resp *http.Response, payload []byte, duration time.Duration) VerboseResult {
	return VerboseResult{
		Response{
			resp.StatusCode,
			resp.Status,
		},
		payload,
		duration,
	}
}
