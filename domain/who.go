package domain

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

// show who is logged on
type WhoWatcher struct {
	Command *command.Cmd
	Loop    time.Duration
	TransportChan chan<- command.Result
}

func (obj *WhoWatcher) Exec() {
	handler := command.NewLinesChangedResultHandler(obj.TransportChan)
	command.RunCmdLoop(obj.Command, obj.Loop, handler)
}
