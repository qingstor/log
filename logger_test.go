package log

import (
	"testing"
	"time"

	"github.com/qingstor/log/level"
)

func TestLogger_Info(t *testing.T) {
	tf, err := NewText(&TextConfig{
		TimeFormat:  time.RFC3339,
		LevelFormat: level.LowerCase,
		EntryFormat: "[{level}] - {time} {value}",
	})
	if err != nil {
		t.Errorf("new text failed for %v", err)
	}

	l := New().WithTransformer(tf)

	l.Info(
		Int("version", 1024),
		String("msg", "test logger"),
	)
}

func TestLogger_InfoWithLoggerEntry(t *testing.T) {
	tf, err := NewText(&TextConfig{
		TimeFormat:  TimeFormatUnixNano,
		LevelFormat: level.LowerCase,
		EntryFormat: "[{level}] - {time} {value}",
	})
	if err != nil {
		t.Errorf("new text failed for %v", err)
	}

	l := New().
		WithTransformer(tf).
		WithFields(
			String("request_id", "abcdefg"),
		)

	l.Info(
		Int("version", 1024),
		String("msg", "test logger"),
	)
}
