package taskgopher

import (
	"fmt"
	"log"
	"strconv"

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
		Printer: newScreenPrinter(),
	}
}

// AddTask is creating a task
func (a *App) AddTask(args []string) error {
	a.load(false)

	parser := &Parser{}
	filter, err := parser.ParseArgs(args)
	if err != nil {
		return err
	}

	task := NewTask(filter)

	fmt.Printf("Created task %d.\n", a.TaskList.add(task))
	a.save()

	return nil
}

// ModifyTask is modifying a task
func (a *App) ModifyTask(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	a.load(false)

	parser := &Parser{}
	filter, err := parser.ParseArgs(args[1:])
	if err != nil {
		return err
	}
	task := a.TaskList.get(id)
	EditTask(task, filter)

	a.TaskList.set(task)
	fmt.Printf("Modified task %d.\n", task.ID)
	a.save()

	return nil
}

// CompleteTask is completing a task
func (a *App) CompleteTask(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	a.load(false)

	task := a.TaskList.get(id)
	task.complete()

	a.TaskList.set(task)
	fmt.Printf("Completed task %d.\n", task.ID)
	a.save()

	return nil
}

// DeleteTask is deleting a task
func (a *App) DeleteTask(args []string) error {
	prompt := promptui.Prompt{
		Label:     "Delete task",
		IsConfirm: true,
	}

	if _, err := prompt.Run(); err != nil {
		fmt.Println("Aborted...")

		return nil
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	a.load(false)

	task := a.TaskList.get(id)
	task.delete()

	a.TaskList.set(task)
	fmt.Printf("Deleted task %d.\n", task.ID)
	a.save()

	return nil
}

// ShowTask is showing details of a task
func (a *App) ShowTask(args []string) error {
	return nil
}

// ListTasks is listing tasks
func (a *App) ListTasks(args []string, all bool) error {
	a.garbageCollect()
	a.clear()
	a.load(all)
	parser := &Parser{}
	filter, err := parser.ParseArgs(args)
	if err != nil {
		return err
	}
	taskFilter := &TaskFilter{Tasks: a.TaskList.ByUrgency(), Filter: filter}
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
