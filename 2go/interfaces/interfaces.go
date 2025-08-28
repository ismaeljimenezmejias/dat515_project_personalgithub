package interfaces

import "errors"

// --- Interfaces ---

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

// --- MemoryStorage ---

type MemoryStorage struct {
	data map[string][]byte
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string][]byte)}
}

func (m *MemoryStorage) Store(key string, value []byte) error {
	if key == "" {
		return errors.New("empty key")
	}
	m.data[key] = value
	return nil
}

func (m *MemoryStorage) Retrieve(key string) ([]byte, error) {
	val, ok := m.data[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return val, nil
}

func (m *MemoryStorage) Delete(key string) error {
	delete(m.data, key)
	return nil
}

// --- MemoryCache ---

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
	val, ok := c.data[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return val, nil
}

func (c *MemoryCache) Clear(key string) error {
	delete(c.data, key)
	return nil
}

// --- SimpleLogger ---

type SimpleLogger struct {
	Logs []string
}

func NewSimpleLogger() *SimpleLogger {
	return &SimpleLogger{Logs: []string{}}
}

func (s *SimpleLogger) Log(message string) error {
	s.Logs = append(s.Logs, message)
	return nil
}

func (s *SimpleLogger) GetLogs() []string {
	return s.Logs
}

// --- CompositeCloudService ---

type CompositeCloudService struct {
	Storage Storage
	Cache   Cache
	Logger  Logger
}

func NewCompositeCloudService(storage Storage, cache Cache, logger Logger) *CompositeCloudService {
	return &CompositeCloudService{
		Storage: storage,
		Cache:   cache,
		Logger:  logger,
	}
}

// Storage methods
func (c *CompositeCloudService) Store(key string, value []byte) error {
	return c.Storage.Store(key, value)
}

func (c *CompositeCloudService) Retrieve(key string) ([]byte, error) {
	return c.Storage.Retrieve(key)
}

func (c *CompositeCloudService) Delete(key string) error {
	return c.Storage.Delete(key)
}

// Cache methods
func (c *CompositeCloudService) Set(key string, value []byte) error {
	return c.Cache.Set(key, value)
}

func (c *CompositeCloudService) Get(key string) ([]byte, error) {
	return c.Cache.Get(key)
}

func (c *CompositeCloudService) Clear(key string) error {
	return c.Cache.Clear(key)
}

// Logger method
func (c *CompositeCloudService) Log(message string) error {
	return c.Logger.Log(message)
}

// ProcessRequest: uses logger, cache, and storage
func (c *CompositeCloudService) ProcessRequest(key string) ([]byte, error) {
	c.Log("Processing request: " + key)
	c.Log("Checking cache for key: " + key)

	val, err := c.Cache.Get(key)
	if err == nil {
		c.Log("Cache hit: " + key)
		return val, nil
	}

	val, err = c.Storage.Retrieve(key)
	if err != nil {
		return nil, err
	}

	c.Cache.Set(key, val)
	c.Log("Request completed: " + key)
	return val, nil
}
