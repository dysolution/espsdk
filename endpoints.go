package espsdk

// These constants represent the root path of the ESP API and the
// relative paths for various endpoints.
const (
	OAuthEndpoint = "https://api.gettyimages.com/oauth2/token"

	ProdAPI    = "https://api.gettyimages.com/esp"
	SandboxAPI = "https://esp-sandbox.api.gettyimages.com/esp"

	APIInvariant               = "/submission/v1"
	BatchesEndpoint            = APIInvariant + "/submission_batches"
	CompositionsEndpoint       = APIInvariant + "/people_metadata/compositions"
	ControlledValuesEndpoint   = APIInvariant + "/controlled_values/index"
	EventsEndpoint             = APIInvariant + "/events"
	ExpressionsEndpoint        = APIInvariant + "/people_metadata/expressions"
	KeywordsEndpoint           = APIInvariant + "/keywords/getty"
	NumberOfPeopleEndpoint     = APIInvariant + "/people_metadata/number_of_people"
	PersonalitiesEndpoint      = APIInvariant + "/personalities"
	TranscoderMappingsEndpoint = APIInvariant + "/video_transcoder_mapping_values"
)

// BatchPath returns the canonical path for a(ll) Submission Batch(es).
func BatchPath(b *Batch) string {
	if b.ID == "" {
		return BatchesEndpoint
	}
	return BatchesEndpoint + "/" + b.ID
}
