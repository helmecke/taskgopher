package storage

import "github.com/helmecke/taskgopher/internal/task"

// Storage interface
type Storage interface {
	Load(bool) []*task.Item
	Save([]*task.Item)
	Complete(*task.Item)
}
