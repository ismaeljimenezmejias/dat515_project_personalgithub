package interfaces

import "fmt"

// Interfaces
type Storage interface {
	Store(key string, value []byte) error
	Retrieve(key string) ([]byte, error)
	Delete(key string) error
}

type Cache interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Clear(key string) error
}

type Logger interface {
	Log(message string) error
}

type CloudService interface {
	Storage
	Cache
	Logger
}

// -------- MemoryStorage --------
type MemoryStorage struct {
	data map[string][]byte
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string][]byte)}
}

func (m *MemoryStorage) Store(key string, value []byte) error {
	if key == "" {
		return fmt.Errorf("empty key")
	}
	m.data[key] = value
	return nil
}

func (m *MemoryStorage) Retrieve(key string) ([]byte, error) {
	v, ok := m.data[key]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}
	return v, nil
}

func (m *MemoryStorage) Delete(key string) error {
	delete(m.data, key)
	return nil
}

// -------- MemoryCache --------
type MemoryCache struct {
	data map[string][]byte
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{data: make(map[string][]byte)}
}

func (c *MemoryCache) Set(key string, value []byte) error {
	c.data[key] = value
	return nil
}

func (c *MemoryCache) Get(key string) ([]byte, error) {
	v, ok := c.data[key]
	if !ok {
		return nil, fmt.Errorf("key not found")
	}
	return v, nil
}

func (c *MemoryCache) Clear(key string) error {
	delete(c.data, key)
	return nil
}

// -------- SimpleLogger --------
type SimpleLogger struct {
	Logs []string
}

func NewSimpleLogger() *SimpleLogger {
	return &SimpleLogger{Logs: []string{}}
}

func (l *SimpleLogger) Log(message string) error {
	l.Logs = append(l.Logs, message)
	return nil
}

func (l *SimpleLogger) GetLogs() []string {
	return l.Logs
}

// -------- CompositeCloudService --------
type CompositeCloudService struct {
	Storage
	Cache
	Logger
}

func NewCompositeCloudService(storage Storage, cache Cache, logger Logger) *CompositeCloudService {
	return &CompositeCloudService{
		Storage: storage,
		Cache:   cache,
		Logger:  logger,
	}
}

func (c *CompositeCloudService) ProcessRequest(key string) ([]byte, error) {
	c.Log("Processing request for key: " + key)

	// 1. Check cache first
	val, err := c.Cache.Get(key)
	if err == nil {
		c.Log("Cache hit for key: " + key)
		return val, nil
	}

	// 2. If not in cache, get from storage
	val, err = c.Storage.Retrieve(key)
	if err != nil {
		c.Log("Storage miss for key: " + key)
		return nil, err
	}

	// 3. Put in cache
	c.Cache.Set(key, val)
	c.Log("Stored key in cache: " + key)
	c.Log("Request processed for key: " + key)
	return val, nil
}
