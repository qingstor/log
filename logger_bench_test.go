package log

import (
	"io/ioutil"
	"testing"
	"time"
)

func benchmarkLogger(b *testing.B, f func(*Logger)) {
	m := MatchHigherLevel(InfoLevel)
	logger := New().WithExecutor(ExecuteMatchWrite(m, ioutil.Discard))

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
		logger.Info(
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
		logger.Info(
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
