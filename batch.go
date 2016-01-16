package espsdk

import (
	"encoding/json"
	"fmt"
	"time"
)

// A Batch is a container for Contributions of the same type and
// any Releases that may be associated with them.
type Batch struct {
	AssignmentID                     string     `json:"assignment_id,omitempty"`
	BatchTags                        []string   `json:"batch_tags,omitempty"`
	BriefID                          string     `json:"brief_id,omitempty"`
	ContributionsAwaitingReviewCount int        `json:"contributions_awaiting_review_count,omitempty"`
	ContributionsCount               int        `json:"contributions_count,omitempty"`
	CreatedAt                        *time.Time `json:"created_at,omitempty"`
	CreatedBy                        string     `json:"created_by,omitempty"`
	CreatorIstockUsername            string     `json:"creator_istock_username,omitempty"`
	EventID                          string     `json:"event_id,omitempty"`
	ID                               int        `json:"id,omitempty"`
	IsGetty                          bool       `json:"is_getty,omitempty"`
	IsIstock                         bool       `json:"is_istock,omitempty"`
	IstockExclusive                  bool       `json:"istock_exclusive,omitempty"`
	LastContributionSubmittedAt      *time.Time `json:"last_contribution_submitted_at,omitempty"`
	LastSubmittedAt                  *time.Time `json:"last_submitted_at,omitempty"`
	Note                             string     `json:"note,omitempty"`
	ProfileID                        int        `json:"profile_id,omitempty"`
	ReviewedContributionsCount       int        `json:"reviewed_contributions_count,omitempty"`
	RevisableContributionsCount      int        `json:"revisable_contributions_count,omitempty"`
	SaveExtractedMetadata            bool       `json:"save_extracted_metadata,omitempty"`
	Status                           string     `json:"status,omitempty"`
	SubmissionName                   string     `json:"submission_name,omitempty"`
	SubmissionType                   string     `json:"submission_type,omitempty"`
	SubmittedContributionsCount      int        `json:"submitted_contributions_count,omitempty"`
	UpdatedAt                        *time.Time `json:"updated_at,omitempty"`
	UserID                           string     `json:"user_id,omitempty"`
}

// NameIsValid provides validation for a proposed SubmissionName.
func (b Batch) NameIsValid() bool { return len(b.SubmissionName) > 0 }

// TypeIsValid reports whether a proposed type is valid for ESP.
func (b *Batch) TypeIsValid() bool { return batchTypeIsValid[b.SubmissionType] }

// ValidTypes are the BatchTypes supported by ESP.
func (b Batch) ValidTypes() []string {
	keys := make([]string, len(batchTypeIsValid))
	i := 0
	for k := range batchTypeIsValid {
		keys[i] = k
		i++
	}
	return keys
}

// Path returns the path for the Batch. If the Batch has no ID, Path returns
// the root for all Batches (the Batch Index).
func (b Batch) Path() string {
	if b.ID == 0 {
		return Batches
	}
	return fmt.Sprintf("%s/%d", Batches, b.ID)
}

// Marshal serializes the Batch into a byte slice.
func (b Batch) Marshal() ([]byte, error) { return indentedJSON(b) }

// Unmarshal serializes the Batch into a byte slice.
func (b Batch) Unmarshal(payload []byte) (*Batch, error) {
	var batch *Batch
	err := json.Unmarshal(payload, &batch)
	if err != nil {
		return batch, err
	}
	return batch, nil
}

// A BatchUpdate contains a Batch. This matches the
// structure of the JSON payload the API expects during a PUT.
type BatchUpdate struct {
	Batch Batch `json:"submission_batch"`
}

// Path returns the path of the batch being updated.
func (bu BatchUpdate) Path() string { return bu.Batch.Path() }

// Marshal serializes a BatchUpdate into a byte slice.
func (bu BatchUpdate) Marshal() ([]byte, error) { return indentedJSON(bu) }

var batchTypeIsValid = map[string]bool{
	"getty_creative_video":  true,
	"getty_editorial_video": true,
	"getty_creative_still":  true,
	"getty_editorial_still": true,
	"istock_creative_video": true,
}

// A BatchList matches the structure of the JSON payload returned
// by the GET (all) Batches API endpoint.
type BatchList struct {
	Items []Batch `json:"items,omitempty"`
	Meta  struct {
		TotalItems int `json:"total_items,omitempty"`
	} `json:"meta,omitempty"`
}

// Unmarshal attempts to deserialize the provided JSON payload
// into the complete metadata returned by a request to the Index (GET all)
// API endpoint.
func (bl BatchList) Unmarshal(payload []byte) BatchList {
	var dest BatchList
	err := json.Unmarshal(payload, &dest)
	if err != nil {
		log.Error(err)
	}
	return dest
}

// Last returns the most recently-created batch.
func (bl BatchList) Last() Batch {
	log.Debugf("getting last of %d batches", bl.Meta.TotalItems)
	return bl.Items[0]
}
