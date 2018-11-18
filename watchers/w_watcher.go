package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
)

const (
	wCommand = "w"
)

func W(c chan<- command.Result, config config.WatcherConfig) {
	cmd := command.NewCmd(wCommand, []string{})
	resultHandler := command.NewDummyResultHandler(c)
	command.RunCmdLoop(cmd, config.GetLoop(), resultHandler)
}
