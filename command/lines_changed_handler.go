package command

import (
	"fmt"
	"github.com/sepuka/gowatcher/parsers"
)

type LinesChangedResultHandler struct {
	c     chan<-Result
	users []string
}

// Send msg to chan if it detect new lines of watcher's output
func NewLinesChangedResultHandler(transportChan chan<-Result) ResultHandler {
	return &LinesChangedResultHandler{
		transportChan,
		[]string{},
	}
}

func (handler *LinesChangedResultHandler) Handle(raw Result) {
	rows := parsers.SplitPerLines(raw.GetContent())
	for _, userInfo := range rows {
		if !handler.isUserRegistered(userInfo) {
			userInfoText := fmt.Sprintln("New user detected:", userInfo)
			result := NewResult(raw.GetName(), userInfoText, nil)
			handler.c <- result
		}
	}

	handler.users = parsers.SplitPerLines(raw.GetContent())
}

func (handler LinesChangedResultHandler) isUserRegistered(userInfo string) bool {
	for _, value := range handler.users {
		if value == userInfo {
			return true
		}
	}

	return false
}
