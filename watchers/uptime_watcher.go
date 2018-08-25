package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"log"
	"time"
)

const (
	uptimeCommand      = "uptime"
	uptimeLoopInterval = time.Hour * 24
)

func Uptime(c chan<- command.Result) {
	cmd := command.NewCmd(uptimeCommand, []string{})
	result := command.Run(cmd)
	c <- result

	for {
		select {
		case <-time.After(uptimeLoopInterval):
			result := command.Run(cmd)
			if result.IsFailure() {
				log.Printf("Watcher %v failed: %v", result.GetName(), result.GetError().Error())
				break
			}
			c <- result
		}
	}
}
