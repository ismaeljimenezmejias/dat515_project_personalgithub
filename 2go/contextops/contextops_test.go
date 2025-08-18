package contextops

import "testing"

func TestProcessWithTimeout(t *testing.T) {
	testProcessWithTimeout(t, func() {})
}

func TestProcessWithCancellation(t *testing.T) {
	testProcessWithCancellation(t, func() {})
}

func TestRequestIDFromContext(t *testing.T) {
	testRequestIDFromContext(t, func() {})
}

func TestContextWithRequestID(t *testing.T) {
	testContextWithRequestID(t, func() {})
}

func TestTimeoutHandler(t *testing.T) {
	testTimeoutHandler(t, func() {})
}
