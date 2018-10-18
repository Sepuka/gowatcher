package command

import (
	"fmt"
	"io/ioutil"
	"log"
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

func RunCmdLoop(cmd Command, period time.Duration, resultHandler ResultHandler) {
	callback := func() Result {
		return Run(cmd)
	}

	doPeriodicalTask(period, resultHandler, callback)
}

func doPeriodicalTask(period time.Duration, resultHandler ResultHandler, f func() Result) {
	for {
		select {
		case <-time.After(period):
			result := f()
			if result.IsFailure() {
				// TODO send msg about err to channel
				log.Printf("Watcher %v failed: %v.", result.GetName(), result.GetError().Error())
				break
			}
			resultHandler.Handle(result)
		}
	}
}
