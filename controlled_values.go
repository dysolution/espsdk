package espsdk

import (
	"encoding/json"
)

// A TermItem is an expression of a concept that has a canonical string to
// describe it and an optional image_uri and help_text. TermItems are the
// base unit of Keywords, Personalities, Facial Expressions, and others.
type TermItem struct {
	Term     string `json:"term,omitempty"`
	TermID   int    `json:"term_id,omitempty"`
	ImageURI string `json:"image_uri,omitempty"`
	HelpText string `json:"help_text,omitempty"`
}

// A TermList is an array (slice) of terms (TermItems).
type TermList []TermItem

// Marshal serializes a TermList into a byte slice of indented JSON.
func (m TermList) Marshal() ([]byte, error) { return indentedJSON(m) }

// Unmarshal attempts to deserialize the provided JSON payload into a
// representation of people metadata.
func (m TermList) Unmarshal(payload []byte) TermList {
	var items TermList
	if err := json.Unmarshal(payload, &items); err != nil {
		log.Fatal(err)
	}
	return items
}
