package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"time"
)

const (
	wCommand = "w"
)

var (
	wConfig = config.GetWatcherConfig(wAgentName)
	w = &wWatcher{
		&command.Cmd{
			Cmd: wCommand,
			Args: wConfig.Args,
		},
		wConfig.GetLoop(),
	}
)

// Show who is logged on and what they are doing.
type wWatcher struct {
	command command.Command
	loop time.Duration
}

func (obj wWatcher) exec() {
	handler := command.NewDfFormatResultHandler()
	command.RunCmdLoop(obj.command, obj.loop, handler)
}