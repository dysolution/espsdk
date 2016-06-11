package espsdk

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {
	r1 := Release{SubmissionBatchID: "10"}
	want := fmt.Sprintf("%s/%s/releases", Endpoints.Batches, "10")
	got := r1.Path()
	if got != want {
		t.Errorf(`got %v.(%V) want %v.(%V) `, got, want)
	}

	r2 := Release{SubmissionBatchID: "10", ID: "42"}
	want = fmt.Sprintf("%s/%s/releases/%s", Endpoints.Batches, "10", "42")
	got = r2.Path()
	if got != want {
		t.Errorf(`got %v.(%V) want %v.(%V) `, got, want)
	}
}
