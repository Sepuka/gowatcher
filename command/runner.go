package command

import (
	"time"
	"log"
)

func Runner(cmd Cmd, period time.Duration, c chan<- Result) {
	for {
		select {
		case <-time.After(period):
			result := Run(cmd)
			if result.IsFailure() {
				// TODO send msg about err to channel
				log.Printf("Watcher %v failed: %v.", result.GetName(), result.GetError().Error())
				break
			}
			c <- result
		}
	}
}
