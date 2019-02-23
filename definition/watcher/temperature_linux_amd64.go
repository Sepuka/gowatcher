package watcher

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	config2 "github.com/sepuka/gowatcher/definition/config"
	"github.com/sepuka/gowatcher/definition/transport"
	"github.com/sepuka/gowatcher/domain"
	"github.com/sepuka/gowatcher/services"
)

const (
	DefWatcherTemperature = "definition.watcher.temperature"
	tempCommand = "sensors"
	withoutParseChipNameArg = "-A"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefWatcherTemperature,
			Tags: []di.Tag{{
				Name: DefWatcherTag,
			}},
			Build: func(ctn di.Container) (interface{}, error) {
				var (
					cfg config.WatcherConfig
					transportChan chan<- command.Result
				)

				if err := services.Container.Fill(config2.DefWatcherConfigTemperature, &cfg); err != nil {
					return nil, err
				}

				if err := services.Container.Fill(transport.DefTransportChan, &transportChan); err != nil {
					return nil, err
				}

				return &domain.TemperatureWatcher{
					Command: command.NewCmd(tempCommand, withoutParseChipNameArg),
					Loop: cfg.GetLoop(),
					Handler: command.NewDummyResultHandler(transportChan),
				}, nil
			},
		})
	})
}
