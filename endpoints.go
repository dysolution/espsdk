package espsdk

import "fmt"

// These constants represent the relative paths for various ESP API endpoints.
const (
	Batches            string = "/submission/v1/submission_batches"
	ControlledValues   string = "/submission/v1/controlled_values/index"
	Keywords           string = "/submission/v1/keywords/getty"
	Personalities      string = "/submission/v1/personalities"
	TranscoderMappings string = "/submission/v1/video_transcoder_mapping_values"
	Compositions       string = "/submission/v1/people_metadata/compositions"
	Expressions        string = "/submission/v1/people_metadata/expressions"
	NumberOfPeople     string = "/submission/v1/people_metadata/number_of_people"
)

func BatchPath(b *Batch) string {
	if b.ID == 0 {
		return Batches
	}
	return fmt.Sprintf("%s/%d", Batches, b.ID)
}

func ReleasePath(batchID int, releaseID int) string {
	if releaseID == 0 {
		return fmt.Sprintf("%s/%d/releases", Batches, batchID)
	}
	return fmt.Sprintf("%s/%d/releases/%d", Batches, batchID, releaseID)
}

func ContributionPath(batchID int, contributionID int) string {
	if contributionID == 0 {
		return fmt.Sprintf("%s/%d/contributions", Batches, batchID)
	}
	return fmt.Sprintf("%s/%d/contributions/%d", Batches, batchID, contributionID)
}
