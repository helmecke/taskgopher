package task

import (
	"github.com/google/uuid"

	"github.com/helmecke/taskgopher/internal/parser"
)

// List hold n task
type List struct {
	Tasks []*Item
}

// Load adds multiple task to list
func (l *List) Load(tasks []*Item) {
	if l.Tasks != nil {
		l.Tasks = append(l.Tasks, tasks...)
	} else {
		l.Tasks = tasks
	}

	for _, task := range l.Tasks {
		task.SetUrgency()
		task.GenerateVirtualTags()
	}
}

// Add adds task to list
func (l *List) Add(task *Item) int {
	task.ID = len(l.Tasks) + 1
	l.Tasks = append(l.Tasks, task)

	return task.ID
}

// GetByID gets task with id
func (l *List) GetByID(id int) *Item {
	for _, task := range l.Tasks {
		if id == task.ID {
			return task
		}
	}

	return nil
}

// GetByUUID gets task with uuid
func (l *List) GetByUUID(uuid uuid.UUID) *Item {
	for _, task := range l.Tasks {
		if uuid == task.UUID {
			return task
		}
	}

	return nil
}

// Set places task at ID
func (l *List) Set(task *Item) {
	l.Tasks[task.ID-1] = task
}

// GarbageCollect moves completed task
func (l *List) GarbageCollect() (completed []*Item) {
	var tasks []*Item

	for _, task := range l.Tasks {
		if !(task.Status == "deleted" || task.Status == "completed") {
			tasks = append(tasks, task)
		}

		if task.Status == "completed" {
			completed = append(completed, task)
		}
	}
	l.Tasks = tasks

	return completed
}

// Filter filters task matching given filter
func (l *List) Filter(filter *parser.Filter) {
	for _, task := range l.Tasks {
		if task.Matches(filter) || !filter.Found {
			task.filtered = true
		}
	}
}

// Filtered return all filtered task
func (l *List) Filtered() (tasks []*Item) {
	for _, task := range l.Tasks {
		if task.filtered {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
