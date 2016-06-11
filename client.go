package espsdk

import (
	"encoding/json"
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/dysolution/sleepwalker"
)

// A Client communicates with the ESP REST API.
type Client struct {
	*sleepwalker.Client
}

// A Result provides information about the response from the ESP REST API.
type Result struct {
	sleepwalker.Result
}

// GetClient provides a client for communicating with the ESP REST API.
func GetClient(key, secret, username, password, apiRoot string, log *logrus.Logger) Client {
	config := &sleepwalker.Config{
		Credentials: &sleepwalker.Credentials{
			APIKey:    key,
			APISecret: secret,
			Username:  username,
			Password:  password,
		},
		OAuthEndpoint: OAuthEndpoint,
		APIRoot:       apiRoot,
		Logger:        log,
	}
	return Client{sleepwalker.GetClient(config)}
}

// ValidateKeywords queries the ESP keywords endpoint and reports whether each
// provided keyword is valid.
func (c Client) ValidateKeywords(keywords []string, mediaType string) []Keyword {
	reqPayload := struct {
		Keywords  []string `json:"keywords"`
		MediaType string   `json:"media_type"`
	}{
		Keywords:  keywords,
		MediaType: mediaType,
	}

	bytes, _ := json.Marshal(reqPayload)
	result, _ := c.GetWithPayload(Endpoints.Keywords, bytes)
	var payload map[string]map[string][]interface{}
	if err := json.Unmarshal(result.Payload, &payload); err != nil {
		Log.Error(err)
	}

	out, _ := json.MarshalIndent(payload, "", "  ")
	Log.WithFields(map[string]interface{}{
		"keyword_response": out,
	}).Debugf("%s", out)

	var outKW []Keyword
	for inKW, matches := range payload["keywords"] {
		Log.Debugf("checking keyword: %v", inKW)
		if len(matches) == 1 {
			outKW = append(outKW, Keyword{Term: inKW, Valid: true})
		} else {
			// no matches, or further disambiguation is required (TODO)
			outKW = append(outKW, Keyword{Term: inKW, Valid: false})
		}
	}
	return outKW
}

// GetControlledValues returns complete lists of values and descriptions for
// fields with controlled vocabularies, grouped by submission type.
func (c Client) GetControlledValues() ControlledValues {
	desc := "Client.GetControlledValues"
	result, err := c.GetPath(Endpoints.ControlledValues)
	if err != nil {
		return ControlledValues{}
	}
	result.Log().Info(desc)
	return parseCV(result)
}

// GetTranscoderMappings lists acceptable transcoder mapping values
// for Getty and iStock video.
func (c Client) GetTranscoderMappings() *TranscoderMappingList {
	desc := "Client.GetTranscoderMappings"
	result, err := c.GetPath(Endpoints.TranscoderMappings)
	if err != nil {
		return &TranscoderMappingList{}
	}
	if result.Payload == nil {
		return &TranscoderMappingList{}
	}
	result.Log().Info(desc)
	return TranscoderMappingList{}.Unmarshal(result.Payload)
}

// GetEvents returns a list of events that match the provided criteria.
func (c Client) GetEvents(params EventQuery) (*EventResponse, error) {
	desc := "Client.GetEvents"
	bytes, _ := json.Marshal(params)
	result, err := c.GetWithPayload(Endpoints.Events, bytes)
	if err != nil {
		return &EventResponse{}, err
	}
	if result.Payload == nil {
		return &EventResponse{}, errors.New("empty payload")
	}
	result.Log().Info(desc)
	return Event{}.Unmarshal(result.Payload), nil
}

// GetTermList lists all possible values for the given controlled vocabulary.
func (c Client) GetTermList(endpoint string) *TermList {
	desc := "Client.GetTermList"
	result, err := c.GetPath(endpoint)
	if err != nil {
		return &TermList{}
	}
	if result.Payload == nil {
		return &TermList{}
	}
	result.Log().Info(desc)
	return TermList{}.Unmarshal(result.Payload)
}

// GetTermIntList lists all possible values for the given controlled vocabulary.
func (c Client) GetTermIntList(endpoint string) *TermIntList {
	desc := "Client.GetTermListInt"
	result, err := c.GetPath(endpoint)
	if err != nil {
		return &TermIntList{}
	}
	if result.Payload == nil {
		return &TermIntList{}
	}
	result.Log().Info(desc)
	return TermIntList{}.Unmarshal(result.Payload)
}

// DeleteLastBatch looks up the newest Batch and deletes it.
func DeleteLastBatch(c sleepwalker.RESTClient) (sleepwalker.Result, error) {
	lastBatch := Batch{}.Index(c).Last()
	return c.Delete(lastBatch)
}

// SubmitLastPhoto subtmits the newest Contribution for review and publication.
func SubmitLastPhoto(c sleepwalker.RESTClient) (sleepwalker.Result, error) {
	newestBatch := Batch{}.Index(c).Last()
	newestContribution, err := Contribution{
		SubmissionBatchID: newestBatch.ID,
	}.Index(c, newestBatch.ID).Last()
	if err != nil {
		return sleepwalker.Result{}, err
	}
	return c.Put(newestContribution, newestContribution.Path()+"/submit")
}
