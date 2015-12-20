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

func BatchPath(b *Batch) string { return fmt.Sprintf("%s/%d", Batches, b.ID) }

func ReleasePath(batchID int, releaseID int) string {
	return fmt.Sprintf("%s/%d/releases/%d", Batches, batchID, releaseID)
}

func ContributionPath(batchID int, contributionID int) string {
	return fmt.Sprintf("%s/%d/contributions/%d", Batches, batchID, contributionID)
}
