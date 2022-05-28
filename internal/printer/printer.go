package printer

import (
	"github.com/helmecke/taskgopher/internal/task"
)

// Printer interface
type Printer interface {
	PrintItem(*task.Item)
	PrintList([]*task.Item)
}
