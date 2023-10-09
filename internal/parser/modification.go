package parser

import "time"

// Modification holds modification results
type Modification struct {
	Description   string
	Due           time.Time
	Project       string
	RemoveDue     bool
	RemoveProject bool
}

// HasDescription is a helper if modification has description
func (m *Modification) HasDescription() bool {
	return m.Description != ""
}

// HasDue is a helper if modification has due date
func (m *Modification) HasDue() bool {
	return !m.Due.IsZero()
}

// HasProject is a helper if modification has project
func (m *Modification) HasProject() bool {
	return m.Project != ""
}
