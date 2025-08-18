package interfaces

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testMemoryStorage(t *testing.T, scoreDec func()) {
	storage := NewMemoryStorage()
	if storage == nil {
		t.Error("NewMemoryStorage() returned nil")
		scoreDec()
		return
	}

	// Test store and retrieve
	key := "test-key"
	value := []byte("test-value")

	err := storage.Store(key, value)
	if err != nil {
		t.Errorf("Store() error = %v", err)
		scoreDec()
	}

	retrieved, err := storage.Retrieve(key)
	if err != nil {
		t.Errorf("Retrieve() error = %v", err)
		scoreDec()
	}

	if !cmp.Equal(retrieved, value) {
		t.Errorf("Retrieve() = %v, want %v\ndiff: %s", retrieved, value, cmp.Diff(value, retrieved))
		scoreDec()
	}

	// Test delete
	err = storage.Delete(key)
	if err != nil {
		t.Errorf("Delete() error = %v", err)
		scoreDec()
	}

	// Should not find deleted key
	_, err = storage.Retrieve(key)
	if err == nil {
		t.Error("Retrieve() should return error for deleted key")
		scoreDec()
	}

	// Test empty key
	err = storage.Store("", value)
	if err == nil {
		t.Error("Store() should return error for empty key")
		scoreDec()
	}
}

func testMemoryCache(t *testing.T, scoreDec func()) {
	cache := NewMemoryCache()
	if cache == nil {
		t.Error("NewMemoryCache() returned nil")
		scoreDec()
		return
	}

	// Test set and get
	key := "cache-key"
	value := []byte("cache-value")

	err := cache.Set(key, value)
	if err != nil {
		t.Errorf("Set() error = %v", err)
		scoreDec()
	}

	retrieved, err := cache.Get(key)
	if err != nil {
		t.Errorf("Get() error = %v", err)
		scoreDec()
	}

	if !cmp.Equal(retrieved, value) {
		t.Errorf("Get() = %v, want %v\ndiff: %s", retrieved, value, cmp.Diff(value, retrieved))
		scoreDec()
	}

	// Test clear
	err = cache.Clear(key)
	if err != nil {
		t.Errorf("Clear() error = %v", err)
		scoreDec()
	}

	// Should not find cleared key
	_, err = cache.Get(key)
	if err == nil {
		t.Error("Get() should return error for cleared key")
		scoreDec()
	}
}

func testSimpleLogger(t *testing.T, scoreDec func()) {
	logger := NewSimpleLogger()
	if logger == nil {
		t.Error("NewSimpleLogger() returned nil")
		scoreDec()
		return
	}

	// Test initial state
	logs := logger.GetLogs()
	if len(logs) != 0 {
		t.Errorf("New logger should have 0 logs, got %d", len(logs))
		scoreDec()
	}

	// Test logging
	message1 := "First log message"
	message2 := "Second log message"

	err := logger.Log(message1)
	if err != nil {
		t.Errorf("Log() error = %v", err)
		scoreDec()
	}

	err = logger.Log(message2)
	if err != nil {
		t.Errorf("Log() error = %v", err)
		scoreDec()
	}

	logs = logger.GetLogs()
	if len(logs) != 2 {
		t.Errorf("Expected 2 logs, got %d", len(logs))
		scoreDec()
	}

	if logs[0] != message1 || logs[1] != message2 {
		t.Errorf("Logs = %v, want [%s, %s]", logs, message1, message2)
		scoreDec()
	}
}

func testCompositeCloudService(t *testing.T, scoreDec func()) {
	storage := NewMemoryStorage()
	cache := NewMemoryCache()
	logger := NewSimpleLogger()

	service := NewCompositeCloudService(storage, cache, logger)
	if service == nil {
		t.Error("NewCompositeCloudService() returned nil")
		scoreDec()
		return
	}

	// Test that service implements CloudService interface
	var _ CloudService = service

	// Test storage functionality through service
	key := "service-key"
	value := []byte("service-value")

	err := service.Store(key, value)
	if err != nil {
		t.Errorf("Service.Store() error = %v", err)
		scoreDec()
	}

	retrieved, err := service.Retrieve(key)
	if err != nil {
		t.Errorf("Service.Retrieve() error = %v", err)
		scoreDec()
	}

	if !cmp.Equal(retrieved, value) {
		t.Errorf("Service.Retrieve() = %v, want %v\ndiff: %s", retrieved, value, cmp.Diff(value, retrieved))
		scoreDec()
	}

	// Test cache functionality through service
	err = service.Set("cache-key", []byte("cache-val"))
	if err != nil {
		t.Errorf("Service.Set() error = %v", err)
		scoreDec()
	}

	// Test logging functionality through service
	err = service.Log("test message")
	if err != nil {
		t.Errorf("Service.Log() error = %v", err)
		scoreDec()
	}
}

func testProcessRequest(t *testing.T, scoreDec func()) {
	storage := NewMemoryStorage()
	cache := NewMemoryCache()
	logger := NewSimpleLogger()

	service := NewCompositeCloudService(storage, cache, logger)

	// First, store some data in storage
	key := "request-key"
	value := []byte("request-value")
	storage.Store(key, value)

	// Test first request (cache miss, storage hit)
	result, err := service.ProcessRequest(key)
	if err != nil {
		t.Errorf("ProcessRequest() error = %v", err)
		scoreDec()
	}

	if !cmp.Equal(result, value) {
		t.Errorf("ProcessRequest() = %v, want %v\ndiff: %s", result, value, cmp.Diff(value, result))
		scoreDec()
	}

	// Check that logs were created
	logs := logger.GetLogs()
	if len(logs) < 3 {
		t.Errorf("Expected at least 3 log messages, got %d", len(logs))
		scoreDec()
	}

	// Verify log content contains expected keywords
	logContent := strings.Join(logs, " ")
	if !strings.Contains(logContent, "Processing request") {
		t.Error("Logs should contain 'Processing request'")
		scoreDec()
	}

	// Test second request (should be cache hit)
	logger.logs = nil // Clear logs for second test
	result2, err := service.ProcessRequest(key)
	if err != nil {
		t.Errorf("Second ProcessRequest() error = %v", err)
		scoreDec()
	}

	if !cmp.Equal(result2, value) {
		t.Errorf("Second ProcessRequest() = %v, want %v\ndiff: %s", result2, value, cmp.Diff(value, result2))
		scoreDec()
	}

	// Check cache hit logs
	logs2 := logger.GetLogs()
	logContent2 := strings.Join(logs2, " ")
	if !strings.Contains(logContent2, "Cache hit") {
		t.Error("Second request should result in cache hit")
		scoreDec()
	}

	// Test request for non-existent key
	_, err = service.ProcessRequest("non-existent")
	if err == nil {
		t.Error("ProcessRequest() should return error for non-existent key")
		scoreDec()
	}
}
