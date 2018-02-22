package watchers

import (
	"time"
	"log"
)

const whoCommand = "who"

func Who(config Configuration) {
	result := RunCommand(whoCommand)
	SendMessage(result, config)

	for {
		select {
		case <-time.After(time.Second * config.MainLoopInterval):
			result := RunCommand(whoCommand)
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
				break
			}
			SendMessage(result, config)
		}
	}
}
