package espsdk

import (
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/dysolution/sleepwalker"
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
func (r Release) Index(client sleepwalker.RESTClient, batchID int) ReleaseList {
	desc := "Release.Index"
	r.SubmissionBatchID = batchID
	result, err := client.Get(r)
	if err != nil {
		result.Log().Error(desc)
		return ReleaseList{}
	}
	if result.StatusCode == 404 {
		result.Log().Error(desc)
		return ReleaseList{}
	}
	result.Log().Info(desc)
	releaseList, _ := ReleaseList{}.Unmarshal(result.Payload)
	return releaseList
}

// Path returns the path for the contribution.
// If the Contribution has no ID, Path returns the root for all
// contributions for the Batch (the Contribution Index).
func (r Release) Path() string {
	bid := r.SubmissionBatchID
	if r.ID == 0 {
		return fmt.Sprintf("%s/%d/releases", Batches, bid)
	}
	return fmt.Sprintf("%s/%d/releases/%d", Batches, bid, r.ID)
}

// ValidTypes are the Release types supported by ESP.
func (r Release) ValidTypes() []string { return []string{"Model", "Property"} }

// Marshal serializes the Release into a byte slice.
func (r Release) Marshal() ([]byte, error) { return sleepwalker.Marshal(r) }

// Unmarshal attempts to deserialize the provided JSON payload into a
// Release object.
func (r Release) Unmarshal(payload []byte) (*Release, error) {
	var release *Release
	err := json.Unmarshal(payload, &release)
	if err != nil {
		return release, err
	}
	return release, nil

}

// A ReleaseList is a slice of zero or more Releases.
type ReleaseList []Release

// Marshal serializes a ReleaseList into a byte slice.
func (rl ReleaseList) Marshal() ([]byte, error) {
	return sleepwalker.Marshal(rl)
}

// Unmarshal attempts to deserialize the provided JSON payload
// into the complete metadata returned by a request to the Index (GET all)
// API endpoint.
func (rl ReleaseList) Unmarshal(payload []byte) (ReleaseList, error) {
	var releaseList ReleaseList
	if err := json.Unmarshal(payload, &releaseList); err != nil {
		var errResponse interface{}
		json.Unmarshal(payload, &errResponse)
		Log.WithFields(logrus.Fields{
			"error":   err,
			"payload": errResponse,
		}).Error("ReleaseList.Unmarshal")
		return ReleaseList{}, err
	}
	return releaseList, nil
}
