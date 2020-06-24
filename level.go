package log

// Level is the log level for logger.
type Level uint8

const (
	// EmptyLevel is the lowest level.
	EmptyLevel Level = iota

	// DebugLevel used for debugging.
	DebugLevel
	// InfoLevel used for printing informational data.
	InfoLevel
	// WarnLevel used for printing warning information.
	WarnLevel
	// ErrorLevel used for printing error.
	ErrorLevel

	// DisableLevel always the highest level.
	DisableLevel
)
