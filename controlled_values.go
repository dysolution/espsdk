package espsdk

import (
	"encoding/json"
	"fmt"

	"github.com/dysolution/sleepwalker"
)

// A TermItem is an expression of a concept that has a canonical string to
// describe it and an optional image_uri and help_text. TermItems are the
// base unit of Keywords, Personalities, Facial Expressions, and others.
type TermItem struct {
	Term     string `json:"term,omitempty"`
	TermID   string `json:"term_id,omitempty"`
	HelpText string `json:"help_text,omitempty"`
	ImageURI string `json:"image_uri,omitempty"`
}

// A TermList is an array (slice) of terms (TermItems).
type TermList []TermItem

// Marshal serializes a TermList into readable JSON.
func (m TermList) Marshal() ([]byte, error) {
	return sleepwalker.Marshal(m)
}

// Unmarshal attempts to deserialize the provided JSON payload into a
// representation of people metadata.
func (m TermList) Unmarshal(payload []byte) *TermList {
	var items TermList
	err := json.Unmarshal(payload, &items)
	if err != nil {
		Log.WithFields(map[string]interface{}{
			"items": fmt.Sprintf("%v", items),
		}).Error(err)
	}
	return &items
}
