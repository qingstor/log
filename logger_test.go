package log

import (
	"os"
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

func ExampleLogger_Info() {
	tf, err := NewText(&TextConfig{
		// Use unix timestamp for time
		TimeFormat: TimeFormatUnixNano,
		// Use upper case level
		LevelFormat: level.UpperCase,
		EntryFormat: "[{level}] - {time} {value}",
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
