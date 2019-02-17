package command

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

const outputFormat = "^(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)\\s+(.*?)$"

type DfFormatResultHandler struct {
	c chan<-Result
}

func NewDfFormatResultHandler(transportChan chan<-Result) ResultHandler {
	return DfFormatResultHandler{transportChan}
}

func (handler DfFormatResultHandler) Handle(raw Result) {
	handler.c <- NewResult(raw.GetName(), parse(raw.GetContent()), raw.GetError())
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
