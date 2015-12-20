package espsdk

import (
	"encoding/json"
	"log"
)

type Contribution struct {
	AlternateID          string   `json:"alternate_id,omitempty"`
	CameraShotDate       string   `json:"camera_shot_date,omitempty"`
	CollectionCode       string   `json:"collection_code,omitempty"`
	ContentProviderName  string   `json:"content_provider_name,omitempty"`
	ContentProviderTitle string   `json:"content_provider_title,omitempty"`
	CountryOfShoot       string   `json:"country_of_shoot,omitempty"`
	CreditLine           string   `json:"credit_line,omitempty"`
	ExternalFileLocation string   `json:"external_file_location,omitempty"`
	FileName             string   `json:"file_name,omitempty"`
	FinalBucket          string   `json:"final_bucket,omitempty"`
	FilePath             string   `json:"file_path,omitempty"`
	Headline             string   `json:"headline,omitempty"`
	ID                   int      `json:"id,omitempty"`
	IptcCategory         string   `json:"iptc_category,omitempty"`
	IptcSubjects         []string `json:"iptc_subjects,omitempty"`
	MasterID             string   `json:"master_id,omitempty"`
	MimeType             string   `json:"mime_type,omitempty"`
	ParentSource         string   `json:"parent_source,omitempty"`
	RecordedDate         string   `json:"recorded_date,omitempty"`
	RiskCategory         string   `json:"risk_category,omitempty"`
	ShotSpeed            string   `json:"shot_speed,omitempty"`
	SiteDestination      []string `json:"site_destination,omitempty"`
	Source               string   `json:"source,omitempty"`
	SubmissionBatchID    int      `json:"submission_batch_id,omitempty"`
	SubmittedToReviewAt  string   `json:"submitted_to_review_at,omitempty"`
	UploadBucket         string   `json:"upload_bucket,omitempty"`
	UploadId             string   `json:"upload_id,omitempty"`
}

func (c Contribution) Marshal() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

func (c Contribution) PrettyPrint() string {
	prettyOutput, err := c.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return string(prettyOutput)
}

type ContributionUpdate struct {
	Contribution Contribution `json:"contribution"`
}

type ContributionDeleteParams struct {
	SubmissionBatchID string `json:"submission_batch_id,omitempty"`
	ContributionID    string `json:"id,omitempty"`
}

func (c ContributionDeleteParams) Marshal() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

func (c ContributionUpdate) Marshal() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}
