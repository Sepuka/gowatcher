package watchers

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/parsers"
	"log"
	"time"
)

const (
	whoCommand      = "who"
	whoLoopInterval = time.Second * 2
)

var (
	users = []string{}
)

func Who(c chan<- command.Result) {
	cmd := command.NewCmd(whoCommand, []string{"-u"}) // with PID
	result := command.Run(cmd)
	notifyAboutNewUsers(result, c)

	for {
		select {
		case <-time.After(whoLoopInterval):
			result := command.Run(cmd)
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError().Error())
				break
			}

			notifyAboutNewUsers(result, c)
		}
	}
}

func notifyAboutNewUsers(raw command.Result, c chan<- command.Result) {
	visitors := parsers.SplitPerLines(raw.GetText())
	for _, userInfo := range visitors {
		if !isUserRegistered(userInfo) {
			userInfoText := fmt.Sprintln("New user detected:", userInfo)
			result := command.NewResult(raw.GetName(), userInfoText, nil)
			c <- result
		}
	}

	users = parsers.SplitPerLines(raw.GetText())
}

func isUserRegistered(userInfo string) bool {
	for _, value := range users {
		if value == userInfo {
			return true
		}
	}

	return false
}
