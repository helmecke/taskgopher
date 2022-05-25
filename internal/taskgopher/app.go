package taskgopher

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// App is the structure of the taskgopher app
type App struct {
	TaskList TaskList
	Store    Store
	Printer  Printer
}

// NewApp is creating the taskgopher app
func NewApp(location string) *App {
	return &App{
		Store:   newFileStore(location),
		Printer: NewScreenPrinter(),
	}
}

// AddTask is creating a task
func (a *App) AddTask(filter *Filter) error {
	a.load(false)

	task := NewTask(filter)

	fmt.Printf("Created task %d.\n", a.TaskList.add(task))
	a.save()

	return nil
}

// ModifyTask is modifying a task
func (a *App) ModifyTask(filter *Filter) error {
	a.load(filter.All)

	taskFilter := &TaskFilter{Tasks: a.TaskList.Tasks, Filter: filter}
	tasks := taskFilter.ApplyFilter()

	for _, task := range tasks {
		EditTask(task, filter)

		a.TaskList.set(task)
		fmt.Printf("Modified task %d.\n", task.ID)
	}
	a.save()

	return nil
}

// CompleteTask is completing a task
func (a *App) CompleteTask(filter *Filter) error {
	a.load(filter.All)

	taskFilter := &TaskFilter{Tasks: a.TaskList.Tasks, Filter: filter}
	tasks := taskFilter.ApplyFilter()

	for _, task := range tasks {
		task.complete()

		a.TaskList.set(task)
		fmt.Printf("Completed task %d.\n", task.ID)
	}
	a.save()

	return nil
}

// DeleteTask is deleting a task
func (a *App) DeleteTask(filter *Filter) error {
	prompt := promptui.Prompt{
		Label:     "Delete task",
		IsConfirm: true,
	}

	if _, err := prompt.Run(); err != nil {
		fmt.Println("Aborted...")

		// nolint:nilerr
		return nil
	}

	a.load(filter.All)

	taskFilter := &TaskFilter{Tasks: a.TaskList.Tasks, Filter: filter}
	tasks := taskFilter.ApplyFilter()

	for _, task := range tasks {
		task.delete()

		a.TaskList.set(task)
		fmt.Printf("Deleted task %d.\n", task.ID)
	}

	a.save()

	return nil
}

// ShowTask is showing details of a task
func (a *App) ShowTask(filter *Filter) error {
	a.load(false)

	task := a.TaskList.get(filter.IDs[0])
	a.Printer.PrintTask(task)

	return nil
}

// ListTasks is listing tasks
func (a *App) ListTasks(filter *Filter) error {
	a.garbageCollect()
	a.clear()
	a.load(filter.All)

	taskFilter := &TaskFilter{Tasks: a.TaskList.Tasks, Filter: filter}
	tasks := taskFilter.ApplyFilter()

	a.Printer.PrintTaskList(tasks)

	return nil
}

func (a *App) load(all bool) {
	tasks := a.Store.load(all)

	a.TaskList.load(tasks)
}

func (a *App) save() {
	a.Store.save(a.TaskList.Tasks)
}

func (a *App) clear() {
	a.TaskList.Tasks = nil
}

func (a *App) garbageCollect() {
	a.load(false)
	completed := a.TaskList.garbageCollect()
	for _, task := range completed {
		fmt.Printf("%+v\n", task)
		a.Store.complete(task)
	}
	a.save()
}
