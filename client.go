package espsdk

import (
	"errors"

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

func getClient(key, secret, username, password string) Client {
	return Client{sleepwalker.GetClient(key, secret, username, password, OAuthEndpoint, ESPAPIRoot, Log)}
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
//
// TODO: not implemented (needs new struct type)
func (c Client) GetControlledValues() ([]byte, error) {
	Log.Info("GetControlledValues")
	result, err := c.GetPath(ControlledValues)
	if err != nil {
		return []byte{}, errors.New("unable to get controlled values")
	}
	return result.Payload, nil
}

// GetTranscoderMappings lists acceptable transcoder mapping values
// for Getty and iStock video.
func GetTranscoderMappings(c Client) *TranscoderMappingList {
	result, err := c.GetPath(ESPAPIRoot + TranscoderMappings)
	if err != nil {
		return &TranscoderMappingList{}
	}
	if result.Payload == nil {
		return &TranscoderMappingList{}
	}
	return TranscoderMappingList{}.Unmarshal(result.Payload)
}

// GetTermList lists all possible values for the given controlled vocabulary.
func (c Client) GetTermList(endpoint string) TermList {
	return TermList{}
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
