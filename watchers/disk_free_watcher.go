package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"time"
)

const (
	diskFreeCommand = "df"
)

// report file system disk space usage
type dfWatcher struct {
	command *command.Cmd
	loop    time.Duration
}

var (
	dfConfig = config.GetWatcherConfig(diskFreeAgentName)
	df       = &dfWatcher{
		command.NewCmd(diskFreeCommand, dfConfig.Args),
		dfConfig.GetLoop(),
	}
)

func (obj dfWatcher) exec() {
	handler := command.NewDfFormatResultHandler()
	command.RunCmdLoop(obj.command, obj.loop, handler)
}
