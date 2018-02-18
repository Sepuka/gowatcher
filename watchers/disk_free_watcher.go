package watchers

import (
	"strings"
	"fmt"
	"regexp"
)

const diskFreeCommand = "df"
const outputFormat = "^(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)$"

func Work(channel chan<- WatcherResult) {
	result, err := Run(diskFreeCommand, "-hl", "--type=ext4", "--type=ext2", "--type=vfat")
	if err != nil {
		channel <- WatcherResult{
			"",
			err,
			"",
		}
	}

	channel <- WatcherResult{
		parse(result),
		nil,
		result,
	}
}

func parse(data string) string {
	var result []string
	rows := strings.Split(data, "\n")
	for _, row := range rows[1:] {
		reg := regexp.MustCompile(outputFormat)
		rowDetails := reg.FindStringSubmatch(row)
		if len(rowDetails) == 0 {
			break
		}
		size := rowDetails[2]
		used := rowDetails[3]
		avail := rowDetails[4]
		percent := rowDetails[5]
		mount := rowDetails[6]
		format := fmt.Sprintf("%v have %v (%v) used, %v free, %v total size", mount, used, percent, avail, size)
		result = append(result, format)
	}

	return strings.Join(result, "\n")
}