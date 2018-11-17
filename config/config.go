package config

import (
	"github.com/stevenroose/gonfig"
	"log"
	"os"
)

const (
	configPath = "./config.json"
)

type Logger struct {
	Level string `id:"level" default:"info"`
}

type KeyValue struct {
	Host     string `id:"host" default:"localhost"`
	Port     int16  `id:"port" default:6379`
	Password string `id:"password" default:""`
	Db       int    `id:"db" default:0`
}

type Configuration struct {
	Transports    map[string]interface{}   `id:"transports"`
	Watchers      []map[string]interface{} `id:"watchers"`
	KeyValueStore KeyValue                 `id:"redis"`
	Logger        Logger                   `id:"log"`
}

var (
	WatchersConfig []WatcherConfig
	AppConfig      Configuration
)

func InitConfig() {
	readConfig()
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
