package espsdk

import "encoding/json"

// A FieldRestriction determines what business logic will be applied to the
// field when it is being populated or changed.
type FieldRestriction struct {
	DefaultValue       string `json:"default_value,omitempty"`
	RequiredForPublish bool   `json:"required_for_publish,omitempty"`
	Restriction        string `json:"restriction,omitempty"`
}

// A FieldRestrictionQuery is sent by a client to the ESP API with optional
// criteria and is answered with a FieldRestrictionResponse.
type FieldRestrictionQuery struct {
	UserID                string `json:"user_id"`
	FieldRestrictionsType string `json:"field_restrictions_type"`
}

// The FieldRestrictionBody contains the values for all FieldRestrictions.
type FieldRestrictionBody struct {
	AlternateID              FieldRestriction `json:"alternate_id,omitempty"`
	Audio                    FieldRestriction `json:"audio,omitempty"`
	CallForImage             FieldRestriction `json:"call_for_image,omitempty"`
	Caption                  FieldRestriction `json:"caption,omitempty"`
	City                     FieldRestriction `json:"city,omitempty"`
	ClipLength               FieldRestriction `json:"clip_length,omitempty"`
	CollectionCode           FieldRestriction `json:"collection_code,omitempty"`
	ContentProviderName      FieldRestriction `json:"content_provider_name,omitempty"`
	ContentProviderTitle     FieldRestriction `json:"content_provider_title,omitempty"`
	ContentWarnings          FieldRestriction `json:"content_warnings,omitempty"`
	Copyright                FieldRestriction `json:"copyright,omitempty"`
	CountryOfShoot           FieldRestriction `json:"country_of_shoot,omitempty"`
	CreditLine               FieldRestriction `json:"credit_line,omitempty"`
	DSAAlternateIDs          FieldRestriction `json:"dsa_alternate_ids,omitempty"`
	EventID                  FieldRestriction `json:"event_id,omitempty"`
	ExclusiveToGetty         FieldRestriction `json:"exclusive_to_getty,omitempty"`
	ExternalFileLocation     FieldRestriction `json:"external_file_location,omitempty"`
	FrameComposition         FieldRestriction `json:"frame_composition,omitempty"`
	FrameRate                FieldRestriction `json:"frame_rate,omitempty"`
	FrameSize                FieldRestriction `json:"frame_size,omitempty"`
	Headline                 FieldRestriction `json:"headline,omitempty"`
	Keywords                 FieldRestriction `json:"keywords,omitempty"`
	Language                 FieldRestriction `json:"language,omitempty"`
	MasteredToCompression    FieldRestriction `json:"mastered_to_compression,omitempty"`
	MediaFormat              FieldRestriction `json:"media_format,omitempty"`
	MimeType                 FieldRestriction `json:"mime_type,omitempty"`
	OriginalFrameComposition FieldRestriction `json:"original_frame_composition,omitempty"`
	OriginalFrameRate        FieldRestriction `json:"original_frame_rate,omitempty"`
	OriginalFrameSize        FieldRestriction `json:"original_frame_size,omitempty"`
	OriginalMediaFormat      FieldRestriction `json:"original_media_format,omitempty"`
	OriginalProduction       FieldRestriction `json:"original_production_title,omitempty"`
	ParentSource             FieldRestriction `json:"parent_source,omitempty"`
	PixelAspectRatio         FieldRestriction `json:"pixel_aspect_ratio,omitempty"`
	PosterTimecode           FieldRestriction `json:"poster_timecode,omitempty"`
	PreferredLicenseModel    FieldRestriction `json:"preferred_license_model,omitempty"`
	ProvinceState            FieldRestriction `json:"province_state,omitempty"`
	Rank                     FieldRestriction `json:"rank,omitempty"`
	RecordedDate             FieldRestriction `json:"recorded_date,omitempty"`
	Releases                 FieldRestriction `json:"releases,omitempty"`
	RiskCategory             FieldRestriction `json:"risk_category,omitempty"`
	ShotSpeed                FieldRestriction `json:"shot_speed,omitempty"`
	SpecialInstructions      FieldRestriction `json:"special_instructions,omitempty"`
	UploadID                 FieldRestriction `json:"upload_id,omitempty"`
	VisualColor              FieldRestriction `json:"visual_color,omitempty"`
}

// A FieldRestrictionResponse is sent by the ESP API to a client in response to
// a query.
type FieldRestrictionResponse struct {
	Body         FieldRestrictionBody `json:"body"`
	CreatedAt    string               `json:"created_at,omitempty"`
	ID           int                  `json:"id,omitempty"`
	MaxBatchSize int                  `json:"max_batch_size"`
	ModelClass   string               `json:"model_class"`
	UpdatedAt    string               `json:"updated_at,omitempty"`
	UserID       int                  `json:"user_id,omitempty"`
}

// Unmarshal attempts to deserialize the provided JSON payload
// into an object.
func (fr FieldRestriction) Unmarshal(payload []byte) *FieldRestrictionResponse {
	dest := &FieldRestrictionResponse{}
	if err := json.Unmarshal(payload, dest); err != nil {
		Log.WithFields(map[string]interface{}{
			"error":   err,
			"payload": string(payload),
		}).Error()
	}
	return dest
}
