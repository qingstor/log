package log

import (
	"github.com/qingstor/log/level"
)

// Matcher used for match an entry.
type Matcher interface {
	Match(e *Entry) bool
}

type matchHigherLevel struct {
	lvl level.Level
}

// MatchHigherLevel used to match an Entry whose level is higher than lvl.
func MatchHigherLevel(lvl level.Level) Matcher {
	return matchHigherLevel{lvl: lvl}
}

// Match implements Matcher.
func (m matchHigherLevel) Match(e *Entry) bool {
	return e.Level > m.lvl
}

type matchLowerLevel struct {
	lvl level.Level
}

// MatchLowerLevel used to match an Entry whose level is lower than lvl.
func MatchLowerLevel(lvl level.Level) Matcher {
	return matchLowerLevel{lvl: lvl}
}

// Match implements Matcher.
func (m matchLowerLevel) Match(e *Entry) bool {
	return e.Level < m.lvl
}
