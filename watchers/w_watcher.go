package watchers

import (
	"time"
	"log"
)

const (
	wCommand = "w"
	loopInterval = time.Hour * 6
)

func W(config Configuration) {
	for {
		select {
			case <-time.After(loopInterval):
				result := RunCommand(wCommand)
				if result.IsFailure() {
					log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
					break
				}
				SendMessage(result, config)
		}
	}
}
