package domain

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

// Show who is logged on and what they are doing.
type WWatcher struct {
	Command       *command.Cmd
	Loop          time.Duration
	Handler command.ResultHandler
}

func (obj *WWatcher) Exec() {
	command.RunCmdLoop(obj.Command, obj.Loop, obj.Handler)
}
