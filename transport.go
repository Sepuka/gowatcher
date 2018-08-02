package main

import (
	"github.com/sepuka/gowatcher/watchers"
	"github.com/stevenroose/gonfig"
	"github.com/sepuka/gowatcher/transports"
)

func Transmitter(c <-chan watchers.WatcherResult) {
	for {
		msg := <- c

		telegramConfig := watchers.TransportTelegram{}
		tconf := config.Transports["telegram"].(map[string]interface{})
		gonfig.LoadMap(&telegramConfig, tconf, gonfig.Conf{})
		go sendToTelegram(msg, telegramConfig)

		slackConfig := watchers.TransportSlack{}
		sconf := config.Transports["slack"].(map[string]interface{})
		gonfig.LoadMap(&slackConfig, sconf, gonfig.Conf{})
		go sendToSlack(msg, slackConfig)
	}
}

func sendToTelegram(msg watchers.WatcherResult, config watchers.TransportTelegram) {
	transports.SendMessage(msg, config)
}

func sendToSlack(msg watchers.WatcherResult, config watchers.TransportSlack) {
	transports.SendMessage2(msg, config)
}