package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"time"
)

const (
	whoCommand = "who"
)

// show who is logged on
type whoWatcher struct {
	command *command.Cmd
	loop    time.Duration
}

var (
	whoConfig = config.GetWatcherConfig(whoAgentName)
	who       = &whoWatcher{
		command.NewEnvedCmd(whoCommand, utConfig.Args, "lang=en_EN"),
		whoConfig.GetLoop(),
	}
)

func (obj whoWatcher) exec() {
	handler := command.NewLinesChangedResultHandler()
	command.RunCmdLoop(obj.command, obj.loop, handler)
}
