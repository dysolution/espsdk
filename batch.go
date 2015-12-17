package espapi

import "encoding/json"

type SubmissionBatchType string

type SubmissionBatch struct {
	SubmissionName        string `json:"submission_name,omitempty"`
	SubmissionType        string `json:"submission_type,omitempty"`
	Note                  string `json:"note,omitempty"`
	SaveExtractedMetadata bool   `json:"save_extracted_metadata,omitempty"`
	AssignmentId          string `json:"assignment_id,omitempty"`
	BriefId               string `json:"brief_id,omitempty"`
	EventId               string `json:"event_id,omitempty"`
}

func (s SubmissionBatch) Marshal() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

func (b *SubmissionBatch) TypeIsValid() bool {
	return batchTypeIsValid[b.SubmissionType]
}

func (b SubmissionBatch) NameIsValid() bool {
	return len(b.SubmissionName) > 0
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

func BatchTypes() []string {
	keys := make([]string, len(batchTypeIsValid))
	i := 0
	for k := range batchTypeIsValid {
		keys[i] = k
		i++
	}
	return keys
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
