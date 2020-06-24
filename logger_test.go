package log

import (
	"testing"
)

func TestLogger_Info(t *testing.T) {
	l := New()

	l.Info(
		Int("version", 1024),
		String("msg", "test logger"),
	)
}
