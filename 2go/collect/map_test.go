package collect

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKeys(t *testing.T) {
	x := map[string]int{
		"a": 1,
		"b": 2,
	}
	got := keys(x)
	want := []string{"a", "b"}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("keys() mismatch (-got +want):\n%s", diff)
	}
}

func TestValues(t *testing.T) {
	x := map[string]int{
		"a": 1,
		"b": 2,
	}
	got := values(x)
	want := []int{1, 2}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("values() mismatch (-got +want):\n%s", diff)
	}
}
