package espsdk

import "encoding/json"

// PrettyPrintable applies to all objects that should have an easy-to-read
// JSON representation of themselves availalbe for printing.
type PrettyPrintable interface {
	PrettyPrint() string
}

type DeserializedObject struct {
	Batch
	Contribution
	Release
}

type Createable interface {
	PrettyPrintable
}

func (do DeserializedObject) PrettyPrint() string {
	prettyOutput, err := Marshal(do)
	if err != nil {
		panic(err)
	}
	return string(prettyOutput)
}

// Unmarshal attempts to deserialize the provided JSON payload
// into an object.
func (do DeserializedObject) Unmarshal(payload []byte) DeserializedObject {
	return Unmarshal(payload)
}

// Create uses the provided path and data to ask the API to create a new
// object and returns the deserialized response.
func Create(path string, object interface{}, client *Client) DeserializedObject {
	marshaledObject := client.post(object, path)
	return Unmarshal(marshaledObject)
}

// Delete destroys a specific object.
func Delete(path string, client *Client) { client._delete(path) }

// Get requests the metadata for a specific object.
func Get(path string, client *Client) DeserializedObject {
	return Unmarshal(client.get(path))
}

// Marshal serializes an object into a byte slice.
func Marshal(object interface{}) ([]byte, error) { return indentedJSON(object) }

// Unmarshal attempts to deserialize the provided JSON payload
// into an object.
func Unmarshal(payload []byte) DeserializedObject {
	var dest DeserializedObject
	if err := json.Unmarshal(payload, &dest); err != nil {
		panic(err)
	}
	return dest
}
