package espsdk

import "encoding/json"

type TranscoderMapping struct {
	FrameSize             string `json:"frame_size,omitempty"`
	FrameRate             string `json:"frame_rate,omitempty"`
	FrameComposition      string `json:"frame_composition,omitempty"`
	MasteredToCompression string `json:"mastered_to_compression,omitempty"`
}

type TranscoderMappingList struct {
	GettyVideoMappings  []TranscoderMapping `json:"getty_video_mappings",omitempty`
	IstockVideoMappings []TranscoderMapping `json:"istock_video_mappings",omitempty`
}

// Unmarshal attempts to deserialize the provided JSON payload
// into an object.
func (tml TranscoderMappingList) Unmarshal(payload []byte) TranscoderMappingList {
	var dest TranscoderMappingList
	if err := json.Unmarshal(payload, &dest); err != nil {
		panic(err)
	}
	return dest
}
