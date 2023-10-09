package task

import (
	"math"
	"slices"
	"time"

	"github.com/google/uuid"

	"github.com/helmecke/taskgopher/internal/parser"
	"github.com/helmecke/taskgopher/pkg/timeutils"
)

const (
	statusPending   = "pending"
	statusCompleted = "completed"
	statusDeleted   = "deleted"
)

// A Item is an item
type Item struct {
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified,omitempty"`
	Completed   time.Time `json:"completed,omitempty"`
	Due         time.Time `json:"due,omitempty"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Project     string    `json:"project,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	VirtualTags []string  `json:"-"`
	Notes       []string  `json:"notes,omitempty"`
	UUID        uuid.UUID `json:"uuid"`
	Urgency     float64   `json:"-"`
	ID          int       `json:"-"`
	filtered    bool      `json:"-"`
}

// NewTask is creating a new task
func NewTask(mod *parser.Modification) *Item {
	task := &Item{
		UUID:        uuid.New(),
		Description: mod.Description,
		// Due:         filter.Due,
		// Tags:        filter.Tags,
		Status:  statusPending,
		Created: time.Now(),
	}

	return task
}

// Modify modifies task with given modification
func (i *Item) Modify(mod *parser.Modification) {
	i.Modified = time.Now()

	if mod.HasDescription() {
		i.Description = mod.Description
	}

	if mod.HasDue() {
		i.Due = mod.Due
	}

	if mod.RemoveDue {
		i.Due = time.Time{}
	}

	if mod.HasProject() {
		i.Project = mod.Project
	}

	if mod.RemoveProject {
		i.Project = ""
	}
}

// Complete completes task
func (i *Item) Complete() {
	now := time.Now()
	i.Completed = now
	i.Status = statusCompleted
}

// Delete deletes task
func (i *Item) Delete() {
	i.Status = statusDeleted
}

// SetUrgency sets urgency of task
func (i *Item) SetUrgency() {
	u := map[string]float64{
		"next":      15,
		"due":       12,
		"blocking":  8,
		"high":      6,
		"medium":    3.9,
		"low":       1.8,
		"scheduled": 5,
		"started":   4,
		"age":       2,
		"tags":      1,
		"project":   1,
		"waiting":   -3,
		"blocked":   -5,
	}

	urgency := 0.0

	if !i.Due.IsZero() {
		urgency += u["due"]
	}

	urgency += math.Floor(u["age"] * (time.Since(i.Created).Hours() / 24 / 39))

	if len(i.Tags) > 0 {
		urgency += u["tags"]
	}

	if i.Project != "" {
		urgency += u["project"]
	}

	i.Urgency = urgency
}

// Age returns age of task in shorthand
func (i *Item) Age() string {
	return timeutils.Diff(time.Now(), i.Created, false)
}

// LastModifiedDiff returns duration since last modification in shorthand
func (i *Item) LastModifiedDiff() string {
	if i.Modified.IsZero() {
		return ""
	}
	return timeutils.Diff(time.Now(), i.Modified, false)
}

// DueDiff return duration to due date in shorthand
func (i *Item) DueDiff() string {
	if i.Due.IsZero() {
		return ""
	}
	return timeutils.Diff(time.Now(), i.Due, false)
}

// GenerateVirtualTags generates virtual tags of task
// nolint:gocognit
func (i *Item) GenerateVirtualTags() {
	// DUE - Does the task have a due date?
	if !i.Due.IsZero() {
		i.VirtualTags = append(i.VirtualTags, "DUE")

		// OVERDUE - Is this task past its due date?
		if i.Due.Before(time.Now()) {
			i.VirtualTags = append(i.VirtualTags, "OVERDUE")
		}

		// YEAR	Is this task due this year?
		if i.Due.Year() == time.Now().Year() {
			i.VirtualTags = append(i.VirtualTags, "YEAR")

			// QUARTER	Is this task due this quarter?
			q := [12]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4}
			if q[i.Due.Month()] == q[time.Now().Month()] {
				i.VirtualTags = append(i.VirtualTags, "QUARTER")

				// MONTH - Is this task due this month?
				if i.Due.Month() == time.Now().Month() {
					i.VirtualTags = append(i.VirtualTags, "MONTH")

					// WEEK - Is this task due this week?
					_, dt := i.Due.ISOWeek()
					_, ct := time.Now().ISOWeek()
					if dt == ct {
						i.VirtualTags = append(i.VirtualTags, "WEEK")

						// TODAY - Is this task due sometime today?
						if i.Due.Equal(time.Now().Truncate(time.Hour * 24)) {
							i.VirtualTags = append(i.VirtualTags, "TODAY")
						}

						// YESTERDAY - Was the task due yesterday?
						if i.Due.Equal(time.Now().Truncate(time.Hour * 24).Add(-24 * time.Hour)) {
							i.VirtualTags = append(i.VirtualTags, "YESTERDAY")
						}

						// TOMORROW - Is the task due tomorrow?
						if i.Due.Equal(time.Now().Truncate(time.Hour * 24).Add(24 * time.Hour)) {
							i.VirtualTags = append(i.VirtualTags, "TOMORROW")
						}
					}
				}
			}
		}
	}

	// PENDING - Is the task in the pending state?
	if i.IsPending() {
		i.VirtualTags = append(i.VirtualTags, "PENDING")
	}

	// TAGGED - Does the task have any tags?
	if i.HasTag() {
		i.VirtualTags = append(i.VirtualTags, "TAGGED")
	}

	// PROJECT - Does the task have a project?
	if i.HasProject() {
		i.VirtualTags = append(i.VirtualTags, "PROJECT")
	}

	// BLOCKED - Is the task dependent on another incomplete task?
	// UNBLOCKED - The opposite of BLOCKED, for convenience.
	// BLOCKING - Does another task depend on this incomplete task?
	// ACTIVE - Is the task active, ie does it have a start date?
	// SCHEDULED - Is the task scheduled, ie does it have a scheduled date?
	// PARENT - Is the task a hidden parent recurring task?
	// CHILD - Is the task a recurring child task?
	// UNTIL - Does the task expire, ie does it have an until date?
	// WAITING - Is the task hidden, ie does it have a wait date?
	// ANNOTATED - Does the task have any annotations?
	// READY - Is the task pending, not blocked, and either not scheduled, or scheduled before now.
	// COMPLETED - Is the task in the completed state?
	// DELETED - Is the task in the deleted state?
	// UDA - Does the task contain any UDA values?
	// ORPHAN - Does the task contain any orphaned UDA values?
	// PRIORITY - Does the task have a priority?
	// LATEST - Is the task the most recently added task?
}

// Matches return if task matches given filter
func (i *Item) Matches(filter *parser.Filter) bool {
	if len(filter.IDs) > 0 && slices.Contains(filter.IDs, i.ID) {
		return true
	}

	if len(filter.UUIDs) > 0 && slices.Contains(filter.UUIDs, i.UUID) {
		return true
	}

	if filter.HasDue() && filter.Due.Equal(i.Due) {
		return true
	}

	if filter.HasProject() && filter.Project == i.Project {
		return true
	}

	return false
}

// IsPending is a helper if a task is pending
func (i *Item) IsPending() bool {
	return i.Status == statusPending
}

// HasTag is a helper if a task has a tag
func (i *Item) HasTag() bool {
	return len(i.Tags) > 0
}

// HasProject is a helper if a task has a project
func (i *Item) HasProject() bool {
	return i.Project != ""
}
