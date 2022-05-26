package taskgopher

import "time"

// Filter holds filter results
type Filter struct {
	Found bool
	IDs   []int
	Due   time.Time
	All   bool
}

func (f *Filter) hasDue() bool {
	return !f.Due.IsZero()
}
