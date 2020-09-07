package log

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/qingstor/log/level"
)

func TestLogger_Info(t *testing.T) {
	tf, err := NewText(&TextConfig{
		TimeFormat:  time.RFC3339,
		LevelFormat: level.LowerCase,
		EntryFormat: defaultFormat,
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
		EntryFormat: defaultFormat,
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

func TestLogger_AddFields(t *testing.T) {
	var buf bytes.Buffer
	tf, err := NewText(&TextConfig{
		TimeFormat:  TimeFormatUnixNano,
		LevelFormat: level.LowerCase,
		EntryFormat: defaultFormat,
	})
	if err != nil {
		t.Errorf("new text failed for %v", err)
	}

	e := ExecuteWrite(&buf)

	l := New().
		WithTransformer(tf).
		WithExecutor(e).
		AddFields(
			String("request_id", "abcdefg"),
		).
		AddFields(
			String("request_id", "gfedcba"),
		)

	l.Info(
		Int("version", 1024),
		String("msg", "test logger"),
	)

	if c := strings.Count(buf.String(), "request_id"); c != 2 {
		t.Errorf("expected contains 2 request_id, only got %d", c)
	}
}

func ExampleLogger_Info() {
	tf, err := NewText(&TextConfig{
		// Use unix timestamp for time
		TimeFormat: TimeFormatUnixNano,
		// Use upper case level
		LevelFormat: level.UpperCase,
		EntryFormat: defaultFormat,
	})
	if err != nil {
		println("text config created failed for: ", err)
		os.Exit(1)
	}

	e := ExecuteMatchWrite(
		// Only print log that level is higher than Debug.
		MatchHigherLevel(level.Debug),
		// Write into stderr.
		os.Stderr,
	)

	logger := New().
		WithExecutor(e).
		WithTransformer(tf).
		WithFields(
			String("request_id", "8da3aceea1ba"),
		)

	logger.Info(
		String("object_key", "test_object"),
		Int("version", 3),
	)
}
