package taskgopher

// Store is an interface to store tasks
type Store interface {
	init()
	load(bool) []*Task
	save([]*Task)
	complete(*Task)
}
