package espsdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dysolution/sleepwalker"
)

// A Contribution is the metadata that represents a media asset from
// a contributor.
type Contribution struct {
	AdditionalFacialExpressions []map[string]interface{} `json:"additional_facial_expressions,omitempty"`
	AlternateID                 string                   `json:"alternate_id,omitempty"`
	CallForImage                bool                     `json:"call_for_image,omitempty"`
	CameraShotDate              string                   `json:"camera_shot_date,omitempty"`
	Caption                     string                   `json:"caption,omitempty"`
	City                        string                   `json:"city,omitempty"`
	CollectionCode              string                   `json:"collection_code,omitempty"`
	ContentProviderName         string                   `json:"content_provider_name,omitempty"`
	ContentProviderTitle        string                   `json:"content_provider_title,omitempty"`
	ContentWarnings             string                   `json:"content_warnings,omitempty"`
	Copyright                   string                   `json:"copyright,omitempty"`
	CountryOfShoot              string                   `json:"country_of_shoot,omitempty"`
	CreatedAt                   *time.Time               `json:"created_at,omitempty"`
	CreatedDate                 *time.Time               `json:"created_date,omitempty"`
	CreditLine                  string                   `json:"credit_line,omitempty"`
	DSAAlternateIds             map[string]string        `json:"dsa_alternate_ids,omitempty"`
	Errors                      interface{}              `json:"errors,omitempty"`
	EventID                     string                   `json:"event_id,omitempty"`
	ExclusionRoutes             string                   `json:"exclusion_routes,omitempty"`
	ExclusiveCoverage           bool                     `json:"exclusive_coverage,omitempty"`
	ExternalFileLocation        string                   `json:"external_file_location,omitempty"`
	ExtractedMetadataPresent    bool                     `json:"extracted_metadata_present,omitempty"`
	FacialExpressions           []TermItem               `json:"facial_expressions,omitempty"`
	FileName                    string                   `json:"file_name,omitempty"`
	FilePath                    string                   `json:"file_path,omitempty"`
	FileUploaded                bool                     `json:"file_uploaded,omitempty"`
	FinalBucket                 string                   `json:"final_bucket,omitempty"`
	Headline                    string                   `json:"headline,omitempty"`
	ID                          string                   `json:"id,omitempty"`
	IPTCCaptionWriter           string                   `json:"iptc_caption_writer,omitempty"`
	IPTCCategory                string                   `json:"iptc_category,omitempty"`
	IPTCSubjects                []string                 `json:"iptc_subjects,omitempty"`
	ImageHeight                 int                      `json:"image_height,omitempty"`
	ImageWidth                  int                      `json:"image_width,omitempty"`
	InactiveDate                *time.Time               `json:"inactive_date,omitempty"`
	InclusionRoutes             interface{}              `json:"inclusion_routes,omitempty"`
	Keywords                    []Keyword                `json:"keywords,omitempty"`
	MasterID                    string                   `json:"master_id,omitempty"`
	MediaType                   string                   `json:"media_type,omitempty"`
	MetadataExtractionStartedAt *time.Time               `json:"metadata_extraction_started_at,omitempty"`
	MetadataExtractionTimeout   bool                     `json:"metadata_extraction_timeout,omitempty"`
	MimeType                    string                   `json:"mime_type,omitempty"`
	// NumberOfPeople              TermItemInt              `json:"number_of_people,omitempty"`
	PaidAssignment            bool        `json:"paid_assignment,omitempty"`
	PaidAssignmentID          string      `json:"paid_assignment_id,omitempty"`
	ParentSource              string      `json:"parent_source,omitempty"`
	PersonCompositions        []TermItem  `json:"person_compositions,omitempty"`
	Personalities             []Keyword   `json:"personalities,omitempty"`
	PicscoutSuggestions       interface{} `json:"picscout_suggestions,omitempty"`
	ProvinceState             string      `json:"province_state,omitempty"`
	PublicistApprovalRequired bool        `json:"publicist_approval_required,omitempty"`
	PublishedAt               *time.Time  `json:"published_at,omitempty"`
	PulledReason              string      `json:"pulled_reason,omitempty"`
	Rank                      int         `json:"rank,omitempty"`
	ReadyForSale              bool        `json:"ready_for_sale,omitempty"`
	RecordedDate              string      `json:"recorded_date,omitempty"`
	RiskCategory              string      `json:"risk_category,omitempty"`
	ShotSpeed                 string      `json:"shot_speed,omitempty"`
	SiteDestination           []string    `json:"site_destination,omitempty"`
	Source                    string      `json:"source,omitempty"`
	SpecialInstructions       string      `json:"special_instructions,omitempty"`
	Status                    string      `json:"status,omitempty"`
	StorageURL                string      `json:"storage_url,omitempty"`
	SubmissionBatchID         string      `json:"submission_batch_id,omitempty"`
	Submittable               bool        `json:"submittable,omitempty"`
	SubmittedAt               *time.Time  `json:"submitted_at,omitempty"`
	SubmittedToReviewAt       string      `json:"submitted_to_review_at,omitempty"`
	ThumbnailURL              string      `json:"thumbnail_url,omitempty"`
	UpdatedAt                 *time.Time  `json:"updated_at,omitempty"`
	UploadBucket              string      `json:"upload_bucket,omitempty"`
	UploadID                  string      `json:"upload_id,omitempty"`
	UserMetadataValid         bool        `json:"user_metadata_valid,omitempty"`
	VisualColor               string      `json:"visual_color,omitempty"`
}

