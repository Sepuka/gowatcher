package command

import (
	"log"
	"time"
)
// TODO add
func Runner(cmd Command, period time.Duration, resultHandler ResultHandler) {
	for {
		select {
		case <-time.After(period):
			result := Run(cmd)
			if result.IsFailure() {
				// TODO send msg about err to channel
				log.Printf("Watcher %v failed: %v.", result.GetName(), result.GetError().Error())
				break
			}
			resultHandler.Handle(result)
		}
	}
}
