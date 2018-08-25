package watchers

import (
	"bytes"
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"log"
	"regexp"
	"strings"
	"time"
)

const (
	diskFreeCommand = "df"
	outputFormat    = "^(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)$"
	dfLoopInterval  = time.Hour * 6
)

func DiskFree(c chan<- command.Result) {
	cmd := command.NewCmd(diskFreeCommand, []string{"-hl", "--type=ext4", "--type=ext2", "--type=vfat"})
	raw := command.Run(cmd)
	result := command.NewResult(raw.GetName(), parse(raw.GetText()), raw.GetError())
	c <- result

	for {
		select {
		case <-time.After(dfLoopInterval):
			raw := command.Run(cmd)
			if raw.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError().Error())
				break
			}
			result := command.NewResult(raw.GetName(), parse(raw.GetText()), raw.GetError())
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
