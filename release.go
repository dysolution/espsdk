package espsdk

import (
	"encoding/json"
)

type Release struct {
	SubmissionBatchId    string   `json:"submission_batch_id"`
	FileName             string   `json:"file_name"`
	FilePath             string   `json:"file_path"`
	ExternalFileLocation string   `json:"external_file_location"`
	ReleaseType          string   `json:"release_type"`
	ModelDateOfBirth     string   `json:"model_date_of_birth"`
	ModelEthnicities     []string `json:"model_ethnicities"`
	ModelGender          string   `json:"model_gender"`
}

func (r Release) Marshal() ([]byte, error) { return json.MarshalIndent(r, "", "  ") }
func (r Release) ValidTypes() []string     { return []string{"Model", "Property"} }
