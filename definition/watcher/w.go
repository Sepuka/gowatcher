package watcher

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/definition/transport"
	"github.com/sepuka/gowatcher/domain"
	"github.com/sepuka/gowatcher/services"
)

const (
	DefWatcherW = "definition.watcher.w"
	wCommand    = "w"
	wAgentName  = "W"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefWatcherW,
			Tags: []di.Tag{{
				Name: DefWatcherTag,
			}},
			Build: func(ctn di.Container) (interface{}, error) {
				var (
					cfg           config.WatcherConfig
					transportChan chan<- command.Result
				)

				if err := params.Fill(wAgentName, &cfg); err != nil {
					return nil, err
				}

				if err := services.Container.Fill(transport.DefTransportChan, &transportChan); err != nil {
					return nil, err
				}

				return &domain.WWatcher{
					Command:       command.NewCmd(wCommand, cfg.Args),
					Loop:          cfg.GetLoop(),
					TransportChan: transportChan,
				}, nil
			},
		})
	})
}
