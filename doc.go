/*
Package espsdk provides a Go interface to the JSON API of Getty Images'
Enterprise Submission Portal (ESP).

You will need a username and password that allows you
to log in to https://sandbox.espaws.com/ as well as
an API Key and and API Secret, both of which are accessible
at https://developer.gettyimages.com/apps/mykeys/.

First, configure the SDK's client with your credentials:

	client := espsdk.GetClient(
		"esp_api_key",
		"esp_api_secret",
		"esp_username",
		"esp_password",
		"oregon",
		)

The client creates and sends a token along with each request. If you'd like
to save and cache it, you can call GetToken directly:

	token := client.GetToken()

Media assets, such as photos and videos (represented as Contributions) must be
uploaded into Submission Batches. You can use the client to create a Submission Batch:

	data := espsdk.Batch{
		SubmissionName: "My Photos",
		SubmissionType: "getty_creative_still",
	}
	batch := client.Create(data.Path(), data)

The list (index) of your Batches now includes the Batch you just created:

	batches := client.Index(espsdk.Batches)

You can proceed from there to add contributions:
    batchID := 81421  // iterate "batches" above to get these
    data := espsdk.Contribution{
        Headline: "My Photo Title",
        FileName: "IMG_9235.JPG",
    }
    client.Create(data.Path(), data)

The list (index) of Contributions for the Batch now includes the one you
just created:

    batchID := 81421
    contributions := espsdk.Contribution{}.Index(&client, batchID)

You can also add Releases to a batch:
    batchID := 81421
    data := espsdk.Release{
    ReleaseType: "Property",
        FileName: "IMG_1735.JPG",
        MimeType: "image/jpeg",
    }
    client.Create(data.Path(), data)

The list (index) of Releases for the Batch now includes the one you
just created:

    batchID := 81421
    releases := espsdk.Release{}.Index(&client, batchID)

Contributions, Releases, and Batches can be deleted as well:
    release := espsdk.Release{ID: 172421, SubmissionBatchID: 81421}
    client.Delete(release.Path())

Each of the three main types has a consistent CRUD interface. Other API
endpoints are expressed either as simple GETs or endpoints that perform
auto-suggest against provided terms in order to match them to Getty's
controlled vocabularies, such as suggesting "Bob Newhart" for "Bob".

*/
package espsdk
