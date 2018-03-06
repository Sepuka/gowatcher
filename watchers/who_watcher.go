package watchers

import (
	"time"
	"log"
	"github.com/sepuka/gowatcher/parsers"
)

const (
	whoCommand = "who"
	whoLoopInterval = time.Minute * 1
	nosendCounterLimit = 59
)

var (
	lastUsers = 0
	nosendCounter = 0
)

func Who(config Configuration) {
	result := RunCommand(whoCommand)
	lastUsers = parsers.Cnt(result.GetText())
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
	if parsers.Cnt(result.GetText()) > lastUsers {
		SendUrgentMessage(result, config)
	} else {
		if nosendCounter > nosendCounterLimit {
			nosendCounter = 0
			SendMessage(result, config)
		}
	}
	nosendCounter++
}