package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"time"
)

const (
	wCommand      = "w"
	wLoopInterval = time.Hour * 6
)

func W(c chan<- command.Result) {
	cmd := command.NewCmd(wCommand, []string{})
	resultHandler := command.NewDummyResultHandler(c)
	command.Runner(cmd, wLoopInterval, resultHandler)
}
