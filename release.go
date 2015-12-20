package espsdk

import (
	"encoding/json"
	"log"
)

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

func (r Release) Marshal() ([]byte, error) { return json.MarshalIndent(r, "", "  ") }
func (r Release) ValidTypes() []string     { return []string{"Model", "Property"} }

func (r Release) PrettyPrint() string {
	prettyOutput, err := r.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return string(prettyOutput)
}

// Unmarshal attempts to deserialize the provided JSON payload into a
// Release object.
func (r Release) Unmarshal(payload []byte) Release {
	var release Release
	if err := json.Unmarshal(payload, &release); err != nil {
		log.Fatal(err)
	}
	return release
}

type ReleaseList []Release

func (rl ReleaseList) Marshal() ([]byte, error) {
	return json.MarshalIndent(rl, "", "  ")
}

func (rl ReleaseList) Unmarshal(payload []byte) ReleaseList {
	var releaseList ReleaseList
	if err := json.Unmarshal(payload, &releaseList); err != nil {
		log.Fatal(err)
	}
	return releaseList
}

func (rl ReleaseList) PrettyPrint() string {
	prettyOutput, err := rl.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return string(prettyOutput)
}
