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

// Composite interface
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
	logs []string
}

func NewSimpleLogger() *SimpleLogger {
	return &SimpleLogger{logs: []string{}}
}

func (s *SimpleLogger) Log(message string) error {
	s.logs = append(s.logs, message)
	return nil
}

func (s *SimpleLogger) GetLogs() []string {
	return s.logs
}

// --- CompositeCloudService ---

type CompositeCloudService struct {
	storage Storage
	cache   Cache
	logger  *SimpleLogger
}

func NewCompositeCloudService(storage Storage, cache Cache, logger Logger) *CompositeCloudService {
	return &CompositeCloudService{
		storage: storage,
		cache:   cache,
		logger:  logger.(*SimpleLogger),
	}
}

// Storage
func (c *CompositeCloudService) Store(key string, value []byte) error {
	return c.storage.Store(key, value)
}

func (c *CompositeCloudService) Retrieve(key string) ([]byte, error) {
	return c.storage.Retrieve(key)
}

func (c *CompositeCloudService) Delete(key string) error {
	return c.storage.Delete(key)
}

// Cache
func (c *CompositeCloudService) Set(key string, value []byte) error {
	return c.cache.Set(key, value)
}

func (c *CompositeCloudService) Get(key string) ([]byte, error) {
	return c.cache.Get(key)
}

func (c *CompositeCloudService) Clear(key string) error {
	return c.cache.Clear(key)
}

// Logger
func (c *CompositeCloudService) Log(message string) error {
	return c.logger.Log(message)
}

// ProcessRequest
func (c *CompositeCloudService) ProcessRequest(key string) ([]byte, error) {
	c.Log("Processing request: " + key)

	val, err := c.cache.Get(key)
	if err == nil {
		c.Log("Cache hit: " + key)
		return val, nil
	}

	c.Log("Cache miss: " + key) // <--- este log faltaba

	val, err = c.storage.Retrieve(key)
	if err != nil {
		return nil, err
	}

	c.cache.Set(key, val)
	c.Log("Request completed: " + key)

	return val, nil
}
