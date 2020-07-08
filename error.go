package log

import (
	"errors"
	"fmt"
)

var (
	// ErrTextConfigInvalidLayout means input layout is invalid.
	ErrTextConfigInvalidLayout = errors.New("text config invalid layout")
)

// NewTextConfigError will return a new TextConfigError
func NewTextConfigError(err error, tc *TextConfig) error {
	return &TextConfigError{
		Err: err,
		tc:  tc,
	}
}

// TextConfigError carries text config error.
type TextConfigError struct {
	Err error

	tc *TextConfig
}

func (e *TextConfigError) Error() string {
	return fmt.Sprintf("text_config: {entry_format: %s, level_format: %v, time_format: %s, value_container: %v}: %s", e.tc.EntryFormat, e.tc.LevelFormat, e.tc.TimeFormat, e.tc.ValueContainer, e.Err.Error())
}

// Unwrap implements xerrors.Wrapper
func (e *TextConfigError) Unwrap() error {
	return e.Err
}
