package parser

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/helmecke/taskgopher/pkg/sliceutils"
)

const rfc3339FullDate = "2006-01-02"

// A Parser parses user input
type Parser struct {
	Command      string
	Filter       *Filter
	Modification *Modification
}

// NewParser creates a new argument parser ready to parse inputs
func NewParser() *Parser {
	return &Parser{
		Command:      "",
		Filter:       &Filter{},
		Modification: &Modification{},
	}
}

// ParseArgs parses args
// nolint:gocognit
func (p *Parser) ParseArgs(args []string) (err error) {
	cmdAtIndex := -1
	var description []string

	for i, arg := range args {
		if p.Command == "" && isFilterCommand(arg) {
			p.Command = arg
			cmdAtIndex = i

			// break here to only parse args before filterCommands, to enforce
			// taskgopher <filter> list
			break
		}

		if u, err := uuid.Parse(arg); err == nil {
			p.Filter.UUIDs = append(p.Filter.UUIDs, u)
			p.Filter.Found = true

			continue
		}

		if s, err := strconv.ParseInt(arg, 10, 64); err == nil {
			p.Filter.IDs = append(p.Filter.IDs, int(s))
			p.Filter.Found = true

			continue
		}

		if arg == "all" {
			p.Filter.All = true

			continue
		}

		if strings.HasPrefix(arg, "due:") {
			date, err := time.Parse(rfc3339FullDate, arg[4:])
			if err != nil {
				log.Fatal(err)
			}
			p.Filter.Due = date
			p.Filter.Found = true

			continue
		}

		if strings.HasPrefix(arg, "project:") {
			p.Filter.Project = arg[8:]
			p.Filter.Found = true

			continue
		}
	}

	if isModifyCommand(p.Command) {
		for _, arg := range args[cmdAtIndex+1:] {
			if strings.HasPrefix(arg, "due:") {
				if arg[4:] == "-" {
					p.Modification.RemoveDue = true

					continue
				}

				date, err := time.Parse(rfc3339FullDate, arg[4:])
				if err != nil {
					log.Fatal(err)
				}
				p.Modification.Due = date

				continue
			}

			if strings.HasPrefix(arg, "project:") {
				if arg[8:] == "-" {
					p.Modification.RemoveProject = true

					continue
				}

				p.Modification.Project = arg[8:]

				continue
			}

			description = append(description, arg)
		}
	}

	if len(description) > 0 {
		p.Modification.Description = strings.Join(description, " ")
	}

	return err
}

func isFilterCommand(cmd string) bool {
	filterCommands := []string{
		"complete",
		"delete",
		"list",
		"modify",
		"mod",
		"show",
	}

	return sliceutils.StrSliceContains(filterCommands, cmd)
}

func isModifyCommand(cmd string) bool {
	modifyCommands := []string{
		"add",
		"complete",
		"modify",
		"mod",
	}

	return sliceutils.StrSliceContains(modifyCommands, cmd)
}
