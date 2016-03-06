package espsdk

import (
	"encoding/json"
	"fmt"

	"github.com/dysolution/sleepwalker"
)

// Each Keyword is validated against a controlled vocabulary and thus can
// be valid or invalid. This structure is also used for Personalities, which
// allows recognizable people to have a canonical representation of their name
// across all Getty systems.
type Keyword struct {
	Term  string `json:"term,omitempty"`
	Valid bool   `json:"valid"`
}

// A TermItem is an expression of a concept that has a canonical string to
// describe it and an optional image_uri and help_text. TermItems are the
// base unit of Facial Expressions and Number of People.
type TermItem struct {
	Term     string `json:"term,omitempty"`
	TermID   string `json:"term_id,omitempty"`
	HelpText string `json:"help_text,omitempty"`
	ImageURI string `json:"image_uri,omitempty"`
}

func (ti TermItem) Validate(input string, corpus TermList) TermItem {
	Log.Debugf("checking ID: %v", input)
	for _, validTermItem := range corpus {
		if validTermItem.TermID == input {
			Log.Debugf("match: %v == %v", input, validTermItem)
			return validTermItem
		}
	}
	return TermItem{}
}

func (ti TermItem) ValidateList(input []string, corpus TermList) []TermItem {
	var validatedItems []TermItem
	for _, candidateID := range input {
		validatedItems = append(validatedItems, ti.Validate(candidateID, corpus))
	}
	return validatedItems
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
