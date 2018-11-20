package main

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/transports"
	"github.com/sirupsen/logrus"
)

func transmitter(c <-chan command.Result) {
	for {
		msg := <-c

		for _, distributor := range services.Container.Get(services.Transports).([]transports.Transport) {

			log.WithFields(logrus.Fields{
				"transport": distributor.GetName(),
				"msg_type": msg.GetType(),
			}).Debugf("Sending '%s' message.", msg.GetName())

			go distributor.Send(msg)
		}
	}
}
