package concurrent

import "testing"

func TestProcessConcurrently(t *testing.T) {
	testProcessConcurrently(t, func() {})
}

func TestWorkerPool(t *testing.T) {
	testWorkerPool(t, func() {})
}

func TestRateLimitedProcessor(t *testing.T) {
	testRateLimitedProcessor(t, func() {})
}

func TestFanOutFanIn(t *testing.T) {
	testFanOutFanIn(t, func() {})
}

func TestSafeCounter(t *testing.T) {
	testSafeCounter(t, func() {})
}
