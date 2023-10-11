package taskgopher

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/helmecke/taskgopher/internal/parser"
	"github.com/helmecke/taskgopher/internal/printer"
	"github.com/helmecke/taskgopher/internal/storage"
	"github.com/helmecke/taskgopher/internal/task"
)

// App is the structure of the taskgopher app
type App struct {
	Store    storage.Storage
	Printer  printer.Printer
	TaskList task.List
}

// NewApp is creating the taskgopher app
func NewApp(location string) *App {
	return &App{
		Store:   storage.NewFileStorage(location),
		Printer: printer.NewText(),
	}
}

// AddTask is creating a task
func (a *App) AddTask(mod *parser.Modification) error {
	a.load(false)

	task := task.NewTask(mod)

	fmt.Printf("Created task %d.\n", a.TaskList.Add(task))
	a.save()

	return nil
}

// ModifyTask is modifying a task
func (a *App) ModifyTask(filter *parser.Filter, mod *parser.Modification) error {
	a.load(filter.All)
	a.TaskList.Filter(filter)

	for _, task := range a.TaskList.Filtered() {
		task.Modify(mod, true)

		a.TaskList.Set(task)
		fmt.Printf("Modified task %d.\n", task.ID)
	}
	a.save()

	return nil
}

// CompleteTask is completing a task
func (a *App) CompleteTask(filter *parser.Filter) error {
	a.load(filter.All)
	a.TaskList.Filter(filter)

	for _, task := range a.TaskList.Filtered() {
		task.Complete()

		a.TaskList.Set(task)
		fmt.Printf("Completed task %d.\n", task.ID)
	}
	a.save()

	return nil
}

// DeleteTask is deleting a task
func (a *App) DeleteTask(filter *parser.Filter) error {
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
	a.TaskList.Filter(filter)

	for _, task := range a.TaskList.Filtered() {
		task.Delete()

		a.TaskList.Set(task)
		fmt.Printf("Deleted task %d.\n", task.ID)
	}

	a.save()

	return nil
}

// ShowTask is showing details of a task
func (a *App) ShowTask(filter *parser.Filter) error {
	a.load(false)
	if len(filter.IDs) > 0 {
		a.Printer.PrintItem(a.TaskList.GetByID(filter.IDs[0]))
	}
	if len(filter.UUIDs) > 0 {
		a.Printer.PrintItem(a.TaskList.GetByUUID(filter.UUIDs[0]))
	}

	return nil
}

// ListTasks is listing tasks
func (a *App) ListTasks(filter *parser.Filter) error {
	a.garbageCollect()
	a.clear()
	a.load(filter.All)
	a.TaskList.Filter(filter)
	a.Printer.PrintList(a.TaskList.Filtered())

	return nil
}

func (a *App) load(all bool) {
	tasks := a.Store.Load(all)

	a.TaskList.Load(tasks)
}

func (a *App) save() {
	a.Store.Save(a.TaskList.Tasks)
}

func (a *App) clear() {
	a.TaskList.Tasks = nil
}

func (a *App) garbageCollect() {
	a.load(false)
	completed := a.TaskList.GarbageCollect()
	for _, task := range completed {
		a.Store.Complete(task)
	}
	a.save()
}
