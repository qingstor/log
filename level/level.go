package level

// Level is the log level for logger.
type Level uint8

const (
	// Empty is the lowest level.
	Empty Level = iota

	// Debug used for debugging.
	Debug
	// Info used for printing informational data.
	Info
	// Warn used for printing warning information.
	Warn
	// Error used for printing error.
	Error

	// Disable always the highest level.
	Disable
)

// FormatCase is the level format case for logger.
type FormatCase uint8

const (
	// PascalCase will make level looks like "Debug"
	PascalCase FormatCase = iota
	// LowerCase will make level looks like "debug"
	LowerCase
	// UpperCase Will make level looks like "DEBUG"
	UpperCase
)

// Format is the format slice used in transform.
var Format = [][]string{
	PascalCase: {
		Debug: "Debug",
		Info:  "Info",
		Warn:  "Warn",
		Error: "Error",
	},
	LowerCase: {
		Debug: "debug",
		Info:  "info",
		Warn:  "warn",
		Error: "error",
	},
	UpperCase: {
		Debug: "DEBUG",
		Info:  "INFO",
		Warn:  "WARN",
		Error: "ERROR",
	},
}
