package espsdk

import (
	"encoding/json"
	"log"
)

// GetKeywords requests suggestions from the Getty controlled vocabulary
// for the keywords provided.
func GetKeywords(client *Client) []byte { return client.get(Keywords) }

// GetPersonalities requests suggestions from the Getty controlled vocabulary
// for the famous personalities provided.
func GetPersonalities(client *Client) []byte { return client.get(Personalities) }

// GetControlledValues returns complete lists of values and descriptions for
// fields with controlled vocabularies, grouped by submission type.
func GetControlledValues(client *Client) []byte { return client.get(ControlledValues) }

// GetTranscoderMappings lists acceptable transcoder mapping values
// for Getty and iStock video.
func GetTranscoderMappings(client *Client) []byte { return client.get(TranscoderMappings) }

// GetCompositions lists all possible composition values.
func GetCompositions(client *Client) []byte { return client.get(Compositions) }

// GetExpressions lists all possible facial expression values.
func GetExpressions(client *Client) []byte { return client.get(Expressions) }

type PeopleMetadata struct {
	Term     string `json:"term,omitempty"`
	TermID   int    `json:"term_id,omitempty"`
	ImageURI string `json:"image_uri,omitempty"`
	HelpText string `json:"help_text,omitempty"`
}

type PeopleMetadataList []PeopleMetadata

// Marshal serializes PeopleMetadata into a byte slice of indented JSON.
func (m PeopleMetadataList) Marshal() ([]byte, error) { return indentedJSON(m) }

// Unmarshal attempts to deserialize the provided JSON payload into a
// representation of people metadata.
func (m PeopleMetadataList) Unmarshal(payload []byte) PeopleMetadataList {
	var items PeopleMetadataList
	if err := json.Unmarshal(payload, &items); err != nil {
		log.Fatal(err)
	}
	return items
}

// GetNumberOfPeople lists all possible values for Number of People.
func (m PeopleMetadataList) GetNumberOfPeople(client *Client) PeopleMetadataList {
	return m.Unmarshal(client.get(NumberOfPeople))
}

// PrettyPrint returns a human-readable serialized JSON representation of
// the provided object.
func (m PeopleMetadataList) PrettyPrint() string { return prettyPrint(m) }
