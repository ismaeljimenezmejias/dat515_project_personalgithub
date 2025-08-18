// Package contextops provides exercises for Go context usage
// which is essential for managing timeouts, cancellation, and request-scoped values in cloud applications.
package contextops

import (
	"context"
	"time"
)

// Task: Context Usage for Cloud Services
//
// This exercise teaches context usage which is essential for managing timeouts,
// cancellation, and request-scoped values in cloud applications. Context is
// widely used in HTTP servers, database operations, and API calls.

type contextKey string

const requestIDKey contextKey = "request_id"

// ProcessWithTimeout simulates a long-running operation that should respect
// a timeout. The operation should return an error if the context times out
// before completion. The operation takes 'duration' time to complete.
func ProcessWithTimeout(ctx context.Context, duration time.Duration) error {
	// TODO: Implement operation that respects context timeout
	// Hint: Use time.Sleep(duration) to simulate work and select with ctx.Done()
	return nil
}

// ProcessWithCancellation simulates an operation that can be cancelled.
// It should return immediately if the context is cancelled.
// The operation should perform work in steps, checking for cancellation
// between each step.
func ProcessWithCancellation(ctx context.Context, steps int) error {
	// TODO: Implement operation that can be cancelled
	// Hint: Loop through steps, checking ctx.Done() between each step
	return nil
}

// RequestIDFromContext extracts a request ID from the context.
// Cloud services often pass request IDs through context for tracing.
// Return the request ID if found, empty string if not found.
func RequestIDFromContext(ctx context.Context) string {
	// TODO: Extract request ID from context
	// Hint: Use ctx.Value() with requestIDKey
	return ""
}

// ContextWithRequestID creates a new context with a request ID.
// This is commonly used in cloud services for request tracing.
func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	// TODO: Add request ID to context
	// Hint: Use context.WithValue() with requestIDKey
	return ctx
}

// TimeoutHandler simulates an HTTP handler that processes requests with timeout.
// It should create a context with the specified timeout and call ProcessWithTimeout.
// Return an error if the operation times out or fails.
func TimeoutHandler(timeout, workDuration time.Duration) error {
	// TODO: Create context with timeout and call ProcessWithTimeout
	// Hint: Use context.WithTimeout()
	return nil
}
