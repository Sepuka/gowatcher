package watcher

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
)

const (
	DefWatcherConfigUptime = "definition.config.watcher.uptime"
	uptimeAgentName        = "Uptime"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefWatcherConfigUptime,
			Build: func(ctn di.Container) (interface{}, error) {
				return config.GetWatcherConfig(uptimeAgentName), nil
			},
		})
	})
}
