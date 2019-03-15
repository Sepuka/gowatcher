package command

import (
	"fmt"
	"github.com/sepuka/gowatcher/definition/logger"
	"github.com/sepuka/gowatcher/services"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

func ReadFileLoop(fileName string, period time.Duration, resultHandler ResultHandler) {
	readLaFunc := func() Result {
		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			return NewResult(
				"Periodically file reader",
				fmt.Sprint("Cannot read ", fileName, " file"),
				err,
			)
		}

		return NewResult(
			"Periodically file reader",
			string(bytes[:]),
			nil,
		)
	}

	doPeriodicalTask(period, resultHandler, readLaFunc)
}

func RunFuncLoop(f func() Result, period time.Duration, resultHandler ResultHandler) {
	doPeriodicalTask(period, resultHandler, f)
}

func RunCmdLoop(cmd *Cmd, period time.Duration, resultHandler ResultHandler) {
	callback := func() Result {
		return runConsoleCommand(cmd)
	}

	doPeriodicalTask(period, resultHandler, callback)
}

func doPeriodicalTask(period time.Duration, resultHandler ResultHandler, f func() Result) {
	var log *zap.Logger
	for {
		select {
		case <-time.After(period):
			result := f()
			if result.IsFailure() {
				log = services.Container.Get(logger.DefLogger).(*zap.Logger)
				log.With(
					zap.String("watcher", result.GetName()),
					zap.String("content", result.GetContent()),
					zap.Error(result.GetError()),
				).Error("watcher failed")
				break
			}
			resultHandler.Handle(result)
		}
	}
}
