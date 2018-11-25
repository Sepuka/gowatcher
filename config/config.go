package config

import (
	"github.com/stevenroose/gonfig"
	"log"
	"os"
	"time"
)

const (
	configPath = "./config.json"
)

type Logger struct {
	Level      string `id:"level" default:"info"`
	WithCaller bool   `id:"withCaller" default:"false"`
	File       string `id:"file" default:"log"`
}

type KeyValue struct {
	Host     string `id:"host" default:"localhost"`
	Port     int16  `id:"port" default:"6379"`
	Password string `id:"password" default:""`
	Db       int    `id:"db" default:"0"`
}

type WatcherConfig struct {
	Name string        `id:"name"`
	Loop time.Duration `id:"loop" default:"86400"`
	Args string        `id:"args"`
}

func (setting WatcherConfig) GetName() string {
	return setting.Name
}

func (setting WatcherConfig) GetLoop() time.Duration {
	return setting.Loop * time.Second
}

type Configuration struct {
	Transports    map[string]interface{} `id:"transports"`
	Watchers      []WatcherConfig        `id:"watchers"`
	KeyValueStore KeyValue               `id:"redis"`
	Logger        Logger                 `id:"log"`
}

func GetWatcherConfig(watcherName string) WatcherConfig {
	for _, cfg := range AppConfig.Watchers {
		if cfg.GetName() == watcherName {
			return cfg
		}
	}

	panic("Cannot find watcher config " + watcherName)
}

var (
	AppConfig Configuration
)

func init() {
	readConfig()
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
