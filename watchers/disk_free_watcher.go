package watchers

import (
	"strings"
	"fmt"
	"regexp"
	"bytes"
	"time"
	"log"
)

const (
	diskFreeCommand = "df"
 	outputFormat = "^(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)$"
	dfLoopInterval = time.Hour * 6
)

func DiskFree(c chan<- WatcherResult) {
	result := RunCommand(diskFreeCommand, "-hl", "--type=ext4", "--type=ext2", "--type=vfat")
	result.text = parse(result.raw)
	c <- result

	for {
		select {
		case <-time.After(dfLoopInterval):
			result := RunCommand(diskFreeCommand, "-hl", "--type=ext4", "--type=ext2", "--type=vfat")
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
				break
			}
			result.text = parse(result.raw)
			c <- result
		}
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
		format := fmt.Sprintf("%v has %v (%v) used, %v free, %v total size\n", mount, used, percent, avail, size)
		buffer.WriteString(format)
	}

	return buffer.String()
}