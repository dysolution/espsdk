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
	*request
	*Result
}

// Marshal serializes a FulfilledRequest into a byte stream.
func (r *FulfilledRequest) Marshal() ([]byte, error) { return json.Marshal(r) }

// A Response contains the HTTP status code and text that represent the API's
// response to a request.
type Response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"-"`
}

// Private

func start(s string) time.Time { return time.Now() }

func elapsed(s string, startTime time.Time) time.Duration {
	duration := time.Now().Sub(startTime)
	return duration
}

func indentedJSON(obj interface{}) ([]byte, error) {
	return json.MarshalIndent(obj, "", "\t")
}

func get(path string, token Token) []byte {
	request := NewRequest("GET", path, token, nil)
	result := Client{}.PerformRequest(request)
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
