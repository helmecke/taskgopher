package taskgopher

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type ScreenPrinter struct{}

func NewScreenPrinter() *ScreenPrinter {
	return &ScreenPrinter{}
}

func (s *ScreenPrinter) PrintTask(task *Task) {
	fmt.Println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.Style().Options.DrawBorder = false
	t.AppendHeader(table.Row{"Name", "Value"})
	t.AppendRow(table.Row{"ID", task.ID})
	t.AppendRow(table.Row{"UUID", task.UUID})
	t.AppendRow(table.Row{"Status", task.Status})
	t.AppendRow(table.Row{"Created", task.Created.Format("2006-01-02 15:04:05") + " (" + task.age() + ")"})

	if task.Modified != nil {
		t.AppendRow(table.Row{"Modified", task.Modified.Format("2006-01-02 15:04:05") + " (" + task.lastModified() + ")"})
	}

	if task.Completed != nil {
		t.AppendRow(table.Row{"Completed", task.Completed})
	}

	if task.Due != nil {
		t.AppendRow(table.Row{"Due", task.Due.Format("2006-01-02 15:04:05") + " (" + task.due() + ")"})
	}

	if len(task.Tags) > 0 {
		t.AppendRow(table.Row{"Tags", task.Tags})
	}

	if task.Project != "" {
		t.AppendRow(table.Row{"Project", task.Project})
	}

	if len(task.Contexts) > 0 {
		t.AppendRow(table.Row{"Contexts", task.Contexts})
	}

	t.AppendRow(table.Row{"Urgency", task.Urgency})

	t.Render()
	fmt.Println("")
}

func (s *ScreenPrinter) PrintTaskList(tasks []*Task) {
	if len(tasks) > 0 {
		fmt.Println("")

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleLight)
		t.Style().Options.DrawBorder = false
		t.AppendHeader(table.Row{"ID", "Age", "Title", "Urgency"})
		for _, task := range tasks {
			t.AppendRow(table.Row{task.ID, task.age(), task.Description, task.Urgency})
		}
		t.Render()

		fmt.Printf("\n%d tasks\n", len(tasks))
	} else {
		fmt.Println("No tasks found.")
	}
}
