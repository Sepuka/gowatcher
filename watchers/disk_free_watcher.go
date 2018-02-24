package watchers

import (
	"strings"
	"fmt"
	"regexp"
	"bytes"
)

const diskFreeCommand = "df"
const outputFormat = "^(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)$"

func DiskFree(channel chan<- WatcherResult) {
	result, err := Run(diskFreeCommand, "-hl", "--type=ext4", "--type=ext2", "--type=vfat")
	if err != nil {
		channel <- WatcherResult{
			diskFreeCommand,
			"",
			err,
			"",
		}
	}

	channel <- WatcherResult{
		diskFreeCommand,
		parse(result),
		nil,
		result,
	}
}

func parse(data string) string {
	var buffer bytes.Buffer
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
		format := fmt.Sprintf("%v have %v (%v) used, %v free, %v total size\n", mount, used, percent, avail, size)
		buffer.WriteString(format)
	}

	return buffer.String()
}