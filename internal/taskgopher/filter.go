package taskgopher

import (
	"time"
)

// Filter holds the filtering results
type Filter struct {
	IDs             []int
	Description     string
	HasDue          bool
	HasContexts     bool
	HasTags         bool
	Due             time.Time
	Contexts        []string
	ExcludeContexts []string
	Tags            []string
	ExcludeTags     []string
	All             bool
}
