package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"time"
)

const (
	wCommand = "w"
)

// Show who is logged on and what they are doing.
type wWatcher struct {
	command *command.Cmd
	loop    time.Duration
}

var (
	wConfig = config.GetWatcherConfig(wAgentName)
	w       = &wWatcher{
		command.NewCmd(wCommand, utConfig.Args),
		wConfig.GetLoop(),
	}
)

func (obj wWatcher) exec() {
	handler := command.NewDfFormatResultHandler()
	command.RunCmdLoop(obj.command, obj.loop, handler)
}
