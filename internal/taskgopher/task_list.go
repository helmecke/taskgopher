package taskgopher

// TaskList hold n task
type TaskList struct {
	Tasks []*Task
}

func (t *TaskList) load(tasks []*Task) {
	if t.Tasks != nil {
		t.Tasks = append(t.Tasks, tasks...)
	} else {
		t.Tasks = tasks
	}

	for _, task := range t.Tasks {
		task.urgency()
	}
}

func (t *TaskList) add(task *Task) int {
	task.ID = len(t.Tasks) + 1
	t.Tasks = append(t.Tasks, task)

	return task.ID
}

func (t *TaskList) get(id int) *Task {
	for _, task := range t.Tasks {
		if id == task.ID {
			return task
		}
	}

	return nil
}

func (t *TaskList) set(task *Task) {
	t.Tasks[task.ID-1] = task
}

func (t *TaskList) garbageCollect() (completed []*Task) {
	var tasks []*Task

	for _, task := range t.Tasks {
		if !(task.Status == "deleted" || task.Status == "completed") {
			tasks = append(tasks, task)
		}

		if task.Status == "completed" {
			completed = append(completed, task)
		}
	}
	t.Tasks = tasks

	return completed
}
