package espsdk

// These constants represent the root path of the ESP API and the
// relative paths for various endpoints.
const (
	ESPAPIRoot    = "https://esp-sandbox.api.gettyimages.com/esp"
	OAuthEndpoint = "https://api.gettyimages.com/oauth2/token"

	APIInvariant       = "/submission/v1"
	Batches            = APIInvariant + "/submission_batches"
	Compositions       = APIInvariant + "/people_metadata/compositions"
	ControlledValues   = APIInvariant + "/controlled_values/index"
	Expressions        = APIInvariant + "/people_metadata/expressions"
	Keywords           = APIInvariant + "/keywords/getty"
	NumberOfPeople     = APIInvariant + "/people_metadata/number_of_people"
	Personalities      = APIInvariant + "/personalities"
	TranscoderMappings = APIInvariant + "/video_transcoder_mapping_values"
)

// BatchPath returns the canonical path for a(ll) Submission Batch(es).
func BatchPath(b *Batch) string {
	if b.ID == "" {
		return Batches
	}
	return Batches + "/" + b.ID
}
