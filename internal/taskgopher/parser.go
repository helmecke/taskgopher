package taskgopher

import (
	"log"
	"strings"
	"time"
)

const RFC3339FullDate = "2006-01-02"

// A Parser parses
type Parser struct{}

// ParseArgs parses args
func (p *Parser) ParseArgs(args []string, task *Task) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "@") {
			task.Contexts = append(task.Contexts, arg[1:])
		}
		if strings.HasPrefix(arg, "#") {
			task.Tags = append(task.Tags, arg[1:])
		}
		if strings.HasPrefix(arg, "due:") {
			date, err := time.Parse(RFC3339FullDate, arg[4:])
			date = date.Local()
			if err != nil {
				log.Fatal(err)
			}
			task.Due = &date
		}
	}
}
