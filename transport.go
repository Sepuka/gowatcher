package main

import (
	"github.com/sepuka/gowatcher/watchers"
	"github.com/stevenroose/gonfig"
	"github.com/sepuka/gowatcher/transports"
)

func Transmitter(c <-chan watchers.WatcherResult) {
	for {
		msg := <- c
		telegramConfig := watchers.TelegramConfig{}
		conf := config.Transports["telegram"].(map[string]interface{})
		gonfig.LoadMap(&telegramConfig, conf, gonfig.Conf{})
		transports.SendMessage(msg, telegramConfig)
	}
}
