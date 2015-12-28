package espsdk

import "encoding/json"

// PrettyPrintable applies to all objects that should have an easy-to-read
// JSON representation of themselves availalbe for printing.
type PrettyPrintable interface {
	PrettyPrint() string
}

type DeserializedObject struct {
	Batch
}

type Createable interface {
	PrettyPrintable
	Path() string
	Marshal() ([]byte, error)
}

func (do DeserializedObject) PrettyPrint() string {
	prettyOutput, err := do.Marshal()
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

func Create(object Createable, client *Client) DeserializedObject {
	marshaledObject := client.newPost(object)
	return Unmarshal(marshaledObject)
}

// Get requests the metadata for a specific object.
func Get(path string, client *Client) DeserializedObject {
	return Unmarshal(client.get(path))
}

// Unmarshal attempts to deserialize the provided JSON payload
// into an object.
func Unmarshal(payload []byte) DeserializedObject {
	var dest DeserializedObject
	if err := json.Unmarshal(payload, &dest); err != nil {
		panic(err)
	}
	return dest
}
