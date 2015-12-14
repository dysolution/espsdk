package espapi

import (
	"encoding/json"
)

type Contribution struct {
	FileName             string `json:"file_name"`
	FilePath             string `json:"file_path"`
	SubmittedToReviewAt  string `json:"submitted_to_review_at"`
	UploadBucket         string `json:"upload_bucket"`
	ExternalFileLocation string `json:"external_file_location"`
	UploadId             string `json:"upload_id"`
	MimeType             string `json:"mime_type"`
}

func (c Contribution) Marshal() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}
