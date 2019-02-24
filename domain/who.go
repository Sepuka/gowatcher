package domain

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

// show who is logged on
type WhoWatcher struct {
	Command *command.Cmd
	Loop    time.Duration
	Handler command.ResultHandler
}

func (obj *WhoWatcher) Exec() {
	command.RunCmdLoop(obj.Command, obj.Loop, obj.Handler)
}
