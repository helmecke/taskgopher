package taskgopher_test

import (
	"strings"
	"testing"
	"time"

	"github.com/helmecke/taskgopher/internal/taskgopher"
	"github.com/stretchr/testify/assert"
)

func TestDue(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	parser := &taskgopher.Parser{}
	_, filter, _ := parser.ParseArgs(strings.Split("here is the subject due:2022-06-01", " "))

	assert.Equal(time.Date(2022, time.June, 1, 0, 0, 0, 0, time.UTC), filter.Due)
}

func TestTag(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	parser := &taskgopher.Parser{}
	_, filter, _ := parser.ParseArgs(strings.Split("here is the subject #bla", " "))

	assert.Equal([]string{"bla"}, filter.Tags)
}

func TestTags(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	parser := &taskgopher.Parser{}
	_, filter, _ := parser.ParseArgs(strings.Split("#test here is the subject #bla", " "))

	assert.Equal([]string{"test", "bla"}, filter.Tags)
}
