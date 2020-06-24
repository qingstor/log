package log

import (
	"io"
)

// Executor will execute on an Entry.
type Executor func(e *Entry)

// ExecuteMultiple allow user execute multiple executor
func ExecuteMultiple(ex ...Executor) Executor {
	return func(e *Entry) {
		for _, fn := range ex {
			fn(e)
		}
	}
}

// ExecuteWrite creates an executor that write event into an io.Writer.
func ExecuteWrite(w io.Writer) Executor {
	return func(e *Entry) {
		e.Transform()
		e.WriteInto(w)
	}
}

// ExecuteMatchWrite creates an executor that write matched event into an io.Writer.
func ExecuteMatchWrite(m Matcher, w io.Writer) Executor {
	return func(e *Entry) {
		if !m.Match(e) {
			return
		}

		e.Transform()
		e.WriteInto(w)
	}
}
