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
	Notes       []string  `json:"notes,omitempty"`
}

// NewTask is creating a new task
func NewTask(filter *Filter) *Task {
	task := &Task{
		UUID:        uuid.New(),
		Description: filter.Description,
		Due:         filter.Due,
		Tags:        filter.Tags,
		Contexts:    filter.Contexts,
		Status:      statusPending,
		Created:     time.Now(),
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
