package taskgopher

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/helmecke/taskgopher/pkg/timeutils"
)

const (
	statusPending   = "pending"
	statusCompleted = "completed"
	statusDeleted   = "deleted"
	statusStarted   = "started"
	statusPaused    = "paused"
)

// A Task is an item
type Task struct {
	ID          int        `json:"-"`
	UUID        uuid.UUID  `json:""`
	Description string     `json:""`
	Status      string     `json:""`
	Created     *time.Time `json:""`
	Modified    *time.Time `json:",omitempty"`
	Completed   *time.Time `json:",omitempty"`
	Due         *time.Time `json:",omitempty"`
	Tags        []string   `json:",omitempty"`
	Project     string     `json:",omitempty"`
	Contexts    []string   `json:",omitempty"`
	Urgency     float64    `json:"-"`
	Notes       []string   `json:",omitempty"`
}

// NewTask returns a new task
func NewTask(filter *Filter) *Task {
	now := time.Now()
	task := &Task{
		UUID:        uuid.New(),
		Description: filter.Description,
		Due:         filter.Due,
		Tags:        filter.Tags,
		Contexts:    filter.Contexts,
		Status:      statusPending,
		Created:     &now,
	}

	task.urgency()

	return task

}

func EditTask(task *Task, filter *Filter) {
	now := time.Now()
	task.Modified = &now

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
	t.Completed = &now
	t.Status = statusCompleted
}

func (t *Task) delete() {
	t.Status = statusDeleted
}

func (t *Task) start() {
	t.Status = statusStarted
}

func (t *Task) stop() {
	t.Status = statusPaused
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

	if t.Due != nil {
		urgency += u["due"]
	}

	if t.Status == statusStarted {
		urgency += u[statusStarted]
	}

	urgency += math.Floor(u["age"] * (time.Since(*t.Created).Hours() / 24 / 39))

	if len(t.Tags) > 0 {
		urgency += u["tags"]
	}

	if t.Project != "" {
		urgency += u["project"]
	}

	t.Urgency = urgency
}

func (t *Task) table() []string {
	return []string{
		fmt.Sprint(t.ID),
		timeutils.Diff(time.Now(), *t.Created),
		t.Description,
		strconv.FormatFloat(t.Urgency, 'f', -1, 64),
	}
}
