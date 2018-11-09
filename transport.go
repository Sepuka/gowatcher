package main

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/transports"
	"log"
)

func Transmitter(c <-chan command.Result) {
	for {
		msg := <-c

		if config.Log.Level == config.LogLevelDebug {
			log.Printf("Sending '%s':'%s' message.", msg.GetName(), msg.GetType())
		}
		go sendToTelegram(msg, config.TelegramConfig)
		go sendToSlack(msg, config.SlackConfig)
	}
}

func sendToTelegram(msg command.Result, config config.TransportTelegram) {
	transports.SendTelegramMessage(msg, config)
}

func sendToSlack(msg command.Result, config config.TransportSlack) {
	transports.SendSlackMessage(msg, config)
}
