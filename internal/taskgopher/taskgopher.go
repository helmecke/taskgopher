package taskgopher

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// Taskgopher hold app information
type Taskgopher struct {
	TaskStore store
	TaskList  TaskList
	Options
}

// Options bla
type Options struct {
	Due string
}

// NewTaskgopher creates a new Taskgopher object
func NewTaskgopher(store string) *Taskgopher {
	return &Taskgopher{
		TaskStore: newFileStore(store),
	}
}

// Add task to task list
func (t *Taskgopher) Add(args []string) error {
	t.load(false)

	parser := &Parser{}
	filter, err := parser.ParseArgs(args)
	if err != nil {
		return err
	}

	task := NewTask(filter)

	fmt.Printf("Created task %d.\n", t.TaskList.add(task))
	t.save()

	return nil
}

// Modify task to task list
func (t *Taskgopher) Modify(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	t.load(false)

	parser := &Parser{}
	filter, err := parser.ParseArgs(args[1:])
	if err != nil {
		return err
	}
	task := t.TaskList.get(id)
	EditTask(task, filter)

	t.TaskList.set(task)
	fmt.Printf("Modified task %d.\n", task.ID)
	t.save()

	return nil
}

// Complete task to task list
func (t *Taskgopher) Complete(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	t.load(false)

	task := t.TaskList.get(id)
	task.complete()

	t.TaskList.set(task)
	fmt.Printf("Completed task %d.\n", task.ID)
	t.save()

	return nil
}

// Delete task from task list
func (t *Taskgopher) Delete(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	t.load(false)

	task := t.TaskList.get(id)
	task.delete()

	t.TaskList.set(task)
	fmt.Printf("Deleted task %d.\n", task.ID)
	t.save()

	return nil
}

// Start task from task list
func (t *Taskgopher) Start(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	t.load(false)

	task := t.TaskList.get(id)
	task.start()

	t.TaskList.set(task)
	fmt.Printf("Started task %d.\n", task.ID)
	t.save()

	return nil
}

// Stop task from task list
func (t *Taskgopher) Stop(args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	t.load(false)

	task := t.TaskList.get(id)
	task.stop()

	t.TaskList.set(task)
	t.save()

	fmt.Printf("Stopped task %d.\n", task.ID)

	return nil
}

// List all takss
func (t *Taskgopher) List(all bool) error {
	t.garbageCollect()
	t.clear()
	t.load(all)

	if len(t.TaskList.Tasks) > 0 {
		fmt.Println("")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorder(false)
		table.SetHeader([]string{"ID", "Age", "Title", "Urgency"})

		for _, task := range t.TaskList.ByUrgency() {
			table.Append(task.table())
		}
		table.Render()
		fmt.Printf("\n%d tasks\n", len(t.TaskList.Tasks))
	} else {
		fmt.Println("No tasks found.")
	}

	return nil
}

func (t *Taskgopher) load(all bool) {
	tasks := t.TaskStore.load(all)

	t.TaskList.load(tasks)
}

func (t *Taskgopher) save() {
	t.TaskStore.save(t.TaskList.Tasks)
}

func (t *Taskgopher) clear() {
	t.TaskList.Tasks = nil
}
func (t *Taskgopher) garbageCollect() {
	t.load(false)
	completed := t.TaskList.garbageCollect()
	for _, task := range completed {

		fmt.Printf("%+v\n", task)
		t.TaskStore.complete(task)
	}
	t.save()
}
