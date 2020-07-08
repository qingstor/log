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
		TimeFormat:  TimeFormatUnixNano,
		LevelFormat: level.LowerCase,
		EntryFormat: "[{level}] - {time} {value}",
	})
	if err != nil {
		println("text transformer created failed for: ", err)
        os.Exit(1)
	}

    l := New().
        WithTransformer(tf).
        WithFields(
            String("request_id", "abcdefg"),
        )
    
    l.Info(
        log.Int("version", 1024),
        log.String("msg", "test logger")
    )

    // [info] - 1594174591970815623 request_id="abcdefg" version=1024 msg="test logger"
}
```