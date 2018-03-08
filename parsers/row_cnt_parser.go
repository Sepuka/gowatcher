package parsers

import "strings"

func GetLines(data string) int {
	return len(strings.Split(data, "\n"))
}