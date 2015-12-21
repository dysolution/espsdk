package espsdk

func GetKeywords(client *Client) []byte           { return client.get(Keywords) }
func GetPersonalities(client *Client) []byte      { return client.get(Personalities) }
func GetControlledValues(client *Client) []byte   { return client.get(ControlledValues) }
func GetTranscoderMappings(client *Client) []byte { return client.get(TranscoderMappings) }
func GetCompositions(client *Client) []byte       { return client.get(Compositions) }
func GetExpressions(client *Client) []byte        { return client.get(Expressions) }
func GetNumberOfPeople(client *Client) []byte     { return client.get(NumberOfPeople) }
