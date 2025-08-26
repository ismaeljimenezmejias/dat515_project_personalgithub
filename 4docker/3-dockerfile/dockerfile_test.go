package dockerfile

import (
	"testing"
)

func TestDockerfileExists(t *testing.T) {
	testDockerfileExists(t, func() {})
}
