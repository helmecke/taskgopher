package taskgopher

type Printer interface {
	PrintTask(*Task)
	PrintTaskList([]*Task)
}
