package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"time"
)

const (
	uptimeCommand = "uptime"
)

// Tell how long the system has been running.
type uptimeWatcher struct {
	command *command.Cmd
	loop    time.Duration
}

var (
	utConfig = config.GetWatcherConfig(upTimeAgentName)
	ut       = &uptimeWatcher{
		command.NewCmd(uptimeCommand, utConfig.Args),
		utConfig.GetLoop(),
	}
)

func (obj uptimeWatcher) exec() {
	handler := command.NewDummyResultHandler()
	command.RunCmdLoop(obj.command, obj.loop, handler)
}
