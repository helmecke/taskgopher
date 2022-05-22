package taskgopher

// Printer is an interface to print tasks
type Printer interface {
	PrintTask(*Task)
	PrintTaskList([]*Task)
}
