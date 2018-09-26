package parsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DataSet struct {
	text     string
	expected []string
}

var perLinesDataProvider = []DataSet{
	{"qwerty", []string{"qwerty"}},
	{"qwerty\nqwerty", []string{"qwerty", "qwerty"}},
}

var perColsDataProvider = []DataSet{
	{"qwerty", []string{"qwerty"}},
	{"qwe rty", []string{"qwe", "rty"}},
}

func TestGetPerLines(t *testing.T) {
	for _, set := range perLinesDataProvider {
		assert.Equal(t, set.expected, SplitPerLines(set.text))
	}
}

func TestSplitPerCols(t *testing.T) {
	for _, set := range perColsDataProvider {
		assert.Equal(t, set.expected, SplitPerCols(set.text, " "))
	}
}
