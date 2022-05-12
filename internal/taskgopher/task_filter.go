package taskgopher

// TaskFilter filters tasks based on patterns
type TaskFilter struct {
	Filter *Filter
	Tasks  []*Task
}

// ApplyFilter filters tasks based on the Filter struct passed in.
func (f *TaskFilter) ApplyFilter() (filtered []*Task) {
	for _, task := range f.Tasks {
		if f.Filter.HasTags {
			if !f.taskPassesFilter(task.Tags, f.Filter.Tags, f.Filter.ExcludeTags) {
				continue
			}
		}

		if f.Filter.HasContexts {
			if !f.taskPassesFilter(task.Contexts, f.Filter.Contexts, f.Filter.ExcludeContexts) {
				continue
			}
		}

		// has exact due date
		if f.Filter.HasDue {
			if task.Due != f.Filter.Due {
				continue
			}
		}

		// finally, if we're still here, append the task to the filtered list
		filtered = append(filtered, task)
	}

	return filtered
}

func (f *TaskFilter) taskPassesFilter(val []string, inclusiveVals []string, exclusiveVals []string) (passes bool) {
	// inclusive vals is evaluated via OR
	if len(inclusiveVals) >= 1 {
		passes = false

		for _, iv := range inclusiveVals {
			for _, v := range val {
				if iv == v {
					passes = true
				}
			}
		}
	}

	// exclusiveVals is evaluated with AND
	for _, ev := range exclusiveVals {
		for _, v := range val {
			if ev == v {
				passes = false
			}
		}
	}

	return passes
}
