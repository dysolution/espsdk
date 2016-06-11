package espsdk

import "encoding/json"

// An Event represents an event's metadata.
type Event struct {
	DateFrom    string `json:"date_from,omitempty"`
	DateTo      string `json:"date_to,omitempty"`
	EventID     string `json:"frame_size,omitempty"`
	Headline    string `json:"headline,omitempty"`
	LastModDate string `json:"last_mod_date,omitempty"`
}

// Unmarshal attempts to deserialize the provided JSON payload
// into an object.
func (e Event) Unmarshal(payload []byte) *EventResponse {
	dest := &EventResponse{}
	if err := json.Unmarshal(payload, dest); err != nil {
		Log.Error(err)
	}
	return dest
}

// An EventQuery is sent to the ESP API with optional criteria.
type EventQuery struct {
	DateFrom         string `json:"date_from,omitempty"`
	DateTo           string `json:"date_to,omitempty"`
	EventName        string `json:"event_name,omitempty"`
	MEID             string `json:"frame_size,omitempty"`
	PhotographerName string `json:"photographer_name,omitempty"`
}

// An EventResponse is sent by the ESP API to the client in response to an
// EventQuery.
type EventResponse struct {
	Events            []Event                `json:"events"`
	SearchInformation map[string]interface{} `json:"search_information"`
	Errors            []string               `json:"errors"`
}
