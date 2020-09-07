package log

import (
	"io"
	"os"

	"github.com/qingstor/log/internal/pkg/bufferpool"
	"github.com/qingstor/log/level"
)

// New will create a new logger with default text config.
func New() *Logger {
	tf, _ := NewText(&TextConfig{
		EntryFormat: defaultFormat,
		TimeFormat:  TimeFormatUnix,
	})
	l := &Logger{
		ew:          os.Stderr,
		executor:    ExecuteWrite(os.Stderr),
		transformer: tf,
	}
	return l
}

// Logger is the logger.
type Logger struct {
	// ew is the error writer of this logger
	// logger will write error message into this while failed to log message
	//
	// NOTE: logger will not check returning error of this writer
	ew io.Writer

	// The executor that logger will execute on every entry.
	executor Executor
	// The transformer that logger used.
	transformer Transformer
	// Logger level fields will be transformed in every entry.
	fields []Transformee
}

// Clone will copy and return a new logger.
func (l *Logger) Clone() *Logger {
	x := &Logger{
		ew: l.ew,

		executor:    l.executor,
		transformer: l.transformer,
		fields:      make([]Transformee, len(l.fields)),
	}

	x.fields = append(x.fields, l.fields...)
	return x
}

// WithExecutor will set logger's executor.
func (l *Logger) WithExecutor(fn Executor) *Logger {
	l.executor = fn
	return l
}

// WithTransformer will set logger's transformer factory factory.
func (l *Logger) WithTransformer(t Transformer) *Logger {
	l.transformer = t
	return l
}

// WithFields will set logger's fields.
func (l *Logger) WithFields(fields ...Transformee) *Logger {
	l.fields = fields
	return l
}

// AddFields will append logger's fields, won't check whether it exists
func (l *Logger) AddFields(fields ...Transformee) *Logger {
	l.fields = append(l.fields, fields...)
	return l
}

// Print will print an empty entry.
func (l *Logger) Print(fields ...Transformee) {
	e := newEntry(level.Empty, fields...)
	defer e.Free()

	l.executor(l, e)
}

// Debug will print an debug entry.
func (l *Logger) Debug(fields ...Transformee) {
	e := newEntry(level.Debug, fields...)
	defer e.Free()

	l.executor(l, e)
}

// Info will print an info entry.
func (l *Logger) Info(fields ...Transformee) {
	e := newEntry(level.Info, fields...)
	defer e.Free()

	l.executor(l, e)
}

// Warn will print an warn entry.
func (l *Logger) Warn(fields ...Transformee) {
	e := newEntry(level.Warn, fields...)
	defer e.Free()

	l.executor(l, e)
}

// Error will print an error entry.
func (l *Logger) Error(fields ...Transformee) {
	e := newEntry(level.Error, fields...)
	defer e.Free()

	l.executor(l, e)
}

// Transform will do transform on this entry
//
// - Start and end an entry container
// - Transform all logger level fields
// - Transform all entry level fields.
func (l *Logger) Transform(e *Entry) {
	if e.transformed {
		return
	}

	l.transformer.Start(e, ContainerEntry)
	defer l.transformer.End(e, ContainerEntry)

	for _, v := range l.fields {
		v.Transform(l, e)
	}

	for _, v := range e.fields {
		v.Transform(l, e)
	}

	e.transformed = true
}

// WriteInto will write entry's bytes value into the writer.
//
// WriteTo looks much better, but requires to have signature WriteTo(io.Writer) (int64, error)
func (l *Logger) WriteInto(e *Entry, w io.Writer) {
	_, err := w.Write(e.buf.Bytes())
	if err != nil {
		buf := bufferpool.Get()
		defer buf.Free()

		buf.AppendString("log write failed: ")
		buf.AppendString(err.Error())
		_, _ = l.ew.Write(buf.Bytes())
	}
}
