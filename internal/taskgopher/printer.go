package taskgopher

type Printer interface {
	PrintTaskList([]*Task)
}
