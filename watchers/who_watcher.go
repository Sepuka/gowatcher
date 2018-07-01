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

func Who(c chan<- WatcherResult) {
	result := RunCommand(whoCommand)
	lastUsers = parsers.GetLines(result.GetText())
	sendMessage(result, c)

	for {
		select {
		case <-time.After(whoLoopInterval):
			result := RunCommand(whoCommand)
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
				break
			}

			sendMessage(result, c)
		}
	}
}

func sendMessage(result WatcherResult, c chan<- WatcherResult) {
	users := parsers.GetLines(result.GetText())
	if users > lastUsers {
		c <- result
		silentCounter = 0
	} else {
		if silentCounter > silentCounterLimit {
			silentCounter = 0
			c <- result
		}
	}
	lastUsers = users
	silentCounter++
}