package log

// Matcher used for match an entry.
type Matcher interface {
	Match(e *Entry) bool
}

type matchHigherLevel struct {
	lvl Level
}

// MatchHigherLevel used to match an Entry whose level is higher than lvl.
func MatchHigherLevel(lvl Level) Matcher {
	return matchHigherLevel{lvl: lvl}
}

// Match implements Matcher.
func (m matchHigherLevel) Match(e *Entry) bool {
	return e.Level > m.lvl
}

type matchLowerLevel struct {
	lvl Level
}

// MatchLowerLevel used to match an Entry whose level is lower than lvl.
func MatchLowerLevel(lvl Level) Matcher {
	return matchLowerLevel{lvl: lvl}
}

// Match implements Matcher.
func (m matchLowerLevel) Match(e *Entry) bool {
	return e.Level < m.lvl
}
