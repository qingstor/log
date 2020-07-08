# log

A log designed for critical mission.

## Goals

- Safe
  - Type Safe: No reflection
  - Runtime Safe: No panic, no error check
- Performance
  - Low Latency: Log entry with 10 fields in 1ms
  - Low Footprint: Nearly no extra objects allocated


## Quick Start

```go
import (
    "os"

    "github.com/qingstor/log"
)

func main()  {
	tf, err := NewText(&TextConfig{
		// Use unix timestamp nano for time
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

    // [INFO] - 1594174591970815623 request_id="8da3aceea1ba" object_key="test_object" version=3
}
```