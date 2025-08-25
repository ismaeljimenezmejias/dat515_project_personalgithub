package httpclient

import (
	"bytes"
	"net/http"
)

// GetRequest makes a GET request to the specified URL and returns the response.
// The caller must close the response body.
func GetRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PostJSONRequest makes a POST request with JSON data to the specified URL.
// The caller must close the response body.
func PostJSONRequest(url string, jsonData []byte) (*http.Response, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CheckHealthEndpoint makes a GET request to a health check endpoint
// and returns true if the status code is 200.
func CheckHealthEndpoint(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}

// GetWithHeaders makes a GET request with custom headers.
func GetWithHeaders(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
