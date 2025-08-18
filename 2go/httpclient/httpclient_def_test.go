package httpclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testGetRequest(t *testing.T, scoreDec func()) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
			scoreDec()
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}))
	defer server.Close()

	resp, err := GetRequest(server.URL)
	if err != nil {
		t.Errorf("GetRequest() error = %v", err)
		scoreDec()
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("GetRequest() status = %d, want %d", resp.StatusCode, http.StatusOK)
		scoreDec()
	}
}

func testPostJSONRequest(t *testing.T, scoreDec func()) {
	testData := []byte(`{"message": "Hello, World!"}`)

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
			scoreDec()
			return
		}

		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", contentType)
			scoreDec()
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	resp, err := PostJSONRequest(server.URL, testData)
	if err != nil {
		t.Errorf("PostJSONRequest() error = %v", err)
		scoreDec()
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("PostJSONRequest() status = %d, want %d", resp.StatusCode, http.StatusOK)
		scoreDec()
	}
}

func testCheckHealthEndpoint(t *testing.T, scoreDec func()) {
	// Test healthy endpoint
	healthyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer healthyServer.Close()

	healthy, err := CheckHealthEndpoint(healthyServer.URL)
	if err != nil {
		t.Errorf("CheckHealthEndpoint() error = %v", err)
		scoreDec()
		return
	}

	if !healthy {
		t.Error("CheckHealthEndpoint() = false, want true for healthy endpoint")
		scoreDec()
	}

	// Test unhealthy endpoint
	unhealthyServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error"))
	}))
	defer unhealthyServer.Close()

	healthy, err = CheckHealthEndpoint(unhealthyServer.URL)
	if err != nil {
		t.Errorf("CheckHealthEndpoint() error = %v", err)
		scoreDec()
		return
	}

	if healthy {
		t.Error("CheckHealthEndpoint() = true, want false for unhealthy endpoint")
		scoreDec()
	}
}

func testGetWithHeaders(t *testing.T, scoreDec func()) {
	expectedHeaders := map[string]string{
		"Authorization": "Bearer token123",
		"User-Agent":    "TestClient/1.0",
	}

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
			scoreDec()
			return
		}

		for key, expectedValue := range expectedHeaders {
			actualValue := r.Header.Get(key)
			if actualValue != expectedValue {
				t.Errorf("Expected header %s: %s, got %s", key, expectedValue, actualValue)
				scoreDec()
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	resp, err := GetWithHeaders(server.URL, expectedHeaders)
	if err != nil {
		t.Errorf("GetWithHeaders() error = %v", err)
		scoreDec()
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("GetWithHeaders() status = %d, want %d", resp.StatusCode, http.StatusOK)
		scoreDec()
	}

	// Test with invalid URL to ensure error handling
	_, err = GetWithHeaders("://invalid-url", expectedHeaders)
	if err == nil {
		t.Error("GetWithHeaders() should return error for invalid URL")
		scoreDec()
	}
}
