package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"time"
)

const (
	whoCommand      = "who"
	WhoLoopInterval = time.Second * 2
)

func Who(c chan<- command.Result, config config.WatcherConfig) {
	cmd := command.NewCmd(whoCommand, []string{"-u"}) // with PID
	resultHandler := command.NewLinesChangedResultHandler(c)
	command.RunCmdLoop(cmd, config.GetLoop(), resultHandler)
}
