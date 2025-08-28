package interfaces

import "fmt"

// -------- Interfaces --------
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
type MemoryStorage struct{ Data map[string][]byte }

func NewMemoryStorage() *MemoryStorage { return &MemoryStorage{Data: make(map[string][]byte)} }

func (m *MemoryStorage) Store(key string, value []byte) error {
	if key == "" {
		return fmt.Errorf("empty key")
	}
	m.Data[key] = value
	return nil
}

func (m *MemoryStorage) Retrieve(key string) ([]byte, error) {
	val, ok := m.Data[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return val, nil
}

func (m *MemoryStorage) Delete(key string) error {
	delete(m.Data, key)
	return nil
}

// -------- MemoryCache --------
type MemoryCache struct{ Data map[string][]byte }

func NewMemoryCache() *MemoryCache { return &MemoryCache{Data: make(map[string][]byte)} }

func (c *MemoryCache) Set(key string, value []byte) error {
	if key == "" {
		return fmt.Errorf("empty key")
	}
	c.Data[key] = value
	return nil
}

func (c *MemoryCache) Get(key string) ([]byte, error) {
	val, ok := c.Data[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return val, nil
}

func (c *MemoryCache) Clear(key string) error {
	delete(c.Data, key)
	return nil
}

// -------- SimpleLogger --------
type SimpleLogger struct{ Logs []string }

func NewSimpleLogger() *SimpleLogger { return &SimpleLogger{Logs: []string{}} }

func (s *SimpleLogger) Log(message string) error {
	s.Logs = append(s.Logs, message)
	return nil
}

func (s *SimpleLogger) GetLogs() []string { return s.Logs }

// -------- CompositeCloudService --------
type CompositeCloudService struct {
	Storage
	Cache
	Logger
}

func NewCompositeCloudService(storage Storage, cache Cache, logger Logger) *CompositeCloudService {
	return &CompositeCloudService{Storage: storage, Cache: cache, Logger: logger}
}

func (c *CompositeCloudService) ProcessRequest(key string) ([]byte, error) {
	c.Logger.Log("Processing request for key: " + key)

	if val, err := c.Cache.Get(key); err == nil {
		c.Logger.Log("Cache hit for key: " + key)
		return val, nil
	}

	val, err := c.Storage.Retrieve(key)
	if err != nil {
		c.Logger.Log("Storage miss for key: " + key)
		return nil, err
	}

	c.Cache.Set(key, val)
	c.Logger.Log("Stored key in cache: " + key)
	c.Logger.Log("Request processed for key: " + key)
	return val, nil
}
