package contextops

import (
	"context"
	"errors"
	"testing"
	"time"
)

func testProcessWithTimeout(t *testing.T, scoreDec func()) {
	// Test operation that completes before timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := ProcessWithTimeout(ctx, 50*time.Millisecond)
	if err != nil {
		t.Errorf("ProcessWithTimeout() should complete successfully, got error: %v", err)
		scoreDec()
	}

	// Test operation that times out
	ctx2, cancel2 := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel2()

	err = ProcessWithTimeout(ctx2, 100*time.Millisecond)
	if err == nil {
		t.Error("ProcessWithTimeout() should return error when context times out")
		scoreDec()
	}

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("ProcessWithTimeout() should return DeadlineExceeded, got: %v", err)
		scoreDec()
	}
}

func testProcessWithCancellation(t *testing.T, scoreDec func()) {
	// Test operation that completes normally
	ctx := context.Background()
	err := ProcessWithCancellation(ctx, 5)
	if err != nil {
		t.Errorf("ProcessWithCancellation() should complete successfully, got error: %v", err)
		scoreDec()
	}

	// Test operation that gets cancelled
	ctx2, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond) // Let it start working
		cancel()
	}()

	err = ProcessWithCancellation(ctx2, 100) // Many steps, should be cancelled
	if err == nil {
		t.Error("ProcessWithCancellation() should return error when context is cancelled")
		scoreDec()
	}

	if !errors.Is(err, context.Canceled) {
		t.Errorf("ProcessWithCancellation() should return Canceled, got: %v", err)
		scoreDec()
	}
}

func testRequestIDFromContext(t *testing.T, scoreDec func()) {
	// Test with request ID present
	ctx := context.WithValue(context.Background(), requestIDKey, "test-123")
	reqID := RequestIDFromContext(ctx)
	if reqID != "test-123" {
		t.Errorf("RequestIDFromContext() = %s, want test-123", reqID)
		scoreDec()
	}

	// Test with no request ID
	emptyCtx := context.Background()
	reqID = RequestIDFromContext(emptyCtx)
	if reqID != "" {
		t.Errorf("RequestIDFromContext() = %s, want empty string", reqID)
		scoreDec()
	}
}

func testContextWithRequestID(t *testing.T, scoreDec func()) {
	ctx := context.Background()
	requestID := "test-456"

	newCtx := ContextWithRequestID(ctx, requestID)
	if newCtx == nil {
		t.Error("ContextWithRequestID() returned nil context")
		scoreDec()
		return
	}

	// Verify the request ID was stored
	retrievedID := RequestIDFromContext(newCtx)
	if retrievedID != requestID {
		t.Errorf("ContextWithRequestID() stored %s, want %s", retrievedID, requestID)
		scoreDec()
	}
}

func testTimeoutHandler(t *testing.T, scoreDec func()) {
	// Test handler that completes before timeout
	err := TimeoutHandler(100*time.Millisecond, 50*time.Millisecond)
	if err != nil {
		t.Errorf("TimeoutHandler() should complete successfully, got error: %v", err)
		scoreDec()
	}

	// Test handler that times out
	err = TimeoutHandler(50*time.Millisecond, 100*time.Millisecond)
	if err == nil {
		t.Error("TimeoutHandler() should return error when operation times out")
		scoreDec()
	}

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("TimeoutHandler() should return DeadlineExceeded, got: %v", err)
		scoreDec()
	}
}
