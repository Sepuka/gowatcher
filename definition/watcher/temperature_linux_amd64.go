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
	DefWatcherTemperature   = "definition.watcher.temperature"
	tempCommand             = "sensors"
	withoutParseChipNameArg = "-A"
	tempAgentName           = "Temp"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		var (
			cfg           config.WatcherConfig
			transportChan chan<- command.Result
		)

		if err := params.Fill(tempAgentName, &cfg); err != nil {
			return err
		}

		if cfg.IsActive == false {
			return nil
		}

		return builder.Add(di.Def{
			Name: DefWatcherTemperature,
			Tags: []di.Tag{{
				Name: DefWatcherTag,
			}},
			Build: func(ctn di.Container) (interface{}, error) {

				if err := services.Container.Fill(transport.DefTransportChan, &transportChan); err != nil {
					return nil, err
				}

				return &domain.TemperatureWatcher{
					Command: command.NewCmd(tempCommand, withoutParseChipNameArg),
					Loop:    cfg.GetLoop(),
					Handler: command.NewDummyResultHandler(transportChan),
				}, nil
			},
		})
	})
}
