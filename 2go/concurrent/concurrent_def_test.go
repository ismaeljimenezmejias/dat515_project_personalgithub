package concurrent

import (
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func testProcessConcurrently(t *testing.T, scoreDec func()) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}

	result := ProcessConcurrently(input)
	if !cmp.Equal(result, expected) {
		t.Errorf("ProcessConcurrently(%v) = %v, want %v\ndiff: %s", input, result, expected, cmp.Diff(expected, result))
		scoreDec()
	}

	// Test empty slice
	emptyResult := ProcessConcurrently([]int{})
	if len(emptyResult) != 0 {
		t.Errorf("ProcessConcurrently([]) should return empty slice, got %v", emptyResult)
		scoreDec()
	}
}

func testWorkerPool(t *testing.T, scoreDec func()) {
	tasks := make(chan int, 5)
	tasks <- 1
	tasks <- 2
	tasks <- 3
	tasks <- 4
	tasks <- 5
	close(tasks)

	results := WorkerPool(3, tasks)
	if results == nil {
		t.Error("WorkerPool() returned nil channel")
		scoreDec()
		return
	}

	var received []int
	for result := range results {
		if result.Error != nil {
			t.Errorf("Worker returned error: %v", result.Error)
			scoreDec()
		}
		received = append(received, result.Result)
	}

	sort.Ints(received)
	expected := []int{2, 4, 6, 8, 10}
	if !cmp.Equal(received, expected) {
		t.Errorf("WorkerPool results = %v, want %v\ndiff: %s", received, expected, cmp.Diff(expected, received))
		scoreDec()
	}
}

func testRateLimitedProcessor(t *testing.T, scoreDec func()) {
	input := []int{1, 2, 3, 4}
	rateLimit := 2

	start := time.Now()
	result := RateLimitedProcessor(input, rateLimit)
	duration := time.Since(start)

	// Sort results since concurrent processing may change order
	sort.Ints(result)
	expected := []int{2, 4, 6, 8}
	sort.Ints(expected)

	if !cmp.Equal(result, expected) {
		t.Errorf("RateLimitedProcessor(%v, %d) = %v, want %v\ndiff: %s", input, rateLimit, result, expected, cmp.Diff(expected, result))
		scoreDec()
	}

	// The rate limiting should not make it too slow for this small test
	if duration > 1*time.Second {
		t.Errorf("RateLimitedProcessor took too long: %v", duration)
		scoreDec()
	}
}

func testFanOutFanIn(t *testing.T, scoreDec func()) {
	input := []int{1, 2, 3, 4, 5}
	numWorkers := 3

	result := FanOutFanIn(input, numWorkers)

	// Sort results since fan-out/fan-in may change order
	sort.Ints(result)
	expected := []int{2, 4, 6, 8, 10}
	sort.Ints(expected)

	if !cmp.Equal(result, expected) {
		t.Errorf("FanOutFanIn(%v, %d) = %v, want %v\ndiff: %s", input, numWorkers, result, expected, cmp.Diff(expected, result))
		scoreDec()
	}

	// Test edge cases
	emptyResult := FanOutFanIn([]int{}, 3)
	if len(emptyResult) != 0 {
		t.Errorf("FanOutFanIn([], 3) should return empty slice, got %v", emptyResult)
		scoreDec()
	}

	zeroWorkerResult := FanOutFanIn(input, 0)
	if len(zeroWorkerResult) != 0 {
		t.Errorf("FanOutFanIn with 0 workers should return empty slice, got %v", zeroWorkerResult)
		scoreDec()
	}
}

func testSafeCounter(t *testing.T, scoreDec func()) {
	counter := NewSafeCounter()
	if counter == nil {
		t.Error("NewSafeCounter() returned nil")
		scoreDec()
		return
	}

	// Test initial value
	if counter.Value() != 0 {
		t.Errorf("New counter value = %d, want 0", counter.Value())
		scoreDec()
	}

	// Test sequential increment
	counter.Increment()
	if counter.Value() != 1 {
		t.Errorf("After one increment, value = %d, want 1", counter.Value())
		scoreDec()
	}

	// Test concurrent increments
	numGoroutines := 100
	incrementsPerGoroutine := 10
	var wg sync.WaitGroup

	for range numGoroutines {
		wg.Go(func() {
			for range incrementsPerGoroutine {
				counter.Increment()
			}
		})
	}

	wg.Wait()

	expected := 1 + (numGoroutines * incrementsPerGoroutine)
	if counter.Value() != expected {
		t.Errorf("After concurrent increments, value = %d, want %d", counter.Value(), expected)
		scoreDec()
	}
}
