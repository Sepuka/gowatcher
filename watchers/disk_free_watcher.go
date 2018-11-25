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
	command command.Command
	loop time.Duration
}

var (
	dfConfig = config.GetWatcherConfig(diskFreeAgentName)
	df = &dfWatcher{
		&command.Cmd{
			Cmd: diskFreeCommand,
			Args: dfConfig.Args,
		},
		dfConfig.GetLoop(),
	}
)

func (obj dfWatcher) exec() {
	handler := command.NewDfFormatResultHandler()
	command.RunCmdLoop(obj.command, obj.loop, handler)
}
