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
		go sendToTelegram(msg)
		go sendToSlack(msg)
	}
}

func sendToTelegram(msg command.Result) {
	telegram := services.Container.Get(services.Telegram).(*transports.Telegram)
	telegram.Send(msg)
}

func sendToSlack(msg command.Result) {
	slack := services.Container.Get(services.Slack).(*transports.Slack)
	slack.Send(msg)
}
