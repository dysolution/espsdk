package espsdk

import (
	"testing"
)

func TestNameIsValid(t *testing.T) {
	b := SubmissionBatch{SubmissionName: ""}
	if b.NameIsValid() {
		t.Errorf("name cannot be blank")
	}
}

func TestTypeIsValidAcceptsValidTypes(t *testing.T) {
	validTypes := []string{
		"getty_creative_video",
		"getty_editorial_video",
		"getty_creative_still",
		"getty_editorial_still",
		"istock_creative_video",
	}
	for _, batchType := range validTypes {
		b := SubmissionBatch{SubmissionType: batchType}
		if b.TypeIsValid() != true {
			t.Errorf("%s should be accepted", batchType)
		}
	}
}

func TestTypeIsValidRejectsInvalidTypes(t *testing.T) {
	invalidTypes := []string{
		"",
		"foo",
		"getty_creative_image",
		"getty_editorial_image",
		"istock_creative_still",
	}
	for _, badType := range invalidTypes {
		b := SubmissionBatch{SubmissionType: badType}
		if b.TypeIsValid() {
			t.Errorf("%s should not be accepted", badType)
		}
	}
}
