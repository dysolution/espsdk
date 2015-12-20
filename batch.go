package espsdk

import (
	"encoding/json"
	"log"
	"time"
)

type SubmissionBatchType string

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

func (s SubmissionBatch) Marshal() ([]byte, error) { return json.MarshalIndent(s, "", "  ") }
func (b *SubmissionBatch) TypeIsValid() bool       { return batchTypeIsValid[b.SubmissionType] }
func (b SubmissionBatch) NameIsValid() bool        { return len(b.SubmissionName) > 0 }

func (b SubmissionBatch) ValidTypes() []string {
	keys := make([]string, len(batchTypeIsValid))
	i := 0
	for k := range batchTypeIsValid {
		keys[i] = k
		i++
	}
	return keys
}

func (b SubmissionBatch) Unmarshal(payload []byte) SubmissionBatch {
	var batch SubmissionBatch
	if err := json.Unmarshal(payload, &batch); err != nil {
		log.Fatal(err)
	}
	return batch
}
func (b SubmissionBatch) PrettyPrint() string {
	prettyOutput, err := b.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return string(prettyOutput)
}

type SubmissionBatchUpdate struct {
	SubmissionBatch SubmissionBatch `json:"submission_batch"`
}

func (s SubmissionBatchUpdate) Marshal() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

var batchTypeIsValid = map[string]bool{
	"getty_creative_video":  true,
	"getty_editorial_video": true,
	"getty_creative_still":  true,
	"getty_editorial_still": true,
	"istock_creative_video": true,
}

type BatchList []SubmissionBatch

func (bl BatchList) Marshal() ([]byte, error) {
	return json.MarshalIndent(bl, "", "  ")
}

func (bl BatchList) Unmarshal(payload []byte) BatchList {
	var batchList BatchList
	if err := json.Unmarshal(payload, &batchList); err != nil {
		log.Fatal(err)
	}
	return batchList
}

func (bl BatchList) PrettyPrint() string {
	prettyOutput, err := bl.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return string(prettyOutput)
}

type BatchListContainer struct {
	Items BatchList `json:"items"`
	Meta  struct {
		TotalItems int `json:"total_items"`
	} `json:"meta"`
}

func (blc BatchListContainer) Marshal() ([]byte, error) {
	return json.MarshalIndent(blc, "", "  ")
}

func (blc BatchListContainer) Unmarshal(payload []byte) BatchListContainer {
	var batchListContainer BatchListContainer
	if err := json.Unmarshal(payload, &batchListContainer); err != nil {
		log.Fatal(err)
	}
	return batchListContainer
}

func (blc BatchListContainer) PrettyPrint() string {
	prettyOutput, err := blc.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return string(prettyOutput)
}
