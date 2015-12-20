package espsdk

import (
	"encoding/json"
	"log"
	"time"
)

// A SubmissionBatch is a container for Contributions of the same type and
// any Releases that may be associated with them.
type SubmissionBatch struct {
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

// Marshal serializes a SubmissionBatch into a byte slice.
func (b SubmissionBatch) Marshal() ([]byte, error) { return indentedJSON(b) }

// NameIsValid provides validation for a proposed SubmissionName.
func (b SubmissionBatch) NameIsValid() bool { return len(b.SubmissionName) > 0 }

// TypeIsValid reports whether a proposed type is valid for ESP.
func (b *SubmissionBatch) TypeIsValid() bool { return batchTypeIsValid[b.SubmissionType] }

// ValidTypes are the SubmissionBatchTypes supported by ESP.
func (b SubmissionBatch) ValidTypes() []string {
	keys := make([]string, len(batchTypeIsValid))
	i := 0
	for k := range batchTypeIsValid {
		keys[i] = k
		i++
	}
	return keys
}

// Unmarshal attempts to deserialize the provided JSON payload
// into a SubmissionBatch object.
func (b SubmissionBatch) Unmarshal(payload []byte) SubmissionBatch {
	var batch SubmissionBatch
	if err := json.Unmarshal(payload, &batch); err != nil {
		log.Fatal(err)
	}
	return batch
}

// PrettyPrint returns a human-readable serialized JSON representation of
// the provided object.
func (b SubmissionBatch) PrettyPrint() string { return prettyPrint(b) }

// A SubmissionBatchUpdate contains a SubmissionBatch. This matches the
// structure of the JSON payload the API expects during a PUT.
type SubmissionBatchUpdate struct {
	SubmissionBatch SubmissionBatch `json:"submission_batch"`
}

// Marshal serializes a SubmissionBatchUpdate into a byte slice.
func (s SubmissionBatchUpdate) Marshal() ([]byte, error) { return indentedJSON(s) }

var batchTypeIsValid = map[string]bool{
	"getty_creative_video":  true,
	"getty_editorial_video": true,
	"getty_creative_still":  true,
	"getty_editorial_still": true,
	"istock_creative_video": true,
}

// A BatchList is a slice of zero or more SubmissionBatches.
type BatchList []SubmissionBatch

// Marshal serializes a BatchList into a byte slice.
func (bl BatchList) Marshal() ([]byte, error) { return indentedJSON(bl) }

func (bl BatchList) unmarshal(payload []byte) BatchList {
	var batchList BatchList
	if err := json.Unmarshal(payload, &batchList); err != nil {
		log.Fatal(err)
	}
	return batchList
}

func (bl BatchList) prettyPrint() string { return prettyPrint(bl) }

// A BatchListContainer matches the structure of the JSON payload returned
// by the GET (all) SubmissionBatches API endpoint.
type BatchListContainer struct {
	Items BatchList `json:"items"`
	Meta  struct {
		TotalItems int `json:"total_items"`
	} `json:"meta"`
}

// Marshal serializes a BatchListContainer into a byte slice.
func (blc BatchListContainer) Marshal() ([]byte, error) { return indentedJSON(blc) }

// Unmarshal attempts to deserialize the provided JSON payload
// into the complete metadata returned by a request to the Index (GET all)
// API endpoint.
func (blc BatchListContainer) Unmarshal(payload []byte) BatchListContainer {
	var batchListContainer BatchListContainer
	if err := json.Unmarshal(payload, &batchListContainer); err != nil {
		log.Fatal(err)
	}
	return batchListContainer
}

// PrettyPrint returns a human-readable serialized JSON representation of
// the provided object.
func (blc BatchListContainer) PrettyPrint() string { return prettyPrint(blc) }
