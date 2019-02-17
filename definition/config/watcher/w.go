package watcher

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
)

const (
	DefWatcherConfigW = "definition.config.watcher.w"
	wAgentName        = "W"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefWatcherConfigW,
			Build: func(ctn di.Container) (interface{}, error) {
				return config.GetWatcherConfig(wAgentName), nil
			},
		})
	})
}
