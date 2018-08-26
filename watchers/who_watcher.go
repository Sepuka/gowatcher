package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

const (
	whoCommand      = "who"
	whoLoopInterval = time.Second * 2
)

func Who(c chan<- command.Result) {
	cmd := command.NewCmd(whoCommand, []string{"-u"}) // with PID
	resultHandler := command.NewLinesChangedResultHandler(c)
	command.Runner(cmd, whoLoopInterval, resultHandler)
}
