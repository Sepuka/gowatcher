package watchers

import (
	"time"
	"log"
)

const (
	whoCommand = "who"
	whoLoopInterval = time.Hour * 3
)

func Who(config Configuration) {
	result := RunCommand(whoCommand)
	SendMessage(result, config)

	for {
		select {
		case <-time.After(whoLoopInterval):
			result := RunCommand(whoCommand)
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
				break
			}
			SendMessage(result, config)
		}
	}
}
