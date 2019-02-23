package domain

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

// report about core temperature
type TemperatureWatcher struct {
	Command *command.Cmd
	Loop    time.Duration
	Handler command.ResultHandler
}

func (obj *TemperatureWatcher) Exec() {
	command.RunCmdLoop(obj.Command, obj.Loop, obj.Handler)
}
