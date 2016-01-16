package espsdk

import (
	"encoding/json"
	"time"

	log "github.com/Sirupsen/logrus"
)

// A Token is a string representation of an OAuth2 token. It grants a user
// access to the ESP API for a limited time.
type Token string

// A Serializable object can be serialized to a byte stream such as JSON.
type serializable interface {
	Marshal() ([]byte, error)
}

// A FulfilledRequest provides an overview of a completed API request and
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
	return log.WithFields(r.Stats())
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

func indentedJSON(obj interface{}) ([]byte, error) {
	return json.MarshalIndent(obj, "", "\t")
}
