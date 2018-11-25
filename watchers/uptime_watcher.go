package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"time"
)

const (
	uptimeCommand = "uptime"
)

var (
	utConfig = config.GetWatcherConfig(upTimeAgentName)
	ut = &uptimeWatcher{
		&command.Cmd{
			Cmd: uptimeCommand,
			Args: utConfig.Args,
		},
		utConfig.GetLoop(),
	}
)

// Tell how long the system has been running.
type uptimeWatcher struct {
	command command.Command
	loop time.Duration
}

func (obj uptimeWatcher) exec() {
	handler := command.NewDummyResultHandler()
	command.RunCmdLoop(obj.command, obj.loop, handler)
}