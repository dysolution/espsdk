package espsdk

import "encoding/json"

// Deserialize attempts to deserialize the provided JSON payload
// into an object.
func Deserialize(payload []byte, dest *interface{}) *interface{} {
	err := json.Unmarshal(payload, &dest)
	if err != nil {
		Log.Error(err)
	}
	return dest
}

// Create uses the provided path and data to ask the API to create a new
// object and returns the deserialized response.
// func Create(object sleepwalker.RESTObject, client *Client) interface{} {
// 	marshaledObject := client.deprecatedPost(object)
// 	return Unmarshal(marshaledObject)
// }
