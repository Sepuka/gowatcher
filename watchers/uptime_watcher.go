package watchers

import (
	"time"
	"log"
)

const uptimeCommand = "uptime"

func Uptime(config Configuration) {
	result := RunCommand(uptimeCommand)
	SendMessage(result, config)

	for {
		select {
		case <-time.After(time.Second * config.MainLoopInterval):
			result := RunCommand(uptimeCommand)
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
				break
			}
			SendMessage(result, config)
		}
	}
}
