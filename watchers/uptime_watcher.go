package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

const (
	uptimeCommand      = "uptime"
	uptimeLoopInterval = time.Hour * 24
)

func Uptime(c chan<- command.Result) {
	cmd := command.NewCmd(uptimeCommand, []string{})
	command.Runner(cmd, uptimeLoopInterval, c)
}
