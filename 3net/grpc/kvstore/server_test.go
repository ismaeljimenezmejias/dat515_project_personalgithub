package kvstore

import (
	"testing"
)

func TestRequestSequence(t *testing.T) {
	testRequestSequence(t, func() {})
}
