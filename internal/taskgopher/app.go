package taskgopher

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type App struct {
	TaskStore store
	TaskList  TaskList
}

// NewApp creates a new Taskgopher app.
func NewApp(store string) *App {
	return &App{
		TaskStore: newFileStore(store),
	}
}

// Add task to task list
func (a *App) Add(args []string) error {
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

// Modify task to task list
func (a *App) Modify(args []string) error {
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

// Complete task to task list
func (a *App) Complete(args []string) error {
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

// Delete task from task list
func (a *App) Delete(args []string) error {
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

// Start task from task list
func (a *App) Start(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	a.load(false)

	task := a.TaskList.get(id)
	task.start()

	a.TaskList.set(task)
	fmt.Printf("Started task %d.\n", task.ID)
	a.save()

	return nil
}

// Stop task from task list
func (a *App) Stop(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	a.load(false)

	task := a.TaskList.get(id)
	task.stop()

	a.TaskList.set(task)
	a.save()

	fmt.Printf("Stopped task %d.\n", task.ID)

	return nil
}

// List all takss
func (a *App) List(args []string, all bool) error {
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

	if len(tasks) > 0 {
		fmt.Println("")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorder(false)
		table.SetHeader([]string{"ID", "Age", "Title", "Urgency"})

		for _, task := range tasks {
			table.Append(task.table())
		}
		table.Render()
		fmt.Printf("\n%d tasks\n", len(tasks))
	} else {
		fmt.Println("No tasks found.")
	}

	return nil
}

func (a *App) load(all bool) {
	tasks := a.TaskStore.load(all)

	a.TaskList.load(tasks)
}

func (a *App) save() {
	a.TaskStore.save(a.TaskList.Tasks)
}

func (a *App) clear() {
	a.TaskList.Tasks = nil
}

func (a *App) garbageCollect() {
	a.load(false)
	completed := a.TaskList.garbageCollect()
	for _, task := range completed {

		fmt.Printf("%+v\n", task)
		a.TaskStore.complete(task)
	}
	a.save()
}
