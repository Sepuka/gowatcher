package parsers

import (
	"testing"
	"reflect"
)

type perLineDataSet struct {
	text string
	expected []string
}

var perLinesDataProvider = []perLineDataSet{
	{"qwerty", []string{"qwerty"}},
	{"qwerty\nqwerty", []string{"qwerty", "qwerty"}},
}

func TestGetPerLines(t *testing.T) {
	for _, set := range perLinesDataProvider {
		if !reflect.DeepEqual(set.expected, SplitPerLines(set.text)) {
			t.Error("Case ", set.text, " failed")
		}
	}
}
