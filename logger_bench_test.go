package log

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/qingstor/log/level"
)

func benchmarkLogger(b *testing.B, f func(*Logger)) {
	tf, _ := NewText(&TextConfig{
		TimeFormat:  "1136239445",
		EntryFormat: "{value}",
	})
	m := MatchHigherLevel(level.Info)
	logger := New().
		WithExecutor(ExecuteMatchWrite(m, ioutil.Discard)).
		WithTransformer(tf)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			f(logger)
		}
	})
}

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = New()
		}
	})
}

func BenchmarkLogger_NoOutput(b *testing.B) {
	now := time.Now()
	x := "Hello, world."

	benchmarkLogger(b, func(logger *Logger) {
		logger.Info(
			String("string", x),
			Bytes("bytes", []byte(x)),
			Int("int64", 1234567890),
			Float("float64", 1234.056789),
			Time("time", now, time.RFC1123),
		)
	})
}

func BenchmarkLogger_Info(b *testing.B) {
	now := time.Now()
	x := "Hello, world."

	benchmarkLogger(b, func(logger *Logger) {
		logger.Error(
			String("string", x),
			Bytes("bytes", []byte(x)),
			Int("int64", 1234567890),
			Float("float64", 1234.056789),
			Time("time", now, time.RFC1123),
		)
	})
}

func BenchmarkLogger_Int(b *testing.B) {
	benchmarkLogger(b, func(logger *Logger) {
		logger.Error(
			Int("int64", 1234567890),
			Int("int64", 1234567890),
			Int("int64", 1234567890),
			Int("int64", 1234567890),
			Int("int64", 1234567890),
			Int("int64", 1234567890),
			Int("int64", 1234567890),
			Int("int64", 1234567890),
			Int("int64", 1234567890),
			Int("int64", 1234567890),
			Int("int64", 1234567890),
		)
	})
}
