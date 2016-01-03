package espsdk

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
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

// Index requests a list of all Batches owned by the user.
func (b Batch) Index(client *Client) BatchListContainer {
	return BatchListContainer{}.Unmarshal(client.get(BatchPath(&b)))
}

// Update changes metadata for an existing Batch.
func (b Batch) Update(client *Client, updatedData BatchUpdate) DeserializedObject {
	return Unmarshal(client.put(updatedData, BatchPath(&b)))
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

// A BatchUpdate contains a Batch. This matches the
// structure of the JSON payload the API expects during a PUT.
type BatchUpdate struct {
	Batch Batch `json:"submission_batch"`
}

// Marshal serializes a BatchUpdate into a byte slice.
func (s BatchUpdate) Marshal() ([]byte, error) { return indentedJSON(s) }

var batchTypeIsValid = map[string]bool{
	"getty_creative_video":  true,
	"getty_editorial_video": true,
	"getty_creative_still":  true,
	"getty_editorial_still": true,
	"istock_creative_video": true,
}

// A BatchList is a slice of zero or more Batches.
type BatchList []Batch

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
// by the GET (all) Batches API endpoint.
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
