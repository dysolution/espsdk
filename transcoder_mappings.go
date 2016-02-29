package espsdk

import (
	"encoding/json"
	"fmt"
)

// A TranscoderMapping is a set of parameters that represent a video
// encoding that can be accepted by ESP.
type TranscoderMapping struct {
	FrameSize             string `json:"frame_size,omitempty"`
	FrameRate             string `json:"frame_rate,omitempty"`
	FrameComposition      string `json:"frame_composition,omitempty"`
	MasteredToCompression string `json:"mastered_to_compression,omitempty"`
}

// The TranscoderMappingList contains valid mappings for both Getty video
// and iStock video.
type TranscoderMappingList struct {
	GettyVideoMappings  []TranscoderMapping `json:"getty_video_mappings,omitempty"`
	IstockVideoMappings []TranscoderMapping `json:"istock_video_mappings,omitempty"`
}

// Unmarshal attempts to deserialize the provided JSON payload
// into an object.
func (tml TranscoderMappingList) Unmarshal(payload []byte) *TranscoderMappingList {
	fmt.Printf("%s\n", payload)
	dest := new(TranscoderMappingList)
	if err := json.Unmarshal(payload, dest); err != nil {
		Log.Error(err)
	}
	if dest == nil {
		var d interface{}
		json.Unmarshal(payload, &d)
		Log.Error(d)
	}
	return dest
}
