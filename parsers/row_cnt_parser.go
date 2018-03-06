package parsers

import "strings"

func Cnt(data interface{}) int {
	switch data.(type) {
		case string:
			return calcStringRows(data.(string))
		default:
			return 0
	}
}

func calcStringRows(data string) int {
	return len(strings.Split(data, "\n"))
}