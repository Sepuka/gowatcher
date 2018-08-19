package parsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type perLineDataSet struct {
	text     string
	expected []string
}

var perLinesDataProvider = []perLineDataSet{
	{"qwerty", []string{"qwerty"}},
	{"qwerty\nqwerty", []string{"qwerty", "qwerty"}},
}

func TestGetPerLines(t *testing.T) {
	for _, set := range perLinesDataProvider {
		assert.Equal(t, set.expected, SplitPerLines(set.text))
	}
}
