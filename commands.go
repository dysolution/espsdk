package espsdk

// PrettyPrintable applies to all objects that should have an easy-to-read
// JSON representation of themselves availalbe for printing.
type PrettyPrintable interface {
	PrettyPrint() string
}

type Createable interface {
	PrettyPrintable
	Path() string
	Marshal() ([]byte, error)
	Unmarshal([]byte) Createable
}

func Create(object Createable, client *Client) Createable {
	marshaledObject := client.newPost(object)
	return object.Unmarshal(marshaledObject)
}
