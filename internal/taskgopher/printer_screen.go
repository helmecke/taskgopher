package taskgopher

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type ScreenPrinter struct{}

func newScreenPrinter() *ScreenPrinter {
	return &ScreenPrinter{}
}

func (s *ScreenPrinter) PrintTaskList(tasks []*Task) {
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
}
