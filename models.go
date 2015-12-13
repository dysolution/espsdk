package espapi

type SubmissionBatchType string

func BatchTypes() []string {
	keys := make([]string, len(batchTypeIsValid))
	i := 0
	for k := range batchTypeIsValid {
		keys[i] = k
		i++
	}
	return keys
}

type SubmissionBatch struct {
	SubmissionName        string `json:"submission_name"`
	SubmissionType        string `json:"submission_type"`
	Note                  string `json:"note"`
	SaveExtractedMetadata bool   `json:"save_extracted_metadata"`
	AssignmentId          string `json:"assignment_id"`
	BriefId               string `json:"brief_id"`
	EventId               string `json:"event_id"`
}

func (b *SubmissionBatch) TypeIsValid() bool {
	return batchTypeIsValid[b.SubmissionType]
}

func (b SubmissionBatch) NameIsValid() bool {
	return len(b.SubmissionName) > 0
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
