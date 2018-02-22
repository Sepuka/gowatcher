package watchers

import (
	"strings"
	"fmt"
	"regexp"
	"time"
	"log"
)

const diskFreeCommand = "df"
const outputFormat = "^(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)$"

func DiskFree(config Configuration) {
	result := RunCommand(diskFreeCommand, "-hl", "--type=ext4", "--type=ext2", "--type=vfat")
	result.text = parse(result.raw)
	SendMessage(result, config)

	for {
		select {
		case <-time.After(time.Second * config.MainLoopInterval):
			result := RunCommand(diskFreeCommand, "-hl", "--type=ext4", "--type=ext2", "--type=vfat")
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
				break
			}
			result.text = parse(result.raw)
			SendMessage(result, config)
		}
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