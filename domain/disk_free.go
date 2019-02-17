package domain

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

// report file system disk space usage
type DfWatcher struct {
	Command *command.Cmd
	Loop    time.Duration
	TransportChan chan<- command.Result
}

func (obj DfWatcher) Exec() {
	handler := command.NewDfFormatResultHandler(obj.TransportChan)
	command.RunCmdLoop(obj.Command, obj.Loop, handler)
}
