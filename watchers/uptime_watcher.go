package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"time"
	"github.com/sepuka/gowatcher/config"
)

const (
	uptimeCommand      = "uptime"
	UptimeLoopInterval = time.Hour * 24
)

func Uptime(c chan<- command.Result, config config.WatcherConfig) {
	cmd := command.NewCmd(uptimeCommand, []string{})
	resultHandler := command.NewDummyResultHandler(c)
	command.Runner(cmd, config.GetLoop(), resultHandler)
}
