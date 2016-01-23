package espsdk

import (
	"errors"

	"github.com/dysolution/sleepwalker"
)

type Client struct {
	*sleepwalker.Client
}

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
func (c Client) GetTranscoderMappings() TranscoderMappingList {
	return TranscoderMappingList{}.Unmarshal([]byte("not implemented"))
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
