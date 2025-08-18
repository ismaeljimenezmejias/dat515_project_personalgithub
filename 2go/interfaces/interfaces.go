// Package interfaces provides exercises for interface design and composition
// which are essential for building clean, testable, and maintainable cloud applications.
package interfaces

// Task: Interface Composition for Cloud Architecture
//
// This exercise teaches interface design and composition which is essential
// for building clean, testable, and maintainable cloud applications.
// You'll learn patterns commonly used in cloud services for abstraction
// and dependency injection.

// Storage represents a basic storage interface
type Storage interface {
	Store(key string, value []byte) error
	Retrieve(key string) ([]byte, error)
	Delete(key string) error
}

// Cache represents a caching interface
type Cache interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Clear(key string) error
}

// Logger represents a logging interface
type Logger interface {
	Log(message string) error
}

// CloudService combines storage, caching, and logging capabilities
// This demonstrates interface composition - a key pattern in Go
type CloudService interface {
	Storage
	Cache
	Logger
}

// MemoryStorage implements the Storage interface using in-memory storage
type MemoryStorage struct {
	// TODO: Add necessary fields
}

// NewMemoryStorage creates a new MemoryStorage instance
func NewMemoryStorage() *MemoryStorage {
	// TODO: Initialize MemoryStorage
	return nil
}

// Store implements the Storage interface
func (m *MemoryStorage) Store(key string, value []byte) error {
	// TODO: Implement storage functionality
	return nil
}

// Retrieve implements the Storage interface
func (m *MemoryStorage) Retrieve(key string) ([]byte, error) {
	// TODO: Implement retrieval functionality
	return nil, nil
}

// Delete implements the Storage interface
func (m *MemoryStorage) Delete(key string) error {
	// TODO: Implement deletion functionality
	return nil
}

// MemoryCache implements the Cache interface using in-memory caching
type MemoryCache struct {
	// TODO: Add necessary fields
}

// NewMemoryCache creates a new MemoryCache instance
func NewMemoryCache() *MemoryCache {
	// TODO: Initialize MemoryCache
	return nil
}

// Set implements the Cache interface
func (m *MemoryCache) Set(key string, value []byte) error {
	// TODO: Implement cache set functionality
	return nil
}

// Get implements the Cache interface
func (m *MemoryCache) Get(key string) ([]byte, error) {
	// TODO: Implement cache get functionality
	return nil, nil
}

// Clear implements the Cache interface
func (m *MemoryCache) Clear(key string) error {
	// TODO: Implement cache clear functionality
	return nil
}

// SimpleLogger implements the Logger interface
type SimpleLogger struct {
	// TODO: Add necessary fields for storing log messages
}

// NewSimpleLogger creates a new SimpleLogger instance
func NewSimpleLogger() *SimpleLogger {
	// TODO: Initialize SimpleLogger
	return nil
}

// Log implements the Logger interface
func (s *SimpleLogger) Log(message string) error {
	// TODO: Implement logging functionality (store messages in memory)
	return nil
}

// GetLogs returns all logged messages (for testing purposes)
func (s *SimpleLogger) GetLogs() []string {
	// TODO: Return stored log messages
	return nil
}

// CompositeCloudService implements CloudService by composing other services
// This demonstrates how to build complex services from simple interfaces
type CompositeCloudService struct {
	// TODO: Add embedded interfaces or separate fields
}

// NewCompositeCloudService creates a new CompositeCloudService
func NewCompositeCloudService(storage Storage, cache Cache, logger Logger) *CompositeCloudService {
	// TODO: Initialize CompositeCloudService with provided components
	return nil
}

// ProcessRequest simulates processing a request using all three capabilities
// It should: 1) Log the request, 2) Check cache, 3) If not in cache, get from storage
// and put in cache, 4) Log completion, 5) Return the data
func (c *CompositeCloudService) ProcessRequest(key string) ([]byte, error) {
	// TODO: Implement request processing logic
	return nil, nil
}
