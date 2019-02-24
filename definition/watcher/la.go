package watcher

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/definition/store"
	"github.com/sepuka/gowatcher/definition/transport"
	"github.com/sepuka/gowatcher/domain"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/stats"
)

const (
	DefLoadAverage = "definition.watcher.la"
	laAgentName    = "LoadAvgGraph"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		var (
			cfg           config.WatcherConfig
			transportChan chan<- command.Result
			redis         stats.SliceStoreReader
		)

		if err := params.Fill(laAgentName, &cfg); err != nil {
			return err
		}

		if cfg.IsActive == false {
			return nil
		}

		return builder.Add(di.Def{
			Name: DefLoadAverage,
			Tags: []di.Tag{{
				Name: DefWatcherTag,
			}},
			Build: func(ctn di.Container) (interface{}, error) {

				if err := services.Container.Fill(transport.DefTransportChan, &transportChan); err != nil {
					return nil, err
				}

				if err := services.Container.Fill(store.DefStoreRedis, &redis); err != nil {
					return nil, err
				}

				return &domain.LoadAvgGraphWatcher{
					Loop:    cfg.GetLoop(),
					Handler: command.NewDummyResultHandler(transportChan),
					Redis:   redis,
				}, nil
			},
		})
	})
}
