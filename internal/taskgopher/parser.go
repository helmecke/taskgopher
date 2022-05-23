package taskgopher

import (
	"log"
	"strings"
	"time"
)

const rfc3339FullDate = "2006-01-02"

// A Parser parses
type Parser struct{}

// ParseArgs parses args
func (p *Parser) ParseArgs(args []string) (*Filter, error) {
	filter := &Filter{
		HasDue: false,
	}

	var descriptionMatches []string

	for _, arg := range args {
		match := false

		if strings.HasPrefix(arg, "@") {
			match = true
			filter.HasContexts = true
			filter.Contexts = append(filter.Contexts, arg[1:])
		}

		if strings.HasPrefix(arg, "#") {
			match = true
			filter.HasTags = true
			filter.Tags = append(filter.Tags, arg[1:])
		}

		if strings.HasPrefix(arg, "due:") {
			match = true
			filter.HasDue = true
			date, err := time.Parse(rfc3339FullDate, arg[4:])
			if err != nil {
				log.Fatal(err)
			}
			filter.Due = date
		}

		if !match {
			descriptionMatches = append(descriptionMatches, arg)
		}

		filter.Description = strings.Join(descriptionMatches, " ")
	}

	return filter, nil
}
