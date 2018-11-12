package config

import (
	"github.com/go-redis/redis"
	"github.com/stevenroose/gonfig"
	"log"
	"os"
	"strconv"
	"strings"
)

type FormatMode string
type LogLevel string

const (
	configPath                  = "./config.json"
	TextModeHTML     FormatMode = "html"
	TextModeMarkdown FormatMode = "markdown"
	TextModeRaw      FormatMode = "raw"
	//https://get.slack.help/hc/en-us/articles/202288908-how-can-i-add-formatting-to-my-messages-
	TextModeSlack   FormatMode = "slack"
	LogLevelDebug   LogLevel   = "debug"
	LogLevelDefault LogLevel   = "default"
)

type Configuration struct {
	Transports    map[string]interface{}   `id:"transports"`
	Watchers      []map[string]interface{} `id:"watchers"`
	KeyValueStore map[string]interface{}   `id:"redis"`
	Logger        Logger                   `id:"log"`
}

var (
	TelegramConfig TransportTelegram
	SlackConfig    TransportSlack
	WatchersConfig []WatcherConfig
	Redis          RedisStore
	Log            Logger
	AppConfig      Configuration
)

type TransportTelegram struct {
	Api          string     `id:"api"`
	Format       string     `id:"format"`
	SilentNotify bool       `id:"silentNotify" default:true`
	TextMode     FormatMode `id:"textMode"`
	ChatId       string     `id:"chatId"`
	BotId        string     `id:"botId"`
	Token        string     `id:"token"`
}

func (r TransportTelegram) IsSilentNotify() string {
	return strconv.FormatBool(r.SilentNotify)
}

type TransportSlack struct {
	Api           string     `id:"api" default:"https://slack.com/api"`
	TextMode      FormatMode `id:"textMode"`
	FileUploadUrl string     `id:"fileUploadUrl" default:"https://slack.com/api/files.upload"`
	Token         string     `id:"token"`
}

type Logger struct {
	Level LogLevel `id:"level" default:"default"`
}

func InitConfig() {
	readConfig()

	tconf := AppConfig.Transports["telegram"].(map[string]interface{})
	gonfig.LoadMap(&TelegramConfig, tconf, gonfig.Conf{})
	TelegramConfig.TextMode = getTextMode(tconf["textMode"].(string))

	sconf := AppConfig.Transports["slack"].(map[string]interface{})
	gonfig.LoadMap(&SlackConfig, sconf, gonfig.Conf{})
	SlackConfig.TextMode = getTextMode(sconf["textMode"].(string))

	redisAddr := AppConfig.KeyValueStore["address"].(string)
	redisPass := AppConfig.KeyValueStore["password"].(string)
	redisDb, _ := AppConfig.KeyValueStore["db"].(float64)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       int(redisDb),
	})
	Redis = RedisStore{redisClient}

	Log = AppConfig.Logger // переписать это все

	initWatcherConfigs()
}

func readConfig() {
	err := gonfig.Load(&AppConfig, gonfig.Conf{
		FileDefaultFilename: configPath,
		FileDecoder:         gonfig.DecoderJSON,
		FlagIgnoreUnknown:   true,
	})
	if err != nil {
		log.Printf("Cannot read AppConfig: %v", err)
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
