package config

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
)

const (
	DefWatcherConfigLoadAverage = "definition.config.watcher.la"
	laAgentName        = "LoadAvgGraph"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefWatcherConfigLoadAverage,
			Build: func(ctn di.Container) (interface{}, error) {
				return config.GetWatcherConfig(laAgentName), nil
			},
		})
	})
}
