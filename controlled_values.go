package espsdk

import (
	"encoding/json"
	"log"
)

// GetKeywords requests suggestions from the Getty controlled vocabulary
// for the keywords provided.
//
// TODO: not implemented (keywords and personalities need a new struct type)
func GetKeywords(client *Client) []byte { return client.get(Keywords) }

// GetPersonalities requests suggestions from the Getty controlled vocabulary
// for the famous personalities provided.
//
// TODO: not implemented (keywords and personalities need a new struct type)
func GetPersonalities(client *Client) []byte { return client.get(Personalities) }

// GetControlledValues returns complete lists of values and descriptions for
// fields with controlled vocabularies, grouped by submission type.
//
// TODO: not implemented (needs new struct type)
func GetControlledValues(client *Client) []byte { return client.get(ControlledValues) }

// GetTranscoderMappings lists acceptable transcoder mapping values
// for Getty and iStock video.
//
// TODO: not implemented (needs new struct type)
func GetTranscoderMappings(client *Client) []byte { return client.get(TranscoderMappings) }

type TermItem struct {
	Term     string `json:"term,omitempty"`
	TermID   int    `json:"term_id,omitempty"`
	ImageURI string `json:"image_uri,omitempty"`
	HelpText string `json:"help_text,omitempty"`
}

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

// GetNumberOfPeople lists all possible values for Number of People.
func (m TermList) GetNumberOfPeople(client *Client) TermList {
	return m.Unmarshal(client.get(NumberOfPeople))
}

// GetExpressions lists all possible facial expression values.
func (m TermList) GetExpressions(client *Client) TermList {
	return m.Unmarshal(client.get(Expressions))
}

// GetCompositions lists all possible composition values.
func (m TermList) GetCompositions(client *Client) TermList {
	return m.Unmarshal(client.get(Compositions))
}

// PrettyPrint returns a human-readable serialized JSON representation of
// the provided object.
func (m TermList) PrettyPrint() string { return prettyPrint(m) }
