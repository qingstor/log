package log

import (
	"io"
	"sync"

	"github.com/qingstor/log/internal/pkg/bufferpool"
	"github.com/qingstor/log/transform"
)

// eventPool is a type safe pool for event, only used internal.
var eventPool = sync.Pool{New: func() interface{} {
	return &Entry{}
}}

// newEntry will create a new event with level and fields.
func newEntry(l *Logger, lvl Level, fields ...transform.Transformee) *Entry {
	e := eventPool.Get().(*Entry)

	e.Level = lvl

	e.l = l
	e.fields = fields
	return e
}

// Entry represents a log entry.
type Entry struct {
	// Level is the entry's level, default to empty.
	Level Level

	// Keep a reference to Logger, so that we can access logger's variable.
	l *Logger
	// Every entry has a transformer for transform.
	m transform.Transformer
	// Fields is the entry level fields.
	fields []transform.Transformee
}

// Transform will transform entry's fields into transformer.
func (e *Entry) Transform() {
	if e.m != nil {
		return
	}

	e.m = e.l.transform()

	for _, v := range e.l.fields {
		v.Transform(e.m)
	}

	for _, v := range e.fields {
		v.Transform(e.m)
	}

	e.m.AppendString("\n")
}

// WriteInto will write entry's bytes value into the writer.
//
// WriteTo looks much better, but requires to have signature WriteTo(io.Writer) (int64, error)
func (e *Entry) WriteInto(w io.Writer) {
	_, err := w.Write(e.m.Bytes())
	if err != nil {
		buf := bufferpool.Get()
		defer buf.Free()

		buf.AppendString("l write failed: ")
		buf.AppendString(err.Error())
		_, _ = e.l.ew.Write(buf.Bytes())
	}
}

// Free will free this entry.
func (e *Entry) Free() {
	e.l = nil
	e.fields = nil

	if e.m != nil {
		e.m.Free()
		e.m = nil
	}

	eventPool.Put(e)
}
