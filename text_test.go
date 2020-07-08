package log

import (
	"testing"
	"time"

	"github.com/qingstor/log/level"
)

func TestNextChunk(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectPre    string
		expectChunk  int
		expectSuffix string
	}{
		{
			"plain text",
			"hello",
			"hello",
			chunkPlain,
			"",
		},
		{
			"time chunk",
			"{time}",
			"",
			chunkTime,
			"",
		},
		{
			"level chunk",
			"{level}",
			"",
			chunkLevel,
			"",
		},
		{
			"value chunk",
			"{value}",
			"",
			chunkValue,
			"",
		},
		{
			"chunk with prefix and suffix",
			"abc {value} def",
			"abc ",
			chunkValue,
			" def",
		},
		{
			"return first chunk",
			"abc {value} {time} def",
			"abc ",
			chunkValue,
			" {time} def",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			actualPre, actualChunk, actualSuffix := nextChunk([]byte(v.input))
			if string(actualPre) != v.expectPre {
				t.Errorf("pre is not correct, expect: %s, actual: %s", v.expectPre, string(actualPre))
			}
			if actualChunk != v.expectChunk {
				t.Errorf("chunk is not correct, expect: %d, actual: %d", v.expectChunk, actualChunk)
			}
			if string(actualSuffix) != v.expectSuffix {
				t.Errorf("pre is not correct, expect: %s, actual: %s", v.expectSuffix, string(actualSuffix))
			}
		})
	}
}

func BenchmarkNextChunk(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _, _ = nextChunk([]byte("[{level}] - {time} {value}"))
		}
	})
}
func BenchmarkNewText(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = NewText(&TextConfig{
				EntryFormat:    "[{level}] - {time} {value}",
				LevelFormat:    level.UpperCase,
				TimeFormat:     time.RFC822,
				ValueContainer: nil,
			})
		}
	})
}
