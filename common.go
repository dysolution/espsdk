/*
Package espsdk provides a Go interface to the JSON API of Getty Images'
Enterprise Submission Portal (ESP).

You will need a username and password that allows you
to log in to https://sandbox.espaws.com/ as well as
an API Key and and API Secret, both of which are accessible
at https://developer.gettyimages.com/apps/mykeys/.

	// Configure the SDK's client with your credentials.
	client := espsdk.Client{
		Credentials: espsdk.Credentials{
			APIKey:      "esp_api_key",
			APISecret:   "esp_api_secret",
			ESPUsername: "esp_username",
			ESPPassword: "esp_password",
		},
		UploadBucket: "oregon",
	}

	// Get a token based on those credentials.
	token := client.GetToken()

	data := espsdk.Batch{
		SubmissionName: "My Photos",
		SubmissionType: "getty_creative_still",
	}
	batch := espsdk.Batch{}.Create(&client, data)

	// Get a list of batches, which should include the one you just created.
	batches := espsdk.Batch{}.Index(&client)

You can proceed from there to add contributions:
    batchID := 81421  // iterate Batch{}.Index() to get these
    data := espsdk.Contribution{
        Headline: "My Photo Title",
        FileName: "IMG_9235.JPG",
    }
    espsdk.Contribution{}.Create(&client, batchID, data)

You can also add Releases to a batch:
    batchID := 81421
    data := espsdk.Release{
    ReleaseType: "Property",
        FileName: "IMG_1735.JPG",
        MimeType: "image/jpeg",
    }
    espsdk.Release{}.Create(&client, batchID, data)

Contributions, Releases, and Batches can be deleted as well:
    batchID := 81421
    releaseID := 172421  // iterate Release{}.Index() to get these
    espsdk.Release{ID: releaseID}.Delete(&client, batchID)

Each of the three main types has a consistent CRUD interface. Other API
endpoints are expressed either as simple GETs or endpoints that perform
auto-suggest against provided terms in order to match them to Getty's
controlled vocabularies, such as suggesting "Bob Newhart" for "Bob".

*/
package espsdk

import (
	"encoding/json"
	"time"

	log "github.com/Sirupsen/logrus"
)

func foo() {
}

// A Token is a string representation of an OAuth2 token. It grants a user
// access to the ESP API for a limited time.
type Token string

// A Serializable object can be serialized to a byte stream such as JSON.
type serializable interface {
	Marshal() ([]byte, error)
}

func prettyPrint(object interface{}) string {
	prettyOutput, err := Marshal(object)
	if err != nil {
		log.Fatal(err)
	}
	return string(prettyOutput)
}

// A FulfilledRequest provides an overview of a completed API request and
// its result, including timing and HTTP status codes.
type fulfilledRequest struct {
	*request
	*result
}

// Marshal serializes a FulfilledRequest into a byte stream.
func (r *fulfilledRequest) Marshal() ([]byte, error) { return json.Marshal(r) }

// Private

// A Response contains the HTTP status code and text that represent the API's
// response to a request.
type response struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"-"`
}

func start(s string) time.Time { return time.Now() }

func elapsed(s string, startTime time.Time) time.Duration {
	duration := time.Now().Sub(startTime)
	return duration
}

func indentedJSON(obj interface{}) ([]byte, error) {
	return json.MarshalIndent(obj, "", "\t")
}

func get(path string, token Token) []byte {
	request := newRequest("GET", path, token, nil)
	result := Client{}.performRequest(request)
	if result.Err != nil {
		log.Fatal(result.Err)
	}
	stats, err := result.Marshal()
	if err != nil {
		log.Fatal(result.Err)
	}
	log.Info(string(stats))
	log.Debugf("%s\n", result.Payload)
	return result.Payload
}
