package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
)

const (
	uptimeCommand = "uptime"
)

func Uptime(c chan<- command.Result, config config.WatcherConfig) {
	cmd := command.NewCmd(uptimeCommand, []string{})
	resultHandler := command.NewDummyResultHandler(c)
	command.RunCmdLoop(cmd, config.GetLoop(), resultHandler)
}
