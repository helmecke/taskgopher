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
func (a *App) AddTask(mod *Modification) error {
	a.load(false)

	task := NewTask(mod)

	fmt.Printf("Created task %d.\n", a.TaskList.add(task))
	a.save()

	return nil
}

// ModifyTask is modifying a task
func (a *App) ModifyTask(filter *Filter, mod *Modification) error {
	a.load(filter.All)
	a.TaskList.filter(filter)

	for _, task := range a.TaskList.filtered() {
		task.modify(mod)

		a.TaskList.set(task)
		fmt.Printf("Modified task %d.\n", task.ID)
	}
	a.save()

	return nil
}

// CompleteTask is completing a task
func (a *App) CompleteTask(filter *Filter) error {
	a.load(filter.All)
	a.TaskList.filter(filter)

	for _, task := range a.TaskList.filtered() {
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
	a.TaskList.filter(filter)

	for _, task := range a.TaskList.filtered() {
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
	a.Printer.PrintTask(a.TaskList.get(filter.IDs[0]))

	return nil
}

// ListTasks is listing tasks
func (a *App) ListTasks(filter *Filter) error {
	a.garbageCollect()
	a.clear()
	a.load(filter.All)
	a.TaskList.filter(filter)
	a.Printer.PrintTaskList(a.TaskList.filtered())

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
		a.Store.complete(task)
	}
	a.save()
}
