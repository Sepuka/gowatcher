package watchers

import (
	"time"
	"log"
)

const (
	uptimeCommand = "uptime"
	uptimeLoopInterval = time.Hour * 24
)

func Uptime(c chan<- WatcherResult) {
	result := RunCommand(uptimeCommand)
	c <- result

	for {
		select {
		case <-time.After(uptimeLoopInterval):
			result := RunCommand(uptimeCommand)
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError())
				break
			}
			c <- result
		}
	}
}
