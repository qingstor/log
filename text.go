package log

import (
	"bytes"
	"time"

	"github.com/qingstor/log/level"
)

const (
	chunkInvalid = iota
	chunkPlain
	chunkTime
	chunkLevel
	chunkValue
)

// nextChunk will parse layout to related chunk.
func nextChunk(layout []byte) (pre []byte, chunk int, suffix []byte) {
	startIdx := bytes.IndexByte(layout, '{')
	if startIdx == -1 {
		return layout, chunkPlain, nil
	}
	endIdx := bytes.IndexByte(layout, '}')
	if endIdx == -1 {
		return nil, chunkInvalid, nil
	}

	switch string(layout[startIdx+1 : endIdx]) {
	case "time":
		return layout[:startIdx], chunkTime, layout[endIdx+1:]
	case "level":
		return layout[:startIdx], chunkLevel, layout[endIdx+1:]
	case "value":
		return layout[:startIdx], chunkValue, layout[endIdx+1:]
	default:
		return layout, chunkInvalid, nil
	}
}

// TextConfig is the config of text transformer.
type TextConfig struct {
	// Whole log format
	// "{level} - {time} {value}"
	EntryFormat string
	// ERROR
	LevelFormat level.FormatCase
	// 1136239445 or time layout
	TimeFormat string
	// a container to byte map.
	ValueContainer map[Container][2]byte
}

// NewText will parse text config and create a new Transformer.
func NewText(tc *TextConfig) (Transformer, error) {
	tx := &Text{}

	// Parse value container.
	tx.vc = [][2]byte{
		ContainerObject: {'{', '}'},
		ContainerMap:    {'{', '}'},
		ContainerArray:  {'[', ']'},
		ContainerQuote:  {'"', '"'},
	}
	for k, v := range tc.ValueContainer {
		tx.vc[k][0], tx.vc[k][1] = v[0], v[1]
	}

	// Parse level layout.
	levelFn := func(e *Entry) {
		e.buf.AppendString(level.Format[tc.LevelFormat][e.Level])
	}

	// Parse time layout.
	var timeFn func(e *Entry)
	switch tc.TimeFormat {
	case TimeFormatUnix:
		timeFn = func(e *Entry) {
			e.buf.AppendInt(time.Now().Unix())
		}
	case TimeFormatUnixNano:
		timeFn = func(e *Entry) {
			e.buf.AppendInt(time.Now().UnixNano())
		}
	default:
		timeFn = func(e *Entry) {
			e.buf.AppendTime(time.Now(), tc.TimeFormat)
		}
	}

	// Parse entry layout.
	layout := []byte(tc.EntryFormat)
	fns := make([]func(e *Entry), 0)
	valuePos := 0
	for len(layout) != 0 {
		pre, chunk, suffix := nextChunk(layout)

		if chunk == chunkInvalid {
			return nil, NewTextConfigError(ErrTextConfigInvalidLayout, tc)
		}

		if len(pre) != 0 {
			fns = append(fns, func(e *Entry) {
				e.buf.AppendString(string(pre))
			})
		}

		switch chunk {
		case chunkTime:
			fns = append(fns, timeFn)
		case chunkLevel:
			fns = append(fns, levelFn)
		case chunkValue:
			// Currently, we don't need to do anything on value chunk, but we do need to record the position.
			valuePos = len(fns)
		}

		layout = suffix
	}

	tx.pre = func(e *Entry) {
		for _, v := range fns[:valuePos] {
			v(e)
		}
	}
	tx.post = func(e *Entry) {
		for _, v := range fns[valuePos:] {
			v(e)
		}
		e.buf.AppendByte('\n')
	}
	return tx, nil
}

// Text is the transform both human and machine friendly.
type Text struct {
	// Value containers
	vc [][2]byte

	pre  func(*Entry)
	post func(*Entry)
}

// Key will append new key.
func (t *Text) Key(e *Entry, key string) {
	if e.keys > e.containers {
		e.buf.AppendByte(' ')
	} else {
		e.keys++
	}

	e.buf.AppendString(key)
	// TODO: Allow user set via value format?
	e.buf.AppendByte('=')
}

// Start will pre a container.
func (t *Text) Start(e *Entry, tc Container) {
	switch tc {
	case ContainerEntry:
		t.pre(e)
		// Entry container should not be counted in containers, return here directly.
		return
	case ContainerObject:
		e.buf.AppendByte(t.vc[ContainerObject][0])
	case ContainerMap:
		e.buf.AppendByte(t.vc[ContainerMap][0])
	case ContainerArray:
		e.buf.AppendByte(t.vc[ContainerArray][0])
	case ContainerQuote:
		e.buf.AppendByte(t.vc[ContainerQuote][0])
	}

	e.containers++

}

// End will post a container.
func (t *Text) End(e *Entry, tc Container) {
	switch tc {
	case ContainerEntry:
		t.post(e)
		// Entry container should not be counted in containers, return here directly.
		return
	case ContainerObject:
		e.buf.AppendByte(t.vc[ContainerObject][1])
	case ContainerMap:
		e.buf.AppendByte(t.vc[ContainerMap][1])
	case ContainerArray:
		e.buf.AppendByte(t.vc[ContainerArray][1])
	case ContainerQuote:
		e.buf.AppendByte(t.vc[ContainerQuote][1])
	}

	e.containers--
}
