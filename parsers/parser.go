package parsers

import (
	"strings"
	"unicode"
)

func SplitPerLines(data string) []string {
	return strings.Split(data, "\n")
}

func SplitPerCols(data string, delimiter string) []string {
	return strings.Split(data, delimiter)
}

func TrimSpaces(data string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, data)
}
