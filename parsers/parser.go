package parsers

import (
	"strings"
	"unicode"
	"strconv"
	"log"
)

func SplitPerLines(data string) []string {
	return strings.Split(data, "\n")
}

func SplitPerCols(data string, delimiter string) []string {
	return strings.Split(data, delimiter)
}

func FetchInt(value string) uint64 {
	res, err := strconv.ParseUint(trimSpaces(value), 10, 0)

	if err != nil {
		log.Println("Cannot read value ", value, " as int, error: ", err)
		return 0
	}

	return res
}

func FetchFloat(value string) float64 {
	res, err := strconv.ParseFloat(trimSpaces(value), 64)

	if err != nil {
		log.Println("Cannot read value ", value, " as float, error: ", err)
		return 0
	}

	return res
}

func trimSpaces(data string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, data)
}