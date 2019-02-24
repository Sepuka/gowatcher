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
	DefWatcherWho = "definition.watcher.who"
	whoCommand    = "who"
	whoAgentName  = "Who"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		var (
			cfg           config.WatcherConfig
			transportChan chan<- command.Result
		)

		if err := params.Fill(whoAgentName, &cfg); err != nil {
			return err
		}

		if cfg.IsActive == false {
			return nil
		}

		return builder.Add(di.Def{
			Name: DefWatcherWho,
			Tags: []di.Tag{{
				Name: DefWatcherTag,
			}},
			Build: func(ctn di.Container) (interface{}, error) {

				if err := services.Container.Fill(transport.DefTransportChan, &transportChan); err != nil {
					return nil, err
				}

				handler := command.NewLinesChangedResultHandler(transportChan)

				return &domain.WhoWatcher{
					Command: command.NewEnvedCmd(whoCommand, cfg.Args, "lang=en_EN"),
					Loop:    cfg.GetLoop(),
					Handler: handler,
				}, nil
			},
		})
	})
}
