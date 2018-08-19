package watchers

import (
	"log"
	"time"
)

const (
	wCommand      = "w"
	wLoopInterval = time.Hour * 6
)

func W(c chan<- WatcherResult) {
	result := RunCommand(wCommand)
	c <- result

	for {
		select {
		case <-time.After(wLoopInterval):
			result := RunCommand(wCommand)
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
				break
			}
			c <- result
		}
	}
}
