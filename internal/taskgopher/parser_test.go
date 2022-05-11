package taskgopher

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDescription(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	filter, _ := parser.ParseArgs(strings.Split("@home here is the subject", " "))

	assert.Equal("here is the subject", filter.Description)
}

func TestDue(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	filter, _ := parser.ParseArgs(strings.Split("here is the subject due:2022-06-01", " "))

	testTime := time.Time(time.Date(2022, time.June, 1, 2, 0, 0, 0, time.Local))
	assert.Equal(&testTime, filter.Due)
}

func TestTag(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	filter, _ := parser.ParseArgs(strings.Split("here is the subject #bla", " "))

	assert.Equal([]string{"bla"}, filter.Tags)
}

func TestTags(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	filter, _ := parser.ParseArgs(strings.Split("#test here is the subject #bla", " "))

	assert.Equal([]string{"test", "bla"}, filter.Tags)
}
