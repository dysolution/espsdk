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
	StorageURL           string   `json:"StorageURL,omitempty"`
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