// Submit requests that the contribution be submitted for review and
// publication.
func (c Contribution) Submit(client sleepwalker.RESTClient) (sleepwalker.Result, error) {
	desc := "Contribution.Submit"
	result, err := client.Put(c, c.Path()+"/submit")
	if err != nil {
		result.Log().Error(desc)
		return result, err
	}
	if result.StatusCode == 404 {
		result.Log().Error(desc)
		return result, err
	}
	result.Log().Info(desc)
	return result, nil
}

// CreateAndSubmit creates a Contribution and submits it for review/publication.
func (c Contribution) CreateAndSubmit(client sleepwalker.RESTClient) (sleepwalker.Result, error) {
	desc := "Contribution.CreateAndSubmit"
	result, err := client.Create(c)
	if err != nil {
		result.Log().Error(desc)
		return result, err
	}
	result.Log().Debug(desc)

	var savedContribution Contribution
	json.Unmarshal(result.Payload, &savedContribution)
	result, err = client.Put(savedContribution, savedContribution.Path()+"/submit")
	if err != nil {
		result.Log().Error(desc)
		return result, err
	}
	result.Log().Info(desc)
	return result, nil
}

// Index requests a list of all Contributions associated with the specified
// Submission Batch.
func (c Contribution) Index(client sleepwalker.RESTClient, batchID string) ContributionList {
	desc := "Contribution.Index"
	c.SubmissionBatchID = batchID
	result, err := client.Get(c)
	if err != nil {
		result.Log().Error(desc)
		return ContributionList{}
	}
	if result.StatusCode == 404 {
		result.Log().Error(desc)
		return ContributionList{}
	}
	result.Log().Info(desc)
	contributionList, _ := ContributionList{}.Unmarshal(result.Payload)
	return contributionList
}

// Path returns the path for the contribution.
// If the Contribution has no ID, Path returns the root for all
// contributions for the Batch (the Contribution Index).
func (c Contribution) Path() string {
	bid := c.SubmissionBatchID
	p := Endpoints.Batches
	if c.ID == "" {
		return fmt.Sprintf("%s/%s/contributions", p, bid)
	}
	return fmt.Sprintf("%s/%s/contributions/%s", p, bid, c.ID)
}

// Marshal serializes the Contribution into a byte slice.
func (c Contribution) Marshal() ([]byte, error) {
	return sleepwalker.Marshal(c)
}

// Unmarshal attempts to deserialize the provided JSON payload into a
// Contribution object.
func (c Contribution) Unmarshal(payload []byte) (*Contribution, error) {
	var contribution *Contribution
	err := json.Unmarshal(payload, &contribution)
	if err != nil {
		return contribution, err
	}
	return contribution, nil

}

// A ContributionUpdate contains a Contribution. This matches the
// structure of the JSON payload the API expects during a PUT.
type ContributionUpdate struct {
	Contribution Contribution `json:"contribution"`
}

// Marshal serializes a ContributionUpdate into a byte slice.
func (c ContributionUpdate) Marshal() ([]byte, error) {
	return sleepwalker.Marshal(c)
}

// Path returns the path of the contribution being updated.
func (c ContributionUpdate) Path() string { return c.Contribution.Path() }

// A ContributionList is a slice of zero or more Contributions.
type ContributionList []Contribution

// Unmarshal attempts to deserialize the provided JSON payload
// into the complete metadata returned by a request to the Index (GET all)
// API endpoint.
func (cl ContributionList) Unmarshal(payload []byte) (ContributionList, error) {
	var contributionList ContributionList
	if err := json.Unmarshal(payload, &contributionList); err != nil {
		var errResponse interface{}
		json.Unmarshal(payload, &errResponse)
		output, _ := json.MarshalIndent(errResponse, "", "    ")
		Log.WithFields(map[string]interface{}{
			"error":   err,
			"payload": string(output),
		}).Error("ContributionList.Unmarshal")
		return ContributionList{}, err
	}
	return contributionList, nil
}

// Last returns the most recently-created Contribution.
func (cl ContributionList) Last() (Contribution, error) {
	desc := "ContributionList.Last"
	Log.WithFields(map[string]interface{}{
		"count":  len(cl),
		"object": "contribution",
	}).Debugf(desc)
	if len(cl) == 0 {
		return Contribution{}, errors.New("no contributions")
	}
	return cl[len(cl)-1], nil
	// return cl[0], nil
}
