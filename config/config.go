package config

import (
	"github.com/go-redis/redis"
	"github.com/stevenroose/gonfig"
	"log"
	"os"
)

const (
	configPath                  = "./config.json"
)

type Configuration struct {
	Transports    map[string]interface{}   `id:"transports"`
	Watchers      []map[string]interface{} `id:"watchers"`
	KeyValueStore map[string]interface{}   `id:"redis"`
	Logger        Logger                   `id:"log"`
}

var (
	WatchersConfig []WatcherConfig
	Redis          RedisStore
	Log            Logger
	AppConfig      Configuration
)

type Logger struct {
	Level string `id:"level" default:"default"`
}

func InitConfig() {
	readConfig()

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
