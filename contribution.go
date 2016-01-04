package espsdk

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
)

// A Contribution is the metadata that represents a media asset from
// a contributor.
type Contribution struct {
	AdditionalFacialExpressions string            `json:"additional_facial_expressions,omitempty"`
	AlternateID                 string            `json:"alternate_id,omitempty"`
	CallForImage                bool              `json:"call_for_image,omitempty"`
	CameraShotDate              string            `json:"camera_shot_date,omitempty"`
	Caption                     string            `json:"caption,omitempty"`
	City                        string            `json:"city,omitempty"`
	CollectionCode              string            `json:"collection_code,omitempty"`
	ContentProviderName         string            `json:"content_provider_name,omitempty"`
	ContentProviderTitle        string            `json:"content_provider_title,omitempty"`
	ContentWarnings             string            `json:"content_warnings,omitempty"`
	Copyright                   string            `json:"copyright,omitempty"`
	CountryOfShoot              string            `json:"country_of_shoot,omitempty"`
	CreatedAt                   *time.Time        `json:"created_at,omitempty"`
	CreatedDate                 *time.Time        `json:"created_date,omitempty"`
	CreditLine                  string            `json:"credit_line,omitempty"`
	Errors                      map[string]string `json:"errors,omitempty"`
	EventID                     string            `json:"event_id,omitempty"`
	ExclusionRoutes             string            `json:"exclusion_routes,omitempty"`
	ExclusiveCoverage           bool              `json:"exclusive_coverage,omitempty"`
	ExternalFileLocation        string            `json:"external_file_location,omitempty"`
	ExtractedMetadataPresent    bool              `json:"extracted_metadata_present,omitempty"`
	FacialExpressions           string            `json:"facial_expressions,omitempty"`
	FileName                    string            `json:"file_name,omitempty"`
	FilePath                    string            `json:"file_path,omitempty"`
	FileUploaded                bool              `json:"file_uploaded,omitempty"`
	FinalBucket                 string            `json:"final_bucket,omitempty"`
	Headline                    string            `json:"headline,omitempty"`
	ID                          int               `json:"id,omitempty"`
	ImageHeight                 int               `json:"image_height,omitempty"`
	ImageWidth                  int               `json:"image_width,omitempty"`
	InactiveDate                string            `json:"inactive_date,omitempty"`
	InclusionRoutes             string            `json:"inclusion_routes,omitempty"`
	IptcCaptionWriter           string            `json:"iptc_caption_writer,omitempty"`
	IptcCategory                string            `json:"iptc_category,omitempty"`
	IptcSubjects                []string          `json:"iptc_subjects,omitempty"`
	Keywords                    []string          `json:"keywords,omitempty"`
	MasterID                    string            `json:"master_id,omitempty"`
	MediaType                   string            `json:"media_type,omitempty"`
	MetadataExtractionStartedAt *time.Time        `json:"metadata_extraction_started_at,omitempty"`
	MetadataExtractionTimeout   bool              `json:"metadata_extraction_timeout,omitempty"`
	MimeType                    string            `json:"mime_type,omitempty"`
	NumberOfPeople              string            `json:"string,omitempty"`
	PaidAssignment              bool              `json:"paid_assignment,omitempty"`
	PaidAssignmentID            string            `json:"paid_assignment_id,omitempty"`
	ParentSource                string            `json:"parent_source,omitempty"`
	PersonCompositions          string            `json:"person_compositions,omitempty"`
	Personalities               []string          `json:"personalities,omitempty"`
	PicscoutSuggestions         string            `json:"picscout_suggestions,omitempty"`
	ProvinceState               string            `json:"province_state,omitempty"`
	PublicistApprovalRequired   bool              `json:"publicist_approval_required,omitempty"`
	PublishedAt                 *time.Time        `json:"published_at,omitempty"`
	PulledReason                string            `json:"pulled_reason,omitempty"`
	Rank                        int               `json:"rank,omitempty"`
	ReadyForSale                bool              `json:"ready_for_sale,omitempty"`
	RecordedDate                string            `json:"recorded_date,omitempty"`
	RiskCategory                string            `json:"risk_category,omitempty"`
	ShotSpeed                   string            `json:"shot_speed,omitempty"`
	SiteDestination             []string          `json:"site_destination,omitempty"`
	Source                      string            `json:"source,omitempty"`
	SpecialInstructions         string            `json:"special_instructions,omitempty"`
	Status                      string            `json:"status,omitempty"`
	StorageURL                  string            `json:"storage_url,omitempty"`
	SubmissionBatchID           int               `json:"submission_batch_id,omitempty"`
	Submittable                 bool              `json:"submittable,omitempty"`
	SubmittedAt                 *time.Time        `json:"submitted_at,omitempty"`
	SubmittedToReviewAt         string            `json:"submitted_to_review_at,omitempty"`
	ThumbnailURL                string            `json:"thumbnail_url,omitempty"`
	UpdatedAt                   *time.Time        `json:"updated_at,omitempty"`
	UploadBucket                string            `json:"upload_bucket,omitempty"`
	UploadID                    string            `json:"upload_id,omitempty"`
	UserMetadataValid           bool              `json:"user_metadata_valid,omitempty"`
	VisualColor                 string            `json:"visual_color,omitempty"`
}

// Index requests a list of all Contributions associated with the specified
// Submission Batch.
func (c Contribution) Index(client *Client, batchID int) ContributionList {
	c.SubmissionBatchID = batchID
	return ContributionList{}.Unmarshal(client.get(c.Path()))
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
