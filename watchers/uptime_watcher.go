package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"time"
)

const (
	uptimeCommand      = "uptime"
	UptimeLoopInterval = time.Hour * 24
)

func Uptime(c chan<- command.Result, config config.WatcherConfig) {
	cmd := command.NewCmd(uptimeCommand, []string{})
	resultHandler := command.NewDummyResultHandler(c)
	command.RunCmdLoop(cmd, config.GetLoop(), resultHandler)
}
