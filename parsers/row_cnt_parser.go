package parsers

import "strings"

func GetNumberOfLines(data string) int {
	return len(strings.Split(data, "\n"))
}

func GetPerLines(data string) []string {
	return strings.Split(data, "\n")
}