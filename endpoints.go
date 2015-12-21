package espsdk

import "fmt"

// These constants represent the root path of the ESP API and the
// relative paths for various endpoints.
const (
	ESPAPIRoot         = "https://esp-sandbox.api.gettyimages.com/esp"
	Batches            = "/submission/v1/submission_batches"
	ControlledValues   = "/submission/v1/controlled_values/index"
	Keywords           = "/submission/v1/keywords/getty"
	Personalities      = "/submission/v1/personalities"
	TranscoderMappings = "/submission/v1/video_transcoder_mapping_values"
	Compositions       = "/submission/v1/people_metadata/compositions"
	Expressions        = "/submission/v1/people_metadata/expressions"
	NumberOfPeople     = "/submission/v1/people_metadata/number_of_people"
	oauthEndpoint      = "https://api.gettyimages.com/oauth2/token"
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
