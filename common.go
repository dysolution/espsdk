package espsdk

import (
	"encoding/json"
)

// A Token is a string representation of an OAuth2 token. It grants a user
// access to the ESP API for a limited time.
type Token string

// A Serializable object can be serialized to a byte stream such as JSON.
type serializable interface {
	Marshal() ([]byte, error)
}

func indentedJSON(obj interface{}) ([]byte, error) {
	return json.MarshalIndent(obj, "", "\t")
}
