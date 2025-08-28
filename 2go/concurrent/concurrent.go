package concurrent

import (
	"sync"
)

// WorkerResult representa el resultado de un worker
type WorkerResult struct {
	ID     int
	Result int
	Error  error
}

// ProcessConcurrently procesa un slice de enteros en paralelo y duplica cada número
func ProcessConcurrently(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	results := make([]int, len(numbers))
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for i, n := range numbers {
		go func(idx, val int) {
			defer wg.Done()
			results[idx] = val * 2
		}(i, n)
	}

	wg.Wait()
	return results
}

// WorkerPool implementa un pool de workers que duplican los números
func WorkerPool(numWorkers int, tasks <-chan int) <-chan WorkerResult {
	out := make(chan WorkerResult)
	if numWorkers <= 0 {
		close(out)
		return out
	}

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for w := 0; w < numWorkers; w++ {
		go func(workerID int) {
			defer wg.Done()
			for task := range tasks {
				out <- WorkerResult{
					ID:     workerID,
					Result: task * 2,
					Error:  nil,
				}
			}
		}(w)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// RateLimitedProcessor procesa elementos con límite de tasa
func RateLimitedProcessor(items []int, rateLimit int) []int {
	if len(items) == 0 || rateLimit <= 0 {
		return []int{}
	}

	results := make([]int, len(items))
	sem := make(chan struct{}, rateLimit)
	var wg sync.WaitGroup
	wg.Add(len(items))

	for i, val := range items {
		go func(idx, v int) {
			sem <- struct{}{}           // bloquear
			defer func() { <-sem }()    // liberar
			defer wg.Done()
			results[idx] = v * 2
		}(i, val)
	}

	wg.Wait()
	return results
}

// FanOutFanIn distribuye trabajo en numWorkers y recopila resultados
func FanOutFanIn(numbers []int, numWorkers int) []int {
	if len(numbers) == 0 || numWorkers <= 0 {
		return []int{}
	}

	results := make([]int, len(numbers))
	done := make(chan struct{})
	chunkSize := (len(numbers) + numWorkers - 1) / numWorkers

	for w := 0; w < numWorkers; w++ {
		start := w * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		go func(s, e int) {
			for i := s; i < e; i++ {
				results[i] = numbers[i] * 2
			}
			done <- struct{}{}
		}(start, end)
	}

	for w := 0; w < numWorkers; w++ {
		<-done
	}

	return results
}

// SafeCounter es un contador seguro para concurrencia
type SafeCounter struct {
	mu sync.Mutex
	v  int
}

// NewSafeCounter crea un nuevo SafeCounter
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

// Increment incrementa de forma segura
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

// Value devuelve el valor de forma segura
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}
