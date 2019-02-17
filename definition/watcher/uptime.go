package watcher

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/definition/config/watcher"
	"github.com/sepuka/gowatcher/definition/transport"
	"github.com/sepuka/gowatcher/domain"
	"github.com/sepuka/gowatcher/services"
)

const (
	DefWatcherUptime = "definition.watcher.uptime"
	uptimeCommand = "uptime"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefWatcherUptime,
			Tags: []di.Tag{{
				Name: DefWatcherTag,
			}},
			Build: func(ctn di.Container) (interface{}, error) {
				var (
					cfg config.WatcherConfig
					transportChan chan<- command.Result
				)

				if err := services.Container.Fill(watcher.DefWatcherConfigUptime, &cfg); err != nil {
					return nil, err
				}

				if err := services.Container.Fill(transport.DefTransportChan, &transportChan); err != nil {
					return nil, err
				}

				return &domain.UptimeWatcher{
					Command: command.NewCmd(uptimeCommand, cfg.Args),
					Loop: cfg.GetLoop(),
					TransportChan: transportChan,
				}, nil
			},
		})
	})
}
