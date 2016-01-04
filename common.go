package espsdk

import (
	"encoding/json"
	"time"

	log "github.com/Sirupsen/logrus"
)

func foo() {
}

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
func (r *fulfilledRequest) Marshal() ([]byte, error) { return json.Marshal(r) }

// Private

// A Response contains the HTTP status code and text that represent the API's
// response to a request.
type response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"-"`
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
	request := newRequest("GET", path, token, nil)
	result := Client{}.performRequest(request)
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
