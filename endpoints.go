package espsdk

// These constants represent the root path of the ESP API and the
// relative paths for various endpoints.
const (
	OAuthEndpoint = "https://api.gettyimages.com/oauth2/token"

	ProdAPI    = "https://api.gettyimages.com/esp"
	SandboxAPI = "https://esp-sandbox.api.gettyimages.com/esp"
)

// Endpoints are the relative paths for the ESP API.
var Endpoints = struct {
	Batches            string
	Compositions       string
	ControlledValues   string
	Events             string
	Expressions        string
	FieldRestrictions  string
	Keywords           string
	NumberOfPeople     string
	Personalities      string
	TranscoderMappings string
}{
	"/submission/v1/submission_batches",
	"/submission/v1/people_metadata/compositions",
	"/submission/v1/controlled_values/index",
	"/submission/v1/events",
	"/submission/v1/people_metadata/expressions",
	"/account/v1/field_restrictions",
	"/submission/v1/keywords/getty",
	"/submission/v1/people_metadata/number_of_people",
	"/submission/v1/personalities",
	"/submission/v1/video_transcoder_mapping_values",
}

// BatchPath returns the canonical path for a(ll) Submission Batch(es).
func BatchPath(b *Batch) string {
	p := Endpoints.Batches
	if b.ID == "" {
		return p
	}
	return p + "/" + b.ID
}
