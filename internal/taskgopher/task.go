package taskgopher

import (
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/helmecke/taskgopher/pkg/timeutils"
)

const (
	statusPending   = "pending"
	statusCompleted = "completed"
	statusDeleted   = "deleted"
)

// A Task is an item
type Task struct {
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
}

// NewTask is creating a new task
func NewTask(_ *Filter) *Task {
	task := &Task{
		UUID: uuid.New(),
		// Description: filter.Description,
		// Due:         filter.Due,
		// Tags:        filter.Tags,
		// Contexts:    filter.Contexts,
		Status:  statusPending,
		Created: time.Now(),
	}

	return task
}

// EditTask is modifying a existing task
func EditTask(task *Task, filter *Filter) {
	now := time.Now()
	task.Modified = now

	if filter.HasDue {
		task.Due = filter.Due
	}
	if filter.Description != "" {
		task.Description = filter.Description
	}

	if len(filter.Tags) > 0 {
		task.Tags = filter.Tags
	}

	if len(filter.Contexts) > 0 {
		task.Contexts = filter.Contexts
	}
}

func (t *Task) complete() {
	now := time.Now()
	t.Completed = now
	t.Status = statusCompleted
}

func (t *Task) delete() {
	t.Status = statusDeleted
}

func (t *Task) urgency() {
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

func (t *Task) age() string {
	return timeutils.Diff(time.Now(), t.Created, false)
}

func (t *Task) lastModified() string {
	return timeutils.Diff(time.Now(), t.Modified, false)
}

func (t *Task) due() string {
	return timeutils.Diff(time.Now(), t.Due, true)
}

// nolint:gocognit
func (t *Task) generateVirtualTags() {
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
