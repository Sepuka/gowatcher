package main

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/transports"
)

func Transmitter(c <-chan command.Result) {
	for {
		msg := <-c

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
