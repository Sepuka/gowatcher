package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"time"
	"github.com/sepuka/gowatcher/config"
)

const (
	whoCommand      = "who"
	WhoLoopInterval = time.Second * 2
)

func Who(c chan<- command.Result, config config.WatcherConfig) {
	cmd := command.NewCmd(whoCommand, []string{"-u"}) // with PID
	resultHandler := command.NewLinesChangedResultHandler(c)
	command.Runner(cmd, config.GetLoop(), resultHandler)
}
