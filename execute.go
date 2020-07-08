package log

import (
	"io"
)

// Executor will execute on an Entry.
type Executor func(l *Logger, e *Entry)

// ExecuteMultiple allow user execute multiple executor
func ExecuteMultiple(ex ...Executor) Executor {
	return func(l *Logger, e *Entry) {
		for _, fn := range ex {
			fn(l, e)
		}
	}
}

// ExecuteWrite creates an executor that write event into an io.Writer.
func ExecuteWrite(w io.Writer) Executor {
	return func(l *Logger, e *Entry) {
		l.Transform(e)
		l.WriteInto(e, w)
	}
}

// ExecuteMatchWrite creates an executor that write matched event into an io.Writer.
func ExecuteMatchWrite(m Matcher, w io.Writer) Executor {
	return func(l *Logger, e *Entry) {
		if !m.Match(e) {
			return
		}

		l.Transform(e)
		l.WriteInto(e, w)
	}
}
