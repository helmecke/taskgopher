package taskgopher

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescription(t *testing.T) {
	assert := assert.New(t)
	parser := &Parser{}
	filter, _ := parser.ParseArgs(strings.Split("@home here is the subject", " "))

	assert.Equal("here is the subject", filter.Description)
}
