package espsdk

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

// A Release is the metadata that represents a legal agreement for
// property owners or models.
type Release struct {
	ExternalFileLocation string   `json:"external_file_location,omitempty"`
	FileName             string   `json:"file_name,omitempty"`
	FilePath             string   `json:"file_path,omitempty"`
	ID                   int      `json:"id,omitempty"`
	MimeType             string   `json:"mime_type,omitempty"`
	ModelDateOfBirth     string   `json:"model_date_of_birth,omitempty"`
	ModelEthnicities     []string `json:"model_ethnicities,omitempty"`
	ModelGender          string   `json:"model_gender,omitempty"`
	ReleaseType          string   `json:"release_type,omitempty"`
	StorageURL           string   `json:"storage_url,omitempty"`
	SubmissionBatchID    int      `json:"submission_batch_id,omitempty"`
	UploadID             int      `json:"upload_id,omitempty"`
}

// Index requests a list of all Releases associated with the specified
// Submission Batch.
func (r Release) Index(client *Client, batchID int) ReleaseList {
	return ReleaseList{}.Unmarshal(client.get(ReleasePath(batchID, 0)))
}

// Get requests the metadata for a specific Release.
func (r Release) Get(client *Client, batchID int) Release {
	return r.Unmarshal(client.get(ReleasePath(batchID, r.ID)))
}

// Create adds a new Release to a Submission Batch.
func (r Release) Create(client *Client, batchID int, releaseData Release) Release {
	return r.Unmarshal(client.post(releaseData, ReleasePath(batchID, r.ID)))
}

// Delete destroys a specific Release.
func (r Release) Delete(client *Client, batchID int) {
	client._delete(ReleasePath(batchID, r.ID))
}

// Marshal serializes a Release into a byte slice.
func (r Release) Marshal() ([]byte, error) { return indentedJSON(r) }

// ValidTypes are the Release types supported by ESP.
func (r Release) ValidTypes() []string { return []string{"Model", "Property"} }

// PrettyPrint returns a human-readable serialized JSON representation of
// the provided object.
func (r Release) PrettyPrint() string { return prettyPrint(r) }

// Unmarshal attempts to deserialize the provided JSON payload into a
// Release object.
func (r Release) Unmarshal(payload []byte) Release {
	var release Release
	if err := json.Unmarshal(payload, &release); err != nil {
		log.Fatal(err)
	}
	return release
}

// A ReleaseList is a slice of zero or more Releases.
type ReleaseList []Release

// Marshal serializes a ReleaseList into a byte slice.
func (rl ReleaseList) Marshal() ([]byte, error) { return indentedJSON(rl) }

// Unmarshal attempts to deserialize the provided JSON payload
// into the complete metadata returned by a request to the Index (GET all)
// API endpoint.
func (rl ReleaseList) Unmarshal(payload []byte) ReleaseList {
	var releaseList ReleaseList
	if err := json.Unmarshal(payload, &releaseList); err != nil {
		log.Fatal(err)
	}
	return releaseList
}

// PrettyPrint returns a human-readable serialized JSON representation of
// the provided object.
func (rl ReleaseList) PrettyPrint() string { return prettyPrint(rl) }
