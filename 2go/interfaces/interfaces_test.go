package interfaces

import "testing"

func TestMemoryStorage(t *testing.T) {
	testMemoryStorage(t, func() {})
}

func TestMemoryCache(t *testing.T) {
	testMemoryCache(t, func() {})
}

func TestSimpleLogger(t *testing.T) {
	testSimpleLogger(t, func() {})
}

func TestCompositeCloudService(t *testing.T) {
	testCompositeCloudService(t, func() {})
}

func TestProcessRequest(t *testing.T) {
	testProcessRequest(t, func() {})
}
