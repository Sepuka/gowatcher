package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"time"
)

const (
	whoCommand = "who"
)

var (
	whoConfig = config.GetWatcherConfig(whoAgentName)
	who = &whoWatcher{
		&command.Cmd{
			Cmd: whoCommand,
			Args: whoConfig.Args,
			Env: []string{"lang=en_EN"},
		},
		whoConfig.GetLoop(),
	}
)

// show who is logged on
type whoWatcher struct {
	command command.Command
	loop time.Duration
}

func (obj whoWatcher) exec() {
	handler := command.NewLinesChangedResultHandler()
	command.RunCmdLoop(obj.command, obj.loop, handler)
}
