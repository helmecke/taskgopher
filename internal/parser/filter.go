package parser

import (
	"time"

	"github.com/google/uuid"
)

// Filter holds filter results
type Filter struct {
	IDs     []int
	Due     time.Time
	Project string
	UUIDs   []uuid.UUID
	Found   bool
	All     bool
}

// HasDue returns true if filter has due date set
func (f *Filter) HasDue() bool {
	return !f.Due.IsZero()
}

// HasProject returns true if filter has project set
func (f *Filter) HasProject() bool {
	return len(f.Project) > 0
}
