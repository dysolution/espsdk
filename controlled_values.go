package espsdk

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

// GetNumberOfPeople lists all possible values for Number of People.
func GetNumberOfPeople(client *Client) []byte { return client.get(NumberOfPeople) }
