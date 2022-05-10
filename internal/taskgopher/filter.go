package taskgopher

import (
	"time"
)

// Filter holds the filtering results
type Filter struct {
	Description string
	HasDue      bool
	Due         *time.Time
	Contexts    []string
	Tags        []string
}
