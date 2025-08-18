package httpclient

import "testing"

func TestGetRequest(t *testing.T) {
	testGetRequest(t, func() {})
}

func TestPostJSONRequest(t *testing.T) {
	testPostJSONRequest(t, func() {})
}

func TestCheckHealthEndpoint(t *testing.T) {
	testCheckHealthEndpoint(t, func() {})
}

func TestGetWithHeaders(t *testing.T) {
	testGetWithHeaders(t, func() {})
}
