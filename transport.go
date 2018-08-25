package main

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/transports"
	"github.com/sepuka/gowatcher/watchers"
	"github.com/stevenroose/gonfig"
	"strings"
)

var telegramConfig watchers.TransportTelegram
var slackConfig watchers.TransportSlack

func initConfig() {
	tconf := config.Transports["telegram"].(map[string]interface{})
	gonfig.LoadMap(&telegramConfig, tconf, gonfig.Conf{})
	telegramConfig.TextMode = getTextMode(tconf["textMode"].(string))

	sconf := config.Transports["slack"].(map[string]interface{})
	gonfig.LoadMap(&slackConfig, sconf, gonfig.Conf{})
	slackConfig.TextMode = getTextMode(sconf["textMode"].(string))
}

func Transmitter(c <-chan command.Result) {
	initConfig()

	for {
		msg := <-c

		go sendToTelegram(msg, telegramConfig)
		go sendToSlack(msg, slackConfig)
	}
}

func getTextMode(mode string) watchers.FormatMode {
	switch strings.ToLower(mode) {
	case "html":
		return watchers.TextModeHTML
	case "markdown":
		return watchers.TextModeMarkdown
	case "slack":
		return watchers.TextModeSlack
	default:
		return watchers.TextModeRaw
	}
}

func sendToTelegram(msg command.Result, config watchers.TransportTelegram) {
	transports.SendTelegramMessage(msg, config)
}

func sendToSlack(msg command.Result, config watchers.TransportSlack) {
	transports.SendSlackMessage(msg, config)
}
