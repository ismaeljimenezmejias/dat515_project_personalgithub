package contextops

import (
	"context"
	"time"
)

type contextKey string

const requestIDKey contextKey = "request_id"

// ProcessWithTimeout simula una operación que respeta el timeout del contexto.
func ProcessWithTimeout(ctx context.Context, duration time.Duration) error {
	done := make(chan struct{})
	go func() {
		time.Sleep(duration) // Simula trabajo
		close(done)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err() // Retorna el error del contexto (timeout o cancelación)
	case <-done:
		return nil // Terminó sin problemas
	}
}

// ProcessWithCancellation simula una operación que puede ser cancelada.
func ProcessWithCancellation(ctx context.Context, steps int) error {
	for i := 0; i < steps; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err() // Cancelado
		default:
			time.Sleep(100 * time.Millisecond) // Simula trabajo por paso
		}
	}
	return nil
}

// RequestIDFromContext extrae el request ID del contexto.
func RequestIDFromContext(ctx context.Context) string {
	if v := ctx.Value(requestIDKey); v != nil {
		if id, ok := v.(string); ok {
			return id
		}
	}
	return ""
}

// ContextWithRequestID añade un request ID al contexto.
func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// TimeoutHandler simula un handler HTTP con timeout.
func TimeoutHandler(timeout, workDuration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return ProcessWithTimeout(ctx, workDuration)
}
