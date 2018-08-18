package main

import (
	"github.com/sepuka/gowatcher/transports"
	"github.com/sepuka/gowatcher/watchers"
	"github.com/stevenroose/gonfig"
)

var telegramConfig watchers.TransportTelegram
var slackConfig watchers.TransportSlack

func initConfig() {
	tconf := config.Transports["telegram"].(map[string]interface{})
	gonfig.LoadMap(&telegramConfig, tconf, gonfig.Conf{})

	sconf := config.Transports["slack"].(map[string]interface{})
	gonfig.LoadMap(&slackConfig, sconf, gonfig.Conf{})
}

func Transmitter(c <-chan watchers.WatcherResult) {
	initConfig()

	for {
		msg := <-c

		go sendToTelegram(msg, telegramConfig)
		go sendToSlack(msg, slackConfig)
	}
}

func sendToTelegram(msg watchers.WatcherResult, config watchers.TransportTelegram) {
	transports.SendTelegramMessage(msg, config)
}

func sendToSlack(msg watchers.WatcherResult, config watchers.TransportSlack) {
	transports.SendSlackMessage(msg, config)
}
