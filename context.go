package log

import (
	"context"
)

type contextKey struct{}

// loggerKey is used as key to store logger in context
var loggerKey contextKey

// ContextWithLogger set *Logger into given context and return
func ContextWithLogger(ctx context.Context, l *Logger) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, loggerKey, l)
}

// FromContext get *Logger from context
// Notice: If ctx is nil or no Logger was set before, it will return nil
func FromContext(ctx context.Context) *Logger {
	if ctx == nil {
		return nil
	}
	l, ok := ctx.Value(loggerKey).(*Logger)
	if !ok {
		return nil
	}
	return l
}
