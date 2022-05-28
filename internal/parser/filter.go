package parser

import "time"

// Filter holds filter results
type Filter struct {
	Found bool
	IDs   []int
	Due   time.Time
	All   bool
}

// HasDue returns if filter has due date set
func (f *Filter) HasDue() bool {
	return !f.Due.IsZero()
}
