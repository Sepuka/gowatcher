package config

import (
	"os"
	"github.com/stevenroose/gonfig"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	configPath = "./config.json"
	watcherName = "name"
	watcherLoop = "loop"
	TextModeHTML     FormatMode = "html"
	TextModeMarkdown FormatMode = "markdown"
	TextModeRaw      FormatMode = "raw"
	//https://get.slack.help/hc/en-us/articles/202288908-how-can-i-add-formatting-to-my-messages-
	TextModeSlack    FormatMode = "slack"
)

var (
	TelegramConfig TransportTelegram
	SlackConfig    TransportSlack
	WatchersConfig []WatcherConfig
	config         configuration
)

type FormatMode string

type TransportTelegram struct {
	Api          string     `id:"api"`
	Format       string     `id:"format"`
	SilentNotify bool       `id:"silentNotify" default:true`
	TextMode     FormatMode `id:"textMode"`
	ChatId       string     `id:"chatId"`
	BotId        string     `id:"botId"`
	Token        string     `id:"token"`
}

func (r *TransportTelegram) IsSilentNotify() string {
	return strconv.FormatBool(r.SilentNotify)
}

type TransportSlack struct {
	Api      string     `id:"api" default:"https://slack.com/api"`
	TextMode FormatMode `id:"textMode"`
}

type configuration struct {
	Transports map[string]interface{} `id:"transports"`
	Watchers []map[string]interface{} `id:"watchers"`
}

func InitConfig() {
	readConfig()

	tconf := config.Transports["telegram"].(map[string]interface{})
	gonfig.LoadMap(&TelegramConfig, tconf, gonfig.Conf{})
	TelegramConfig.TextMode = getTextMode(tconf["textMode"].(string))

	sconf := config.Transports["slack"].(map[string]interface{})
	gonfig.LoadMap(&SlackConfig, sconf, gonfig.Conf{})
	SlackConfig.TextMode = getTextMode(sconf["textMode"].(string))

	for _, watcher := range config.Watchers {
		a := WatcherConfig{
			watcher[watcherName].(string),
			time.Duration(watcher[watcherLoop].(float64)) * time.Second,
		}
		WatchersConfig = append(WatchersConfig, a)
	}
}

func readConfig() {
	err := gonfig.Load(&config, gonfig.Conf{
		FileDefaultFilename: configPath,
		FileDecoder:         gonfig.DecoderJSON,
		FlagIgnoreUnknown:   true,
	})
	if err != nil {
		log.Printf("Cannot read config: %v", err)
		os.Exit(1)
	}
}

func getTextMode(mode string) FormatMode {
	switch strings.ToLower(mode) {
	case "html":
		return TextModeHTML
	case "markdown":
		return TextModeMarkdown
	case "slack":
		return TextModeSlack
	default:
		return TextModeRaw
	}
}