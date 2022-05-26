package taskgopher

// TaskList hold n task
type TaskList struct {
	Tasks []*Task
}

func (tl *TaskList) load(tasks []*Task) {
	if tl.Tasks != nil {
		tl.Tasks = append(tl.Tasks, tasks...)
	} else {
		tl.Tasks = tasks
	}

	for _, task := range tl.Tasks {
		task.urgency()
		task.generateVirtualTags()
	}
}

func (tl *TaskList) add(task *Task) int {
	task.ID = len(tl.Tasks) + 1
	tl.Tasks = append(tl.Tasks, task)

	return task.ID
}

func (tl *TaskList) get(id int) *Task {
	for _, task := range tl.Tasks {
		if id == task.ID {
			return task
		}
	}

	return nil
}

func (tl *TaskList) set(task *Task) {
	tl.Tasks[task.ID-1] = task
}

func (tl *TaskList) garbageCollect() (completed []*Task) {
	var tasks []*Task

	for _, task := range tl.Tasks {
		if !(task.Status == "deleted" || task.Status == "completed") {
			tasks = append(tasks, task)
		}

		if task.Status == "completed" {
			completed = append(completed, task)
		}
	}
	tl.Tasks = tasks

	return completed
}

func (tl *TaskList) filter(filter *Filter) {
	for _, task := range tl.Tasks {
		if task.matches(filter) || !filter.Found {
			task.filtered = true
		}
	}
}

func (tl *TaskList) filtered() (tasks []*Task) {
	for _, task := range tl.Tasks {
		if task.filtered {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
