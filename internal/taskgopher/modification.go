package taskgopher

import "time"

// Modification holds modification results
type Modification struct {
	Description string
	Due         time.Time
	RemoveDue   bool
}

func (m *Modification) hasDescription() bool {
	return m.Description != ""
}

func (m *Modification) hasDue() bool {
	return !m.Due.IsZero()
}
