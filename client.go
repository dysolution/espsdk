package espsdk

import (
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
func GetClient(key, secret, username, password string, log *logrus.Logger) Client {
	return Client{sleepwalker.GetClient(key, secret, username, password, OAuthEndpoint, ESPAPIRoot, log)}
}

// GetKeywords requests suggestions from the Getty controlled vocabulary
// for the keywords provided.
//
// TODO: not implemented (keywords and personalities need a new struct type)
func (c Client) GetKeywords() []byte { return []byte("not implemented") }

// GetPersonalities requests suggestions from the Getty controlled vocabulary
// for the famous personalities provided.
//
// TODO: not implemented (keywords and personalities need a new struct type)
func (c Client) GetPersonalities() []byte { return []byte("not implemented") }

// GetControlledValues returns complete lists of values and descriptions for
// fields with controlled vocabularies, grouped by submission type.
func (c Client) GetControlledValues() ControlledValues {
	desc := "Client.GetControlledValues"
	result, err := c.GetPath(ControlledValuesEndpoint)
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
	result, err := c.GetPath(TranscoderMappingsEndpoint)
	if err != nil {
		return &TranscoderMappingList{}
	}
	if result.Payload == nil {
		return &TranscoderMappingList{}
	}
	result.Log().Info(desc)
	return TranscoderMappingList{}.Unmarshal(result.Payload)
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
