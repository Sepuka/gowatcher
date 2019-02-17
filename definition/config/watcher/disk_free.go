package watcher

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
)

const (
	DefWatcherConfigDiskFree = "definition.config.watcher.disk_free"
	dfAgentName              = "DiskFree"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefWatcherConfigDiskFree,
			Build: func(ctn di.Container) (interface{}, error) {
				return config.GetWatcherConfig(dfAgentName), nil
			},
		})
	})
}
