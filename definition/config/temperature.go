package config

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
)

const (
	DefWatcherConfigTemperature = "definition.config.watcher.temperature"
	tempAgentName               = "Temp"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefWatcherConfigTemperature,
			Build: func(ctn di.Container) (interface{}, error) {
				return config.GetWatcherConfig(tempAgentName), nil
			},
		})
	})
}
