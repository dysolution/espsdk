package espsdk

import (
	"encoding/json"
	"time"
)

type SubmissionBatchType string

type SubmissionBatch struct {
	AssignmentId                     string     `json:"assignment_id,omitempty"`
	BatchTags                        []string   `json:"batch_tags,omitempty"`
	BriefId                          string     `json:"brief_id,omitempty"`
	ContributionsAwaitingReviewCount int        `json:"contributions_awaiting_review_count"`
	ContributionsCount               int        `json:"contributions_count"`
	CreatedAt                        time.Time  `json:"created_at"`
	CreatedBy                        string     `json:"created_by"`
	CreatorIstockUsername            string     `json:"creator_istock_username"`
	EventId                          string     `json:"event_id,omitempty"`
	Id                               int        `json:"id"`
	IsGetty                          bool       `json:"is_getty"`
	IsIstock                         bool       `json:"is_istock"`
	IstockExclusive                  bool       `json:"istock_exclusive"`
	LastContributionSubmittedAt      *time.Time `json:"last_contribution_submitted_at,omitempty"`
	LastSubmittedAt                  *time.Time `json:"last_submitted_at,omitempty"`
	Note                             string     `json:"note,omitempty"`
	ProfileId                        int        `json:"profile_id"`
	ReviewedContributionsCount       int        `json:"reviewed_contributions_count"`
	RevisableContributionsCount      int        `json:"revisable_contributions_count"`
	SaveExtractedMetadata            bool       `json:"save_extracted_metadata,omitempty"`
	Status                           string     `json:"status,omitempty"`
	SubmissionName                   string     `json:"submission_name,omitempty"`
	SubmissionType                   string     `json:"submission_type,omitempty"`
	SubmittedContributionsCount      int        `json:"submitted_contributions_count,omitempty"`
	UpdatedAt                        time.Time  `json:"updated_at"`
	UserId                           string     `json:"user_id"`
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

type SubmissionBatchUpdate struct {
	SubmissionBatch SubmissionBatchChanges `json:"submission_batch"`
}

func (s SubmissionBatchUpdate) Marshal() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

type SubmissionBatchChanges struct {
	SubmissionName string `json:"submission_name,omitempty"`
	Note           string `json:"note,omitempty"`
}

type ExtantSubmissionBatch struct {
	Id                               int
	submissionName                   string
	submissionType                   string
	userId                           string
	profileId                        int
	createdAt                        string
	updatedAt                        string
	status                           string
	note                             string
	createdBy                        string
	lastContributionSubmittedAt      string
	contributionsCount               int
	submittedContributionsCount      int
	istockExclusive                  bool
	contributionsAwaitingReviewCount int
	reviewedContributionsCount       int
	revisableContributionsCount      int
	lastSubmittedAt                  string
	isGetty                          bool
	isIstock                         bool
	totalProfileNotes                int
	saveExtractedMetadata            bool
	assignmentId                     string
	briefId                          string
	eventId                          string
}

var batchTypeIsValid = map[string]bool{
	"getty_creative_video":  true,
	"getty_editorial_video": true,
	"getty_creative_still":  true,
	"getty_editorial_still": true,
	"istock_creative_video": true,
}
