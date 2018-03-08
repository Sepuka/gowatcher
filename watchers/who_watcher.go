package watchers

import (
	"time"
	"log"
	"github.com/sepuka/gowatcher/parsers"
)

const (
	whoCommand         = "who"
	whoLoopInterval    = time.Minute * 1
	silentCounterLimit = 59
)

var (
	lastUsers     = 0
	silentCounter = 0
)

func Who(config Configuration) {
	result := RunCommand(whoCommand)
	lastUsers = parsers.GetLines(result.GetText())
	SendMessage(result, config)

	for {
		select {
		case <-time.After(whoLoopInterval):
			result := RunCommand(whoCommand)
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
				break
			}

			sendMessage(result, config)
		}
	}
}

func sendMessage(result WatcherResult, config Configuration) {
	users := parsers.GetLines(result.GetText())
	if users > lastUsers {
		SendUrgentMessage(result, config)
		silentCounter = 0
	} else {
		if silentCounter > silentCounterLimit {
			silentCounter = 0
			SendMessage(result, config)
		}
	}
	lastUsers = users
	silentCounter++
}