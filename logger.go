package log

import (
	"io"
	"os"

	"github.com/qingstor/log/transform"
)

// New will create a new logger.
func New() *Logger {
	l := &Logger{
		ew:        os.Stderr,
		executor:  ExecuteWrite(os.Stderr),
		transform: transform.NewText,
	}
	return l
}

// Logger is the logger.
type Logger struct {
	ew io.Writer

	executor  Executor
	transform transform.FactoryFunc
	fields    []transform.Transformee
}

// Clone will copy and return a new logger.
func (l *Logger) Clone() *Logger {
	x := &Logger{
		ew:       l.ew,
		fields:   make([]transform.Transformee, len(l.fields)),
		executor: l.executor,
	}

	x.fields = append(x.fields, l.fields...)
	return x
}

// WithExecutor will set logger's executor.
func (l *Logger) WithExecutor(fn Executor) *Logger {
	l.executor = fn
	return l
}

// WithFields will set logger's fields.
func (l *Logger) WithFields(fields ...transform.Transformee) *Logger {
	l.fields = fields
	return l
}

// Print will print an empty entry.
func (l *Logger) Print(fields ...transform.Transformee) {
	e := newEntry(l, EmptyLevel, fields...)
	defer e.Free()

	l.executor(e)
}

// Debug will print an debug entry.
func (l *Logger) Debug(fields ...transform.Transformee) {
	e := newEntry(l, DebugLevel, fields...)
	defer e.Free()

	l.executor(e)
}

// Info will print an info entry.
func (l *Logger) Info(fields ...transform.Transformee) {
	e := newEntry(l, InfoLevel, fields...)
	defer e.Free()

	l.executor(e)
}

// Warn will print an warn entry.
func (l *Logger) Warn(fields ...transform.Transformee) {
	e := newEntry(l, WarnLevel, fields...)
	defer e.Free()

	l.executor(e)
}

// Error will print an error entry.
func (l *Logger) Error(fields ...transform.Transformee) {
	e := newEntry(l, ErrorLevel, fields...)
	defer e.Free()

	l.executor(e)
}
