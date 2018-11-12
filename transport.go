package main

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/services/logger"
	"github.com/sepuka/gowatcher/transports"
	"log"
)

func Transmitter(c <-chan command.Result) {
	log.Println(services.Container)
	for {
		msg := <-c

		var log = services.Container.Get(services.LoggerComponent).(*logger.Logger)
		log.Debug("Sending '%s':'%s' message.", msg.GetName(), msg.GetType())
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
