package domain

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

// Show who is logged on and what they are doing.
type WWatcher struct {
	Command       *command.Cmd
	Loop          time.Duration
	TransportChan chan<- command.Result
}

func (obj *WWatcher) Exec() {
	handler := command.NewDfFormatResultHandler(obj.TransportChan)
	command.RunCmdLoop(obj.Command, obj.Loop, handler)
}
