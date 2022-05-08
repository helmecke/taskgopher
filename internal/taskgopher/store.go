package taskgopher

type store interface {
	init()
	load(bool) []*Task
	save([]*Task)
	complete(*Task)
}
