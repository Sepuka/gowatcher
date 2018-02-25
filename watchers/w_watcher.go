package watchers

import (
	"time"
	"log"
)

const (
	wCommand      = "w"
	wLoopInterval = time.Hour * 6
)

func W(config Configuration) {
	result := RunCommand(wCommand)
	SendMessage(result, config)

	for {
		select {
			case <-time.After(wLoopInterval):
				result := RunCommand(wCommand)
				if result.IsFailure() {
					log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
					break
				}
				SendMessage(result, config)
		}
	}
}
