package watchers

import (
	"time"
	"log"
	"github.com/sepuka/gowatcher/parsers"
	"fmt"
)

const (
	whoCommand         = "who"
	whoLoopInterval    = time.Second * 2
)

var (
	users = []string{}
)

func Who(c chan<- WatcherResult) {
	result := RunCommand(whoCommand, "-u")// with PID
	notifyAboutNewUsers(result, c)

	for {
		select {
		case <-time.After(whoLoopInterval):
			result := RunCommand(whoCommand, "-u")
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
				break
			}

			notifyAboutNewUsers(result, c)
		}
	}
}

func notifyAboutNewUsers(result WatcherResult, c chan<- WatcherResult) {
	visitors := parsers.GetPerLines(result.GetText())
	for _, userInfo := range visitors {
		if !isUserRegistered(userInfo) {
			userInfoText := fmt.Sprintln("New user detected:", userInfo)
			c <- WatcherResult{watcherName: whoCommand, text: userInfoText}
		}
	}

	users = parsers.GetPerLines(result.GetText())
}

func isUserRegistered(userInfo string) bool {
	for _, value := range users {
		if value == userInfo {
			return true
		}
	}

	return false
}