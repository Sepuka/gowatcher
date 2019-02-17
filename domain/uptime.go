package domain

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

// Tell how long the system has been running.
type UptimeWatcher struct {
	Command       *command.Cmd
	Loop          time.Duration
	TransportChan chan<- command.Result
}

func (obj UptimeWatcher) Exec() {
	handler := command.NewDummyResultHandler(obj.TransportChan)
	command.RunCmdLoop(obj.Command, obj.Loop, handler)
}
