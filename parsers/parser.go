package parsers

import "strings"

func SplitPerLines(data string) []string {
	return strings.Split(data, "\n")
}