package main

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/transports"
)

func transmitter(c <-chan command.Result) {
	for {
		msg := <-c

		log.Debugf("Sending '%s':'%s' message.", msg.GetName(), msg.GetType())
		for _, distributor := range services.Container.Get(services.Transports).([]transports.Transport) {
			go distributor.Send(msg)
		}
	}
}
