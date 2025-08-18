// Package httpclient provides exercises for HTTP client operations
// essential for cloud service communication and API integration.
package httpclient

import (
	"net/http"
)

// Task: HTTP Client for API Communication
//
// This exercise teaches HTTP client operations which are essential for
// communicating with REST APIs and cloud services. You'll implement functions
// to make GET and POST requests to APIs.

// APIResponse represents a generic API response
type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// GetRequest makes a GET request to the specified URL and returns the response.
// It should return the HTTP response object and any error that occurs.
// Don't forget to close the response body when you're done with it.
func GetRequest(url string) (*http.Response, error) {
	// TODO: Implement GET request
	return nil, nil
}

// PostJSONRequest makes a POST request with JSON data to the specified URL.
// The jsonData parameter contains the JSON payload to send.
// It should set the Content-Type header to "application/json".
// Don't forget to close the response body when you're done with it.
func PostJSONRequest(url string, jsonData []byte) (*http.Response, error) {
	// TODO: Implement POST request with JSON
	return nil, nil
}

// CheckHealthEndpoint makes a GET request to a health check endpoint
// and returns true if the status code is 200, false otherwise.
// This is a common pattern for checking if a service is healthy.
func CheckHealthEndpoint(url string) (bool, error) {
	// TODO: Implement health check
	return false, nil
}

// GetWithHeaders makes a GET request with custom headers.
// The headers parameter is a map of header names to header values.
func GetWithHeaders(url string, headers map[string]string) (*http.Response, error) {
	// TODO: Implement GET request with custom headers
	return nil, nil
}
