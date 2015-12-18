package espsdk

import (
	"encoding/json"
)

type Contribution struct {
	CameraShotDate       string   `json:"camera_shot_date"`
	CollectionCode       string   `json:"collection_code"`
	ContentProviderName  string   `json:"content_provider_name"`
	ContentProviderTitle string   `json:"content_provider_title"`
	CountryOfShoot       string   `json:"country_of_shoot"`
	CreditLine           string   `json:"credit_line"`
	ExternalFileLocation string   `json:"external_file_location"`
	FileName             string   `json:"file_name"`
	FilePath             string   `json:"file_path"`
	Headline             string   `json:"headline"`
	IptcCategory         string   `json:"iptc_category"`
	IptcSubjects         []string `json:"iptc_subjects"`
	MimeType             string   `json:"mime_type"`
	ParentSource         string   `json:"parent_source"`
	RecordedDate         string   `json:"recorded_date"`
	RiskCategory         string   `json:"risk_category"`
	ShotSpeed            string   `json:"shot_speed"`
	SiteDestination      []string `json:"site_destination"`
	Source               string   `json:"source"`
	SubmittedToReviewAt  string   `json:"submitted_to_review_at"`
	UploadBucket         string   `json:"upload_bucket"`
	UploadId             string   `json:"upload_id"`
}

func (c Contribution) Marshal() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}
