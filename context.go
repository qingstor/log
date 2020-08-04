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
	// if nil logger was given, return ctx directly
	if l == nil {
		return ctx
	}

	return context.WithValue(ctx, loggerKey, l)
}

// FromContext get *Logger from context
// Notice: If ctx is nil or no Logger was set before, it will return a default logger
func FromContext(ctx context.Context) *Logger {
	if ctx == nil {
		return New()
	}
	l, ok := ctx.Value(loggerKey).(*Logger)
	if !ok {
		return New()
	}
	return l
}
