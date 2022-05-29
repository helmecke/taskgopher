package parser

import (
	"time"

	"github.com/google/uuid"
)

// Filter holds filter results
type Filter struct {
	Found   bool
	IDs     []int
	UUIDs   []uuid.UUID
	Due     time.Time
	Project string
	All     bool
}

// HasDue returns if filter has due date set
func (f *Filter) HasDue() bool {
	return !f.Due.IsZero()
}
