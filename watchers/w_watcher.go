package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"time"
	"github.com/sepuka/gowatcher/config"
)

const (
	wCommand      = "w"
	WLoopInterval = time.Hour * 6
)

func W(c chan<- command.Result, config config.WatcherConfig) {
	cmd := command.NewCmd(wCommand, []string{})
	resultHandler := command.NewDummyResultHandler(c)
	command.Runner(cmd, config.GetLoop(), resultHandler)
}
