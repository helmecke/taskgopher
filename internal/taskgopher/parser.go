package taskgopher

import (
	"log"
	"strconv"
	"strings"
	"time"
)

const rfc3339FullDate = "2006-01-02"

var filterCommands = []string{
	"add",
	"complete",
	"delete",
	"list",
	"modify",
	"show",
}

// A Parser parses
type Parser struct{}

// ParseArgs parses args
func (p *Parser) ParseArgs(args []string) (string, *Filter, error) {
	filter := &Filter{}
	cmd := ""

	for _, arg := range args {
		lowerCased := strings.ToLower(arg)

		if cmd == "" && contains(filterCommands, lowerCased) {
			cmd = lowerCased

			// break here to only parse args before filterCommands, to enforce
			// taskgopher <filter> list
			break
		}

		if s, err := strconv.ParseInt(arg, 10, 64); err == nil {
			filter.IDs = append(filter.IDs, int(s))

			continue
		}

		if arg == "all" {
			filter.All = true

			continue
		}

		if strings.HasPrefix(arg, "@") {
			filter.HasContexts = true
			filter.Contexts = append(filter.Contexts, arg[1:])
		}

		if strings.HasPrefix(arg, "#") {
			filter.HasTags = true
			filter.Tags = append(filter.Tags, arg[1:])
		}

		if strings.HasPrefix(arg, "due:") {
			filter.HasDue = true
			date, err := time.Parse(rfc3339FullDate, arg[4:])
			if err != nil {
				log.Fatal(err)
			}
			filter.Due = date
		}
	}

	return cmd, filter, nil
}

// contains returns true if string is in slice of strings
func contains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}
