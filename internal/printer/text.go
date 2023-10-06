package printer

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/helmecke/taskgopher/internal/task"
)

// TextPrinter is the structure of the screen printer
type TextPrinter struct{}

// NewText is creating the screen printer
func NewText() *TextPrinter {
	return &TextPrinter{}
}

// PrintItem is printing detailed information on a task
func (p *TextPrinter) PrintItem(task *task.Item) {
	fmt.Println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.Style().Options.DrawBorder = false
	t.AppendHeader(table.Row{"Name", "Value"})
	t.AppendRow(table.Row{"ID", task.ID})
	t.AppendRow(table.Row{"UUID", task.UUID})
	t.AppendRow(table.Row{"Description", task.Description})
	t.AppendRow(table.Row{"Status", task.Status})
	t.AppendRow(table.Row{"Created", task.Created.Format("2006-01-02 15:04:05") + " (" + task.Age() + ")"})

	if !task.Modified.IsZero() {
		t.AppendRow(table.Row{"Modified", task.Modified.Format("2006-01-02 15:04:05") + " (" + task.LastModifiedDiff() + ")"})
	}

	if !task.Completed.IsZero() {
		t.AppendRow(table.Row{"Completed", task.Completed})
	}

	if !task.Due.IsZero() {
		t.AppendRow(table.Row{"Due", task.Due.Format("2006-01-02 15:04:05") + " (" + task.DueDiff() + ")"})
	}

	if len(task.Tags) > 0 {
		t.AppendRow(table.Row{"Tags", task.Tags})
	}

	if task.Project != "" {
		t.AppendRow(table.Row{"Project", task.Project})
	}

	t.AppendRow(table.Row{"Urgency", task.Urgency})

	if len(task.Tags) > 0 {
		t.AppendRow(table.Row{"VirtualTags", strings.Join(task.VirtualTags, " ")})
	}

	t.Render()
	fmt.Println("")
}

// PrintList is printing general information on all tasks
func (p *TextPrinter) PrintList(tasks []*task.Item) {
	if len(tasks) > 0 {
		fmt.Println("")

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleLight)
		t.Style().Options.DrawBorder = false
		t.SortBy([]table.SortBy{
			{Name: "Urgency", Mode: table.DscNumeric},
		})
		t.AppendHeader(table.Row{"ID", "Age", "Title", "Urgency"})
		for _, task := range tasks {
			t.AppendRow(table.Row{task.ID, task.Age(), task.Description, task.Urgency})
		}
		t.Render()

		fmt.Printf("\n%d tasks\n", len(tasks))
	} else {
		fmt.Println("No tasks found.")
	}
}
