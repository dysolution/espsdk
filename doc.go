/*
Package espsdk provides a Go interface to the JSON API of Getty Images'
Enterprise Submission Portal (ESP).

You will need a username and password that allows you
to log in to https://sandbox.espaws.com/ as well as
an API Key and and API Secret, both of which are accessible
at https://developer.gettyimages.com/apps/mykeys/.

	// Configure the SDK's client with your credentials.
	client := espsdk.GetClient(
	  "esp_api_key",
		"esp_api_secret",
		"esp_username",
		"esp_password",
		"oregon",
		)

	// Get a token based on those credentials.
	token := client.GetToken()

	data := espsdk.Batch{
		SubmissionName: "My Photos",
		SubmissionType: "getty_creative_still",
	}
	batch := client.Create(data.Path(), data)

	// Get a list of batches, which should include the one you just created.
	batches := client.Index(espsdk.Batches)

You can proceed from there to add contributions:
    batchID := 81421  // iterate "batches" above to get these
    data := espsdk.Contribution{
        Headline: "My Photo Title",
        FileName: "IMG_9235.JPG",
    }
		client.Create(data.Path(), data)

You can also add Releases to a batch:
    batchID := 81421
    data := espsdk.Release{
    ReleaseType: "Property",
        FileName: "IMG_1735.JPG",
        MimeType: "image/jpeg",
    }
    espsdk.Release{}.Create(&client, batchID, data)
		client.Create(data.Path(), data)

Contributions, Releases, and Batches can be deleted as well:
    batchID := 81421
    releaseID := 172421  // iterate Release{}.Index() to get these
		release := espsdk.Release{ID: 172421, SubmissionBatchID: 81421}
		client.Delete(release.Path())

Each of the three main types has a consistent CRUD interface. Other API
endpoints are expressed either as simple GETs or endpoints that perform
auto-suggest against provided terms in order to match them to Getty's
controlled vocabularies, such as suggesting "Bob Newhart" for "Bob".

*/
package espsdk
