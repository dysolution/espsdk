package espsdk

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

// A Contribution is the metadata that represents a media asset from
// a contributor.
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
	UploadID             string   `json:"upload_id,omitempty"`
}

// Index requests a list of all Contributions associated with the specified
// Submission Batch.
func (c Contribution) Index(client *Client, batchID int) ContributionList {
	return ContributionList{}.Unmarshal(client.get(ContributionPath(batchID, 0)))
}

// Path returns the path for the contribution.
// If the Contribution has no ID, Path returns the root for all
// contributions for the Batch (the Contribution Index).
func (c Contribution) Path() string {
	bid := c.SubmissionBatchID
	if c.ID == 0 {
		return fmt.Sprintf("%s/%d/contributions", Batches, bid)
	}
	return fmt.Sprintf("%s/%d/contributions/%d", Batches, bid, c.ID)
}

// A ContributionUpdate contains a Contribution. This matches the
// structure of the JSON payload the API expects during a PUT.
type ContributionUpdate struct {
	Contribution Contribution `json:"contribution"`
}

// Marshal serializes a ContributionUpdate into a byte slice.
func (c ContributionUpdate) Marshal() ([]byte, error) { return indentedJSON(c) }

// Path returns the path of the contribution being updated.
func (c ContributionUpdate) Path() string { return c.Contribution.Path() }

// A ContributionList is a slice of zero or more Contributions.
type ContributionList []Contribution

// Unmarshal attempts to deserialize the provided JSON payload
// into the complete metadata returned by a request to the Index (GET all)
// API endpoint.
func (cl ContributionList) Unmarshal(payload []byte) ContributionList {
	var contributionList ContributionList
	if err := json.Unmarshal(payload, &contributionList); err != nil {
		log.Fatal(err)
	}
	return contributionList
}

// PrettyPrint returns a human-readable serialized JSON representation of
// the provided object.
func (cl ContributionList) PrettyPrint() string { return prettyPrint(cl) }
