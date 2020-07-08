package log

import (
	"sync"

	bp "github.com/Xuanwo/go-bufferpool"
	"github.com/qingstor/log/internal/pkg/bufferpool"
	"github.com/qingstor/log/level"
)

// eventPool is a type safe pool for event, only used internal.
var eventPool = sync.Pool{New: func() interface{} {
	return &Entry{}
}}

// newEntry will create a new event with level and fields.
func newEntry(lvl level.Level, fields ...Transformee) *Entry {
	e := eventPool.Get().(*Entry)

	e.Level = lvl
	e.buf = bufferpool.Get()
	e.fields = fields

	// Reset all runtime fields.
	e.transformed = false
	e.keys = 0
	e.containers = 0
	return e
}

// Entry represents a log entry.
type Entry struct {
	// Level is the entry's level, default to empty.
	Level level.Level
	// Fields is the entry level fields.
	fields []Transformee

	buf         *bp.Buffer // entry's internal buffer.
	transformed bool       // whether this entry has been transformed or not.
	keys        int        // only be used in transformer: how many key have been transformed.
	containers  int        // only be used in transformer: how many containers have been transformed.
}

// Free will free this entry.
func (e *Entry) Free() {
	// Deref files
	e.fields = nil
	// Deref buffer
	e.buf.Free()
	e.buf = nil

	eventPool.Put(e)
}
