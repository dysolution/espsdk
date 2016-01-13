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
type fulfilledRequest struct {
	*request
	*result
}

// Marshal serializes a FulfilledRequest into a byte stream.
func (r *fulfilledRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *fulfilledRequest) MarshalIndent() ([]byte, error) {
	return json.MarshalIndent(r, "", "    ")
}

func (r *fulfilledRequest) Stats() log.Fields {
	return log.Fields{
		"method":        r.Verb,
		"path":          r.Path,
		"response_time": r.Duration * time.Millisecond,
		"status":        r.Response.Status,
		"status_code":   r.Response.StatusCode,
	}
}

// Private

// A Response contains the HTTP status code and text that represent the API's
// response to a request.
type response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}

func indentedJSON(obj interface{}) ([]byte, error) {
	return json.MarshalIndent(obj, "", "\t")
}
