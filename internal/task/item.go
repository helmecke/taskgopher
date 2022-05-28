package task

import (
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/helmecke/taskgopher/internal/parser"
	"github.com/helmecke/taskgopher/pkg/sliceutils"
	"github.com/helmecke/taskgopher/pkg/timeutils"
)

const (
	statusPending   = "pending"
	statusCompleted = "completed"
	statusDeleted   = "deleted"
)

// A Item is an item
type Item struct {
	ID          int       `json:"-"`
	UUID        uuid.UUID `json:"uuid"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified,omitempty"`
	Completed   time.Time `json:"completed,omitempty"`
	Due         time.Time `json:"due,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	Project     string    `json:"project,omitempty"`
	Contexts    []string  `json:"contexts,omitempty"`
	Urgency     float64   `json:"-"`
	VirtualTags []string  `json:"-"`
	Notes       []string  `json:"notes,omitempty"`
	filtered    bool      `json:"-"`
}

// NewTask is creating a new task
func NewTask(mod *parser.Modification) *Item {
	task := &Item{
		UUID:        uuid.New(),
		Description: mod.Description,
		// Due:         filter.Due,
		// Tags:        filter.Tags,
		// Contexts:    filter.Contexts,
		Status:  statusPending,
		Created: time.Now(),
	}

	return task
}

// Modify modifies task with given modification
func (t *Item) Modify(mod *parser.Modification) {
	t.Modified = time.Now()

	if mod.HasDescription() {
		t.Description = mod.Description
	}

	if mod.HasDue() {
		t.Due = mod.Due
	}

	if mod.RemoveDue {
		t.Due = time.Time{}
	}
}

// Complete completes task
func (t *Item) Complete() {
	now := time.Now()
	t.Completed = now
	t.Status = statusCompleted
}

// Delete deletes task
func (t *Item) Delete() {
	t.Status = statusDeleted
}

// SetUrgency sets urgency of task
func (t *Item) SetUrgency() {
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

	if !t.Due.IsZero() {
		urgency += u["due"]
	}

	urgency += math.Floor(u["age"] * (time.Since(t.Created).Hours() / 24 / 39))

	if len(t.Tags) > 0 {
		urgency += u["tags"]
	}

	if t.Project != "" {
		urgency += u["project"]
	}

	t.Urgency = urgency
}

// Age returns age of task in shorthand
func (t *Item) Age() string {
	return timeutils.Diff(time.Now(), t.Created, false)
}

// LastModifiedDiff returns duration since last modification in shorthand
func (t *Item) LastModifiedDiff() string {
	return timeutils.Diff(time.Now(), t.Modified, false)
}

// DueDiff return duration to due date in shorthand
func (t *Item) DueDiff() string {
	return timeutils.Diff(time.Now(), t.Due, true)
}

// GenerateVirtualTags generates virtual tags of task
// nolint:gocognit
func (t *Item) GenerateVirtualTags() {
	// DUE - Does the task have a due date?
	if !t.Due.IsZero() {
		t.VirtualTags = append(t.VirtualTags, "DUE")

		// OVERDUE - Is this task past its due date?
		if t.Due.Before(time.Now()) {
			t.VirtualTags = append(t.VirtualTags, "OVERDUE")
		}

		// YEAR	Is this task due this year?
		if t.Due.Year() == time.Now().Year() {
			t.VirtualTags = append(t.VirtualTags, "YEAR")

			// QUARTER	Is this task due this quarter?
			q := [12]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4}
			if q[t.Due.Month()] == q[time.Now().Month()] {
				t.VirtualTags = append(t.VirtualTags, "QUARTER")

				// MONTH - Is this task due this month?
				if t.Due.Month() == time.Now().Month() {
					t.VirtualTags = append(t.VirtualTags, "MONTH")

					// WEEK - Is this task due this week?
					_, dt := t.Due.ISOWeek()
					_, ct := time.Now().ISOWeek()
					if dt == ct {
						t.VirtualTags = append(t.VirtualTags, "WEEK")

						// TODAY - Is this task due sometime today?
						if t.Due.Equal(time.Now().Truncate(time.Hour * 24)) {
							t.VirtualTags = append(t.VirtualTags, "TODAY")
						}

						// YESTERDAY - Was the task due yesterday?
						if t.Due.Equal(time.Now().Truncate(time.Hour * 24).Add(-24 * time.Hour)) {
							t.VirtualTags = append(t.VirtualTags, "YESTERDAY")
						}

						// TOMORROW - Is the task due tomorrow?
						if t.Due.Equal(time.Now().Truncate(time.Hour * 24).Add(24 * time.Hour)) {
							t.VirtualTags = append(t.VirtualTags, "TOMORROW")
						}
					}
				}
			}
		}
	}

	// PENDING - Is the task in the pending state?
	if t.Status == statusPending {
		t.VirtualTags = append(t.VirtualTags, "PENDING")
	}

	// TAGGED - Does the task have any tags?
	if len(t.Tags) > 0 {
		t.VirtualTags = append(t.VirtualTags, "TAGGED")
	}

	// PROJECT - Does the task have a project?
	if t.Project != "" {
		t.VirtualTags = append(t.VirtualTags, "PROJECT")
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
func (t *Item) Matches(filter *parser.Filter) bool {
	if len(filter.IDs) > 0 && sliceutils.IntSliceContains(filter.IDs, t.ID) {
		return true
	}

	if filter.HasDue() && filter.Due.Equal(t.Due) {
		return true
	}

	return false
}