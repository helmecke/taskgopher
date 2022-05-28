package parser

import "time"

// Modification holds modification results
type Modification struct {
	Description string
	Due         time.Time
	RemoveDue   bool
}

// HasDescription returns if modification has description set
func (m *Modification) HasDescription() bool {
	return m.Description != ""
}

// HasDue returns if modification has due date set
func (m *Modification) HasDue() bool {
	return !m.Due.IsZero()
}
