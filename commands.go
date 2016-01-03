package espsdk

import "encoding/json"

// A DeserializedObject contains JSON struct tags that map object properties
// to JSON fields.
type DeserializedObject struct {
	Batch
	Contribution
	Release
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
