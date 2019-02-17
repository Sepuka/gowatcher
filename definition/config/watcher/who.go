package watcher

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
)

const (
	DefWatcherConfigWho = "definition.config.watcher.who"
	whoAgentName        = "Who"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefWatcherConfigWho,
			Build: func(ctn di.Container) (interface{}, error) {
				return config.GetWatcherConfig(whoAgentName), nil
			},
		})
	})
}
