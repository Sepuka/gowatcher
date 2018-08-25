package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"log"
	"time"
)

const (
	wCommand      = "w"
	wLoopInterval = time.Hour * 6
)

func W(c chan<- command.Result) {
	cmd := command.NewCmd(wCommand, []string{})
	result := command.Run(cmd)
	c <- result

	for {
		select {
		case <-time.After(wLoopInterval):
			result := command.Run(cmd)
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError().Error())
				break
			}
			c <- result
		}
	}
}
