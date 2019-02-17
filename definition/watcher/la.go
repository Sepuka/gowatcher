package watcher

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	config2 "github.com/sepuka/gowatcher/definition/config"
	"github.com/sepuka/gowatcher/definition/store"
	"github.com/sepuka/gowatcher/definition/transport"
	"github.com/sepuka/gowatcher/domain"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/stats"
)

const (
	DefLoadAverage = "definition.watcher.la"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefLoadAverage,
			Tags: []di.Tag{{
				Name: DefWatcherTag,
			}},
			Build: func(ctn di.Container) (interface{}, error) {
				var (
					cfg config.WatcherConfig
					transportChan chan<- command.Result
					redis stats.SliceStoreReader
				)

				if err := services.Container.Fill(config2.DefWatcherConfigLoadAverage, &cfg); err != nil {
					return nil, err
				}

				if err := services.Container.Fill(transport.DefTransportChan, &transportChan); err != nil {
					return nil, err
				}

				if err := services.Container.Fill(store.DefStoreRedis, &redis); err != nil {
					return nil, err
				}

				return &domain.LoadAvgGraphWatcher{
					Loop: cfg.GetLoop(),
					Handler: command.NewDummyResultHandler(transportChan),
					Redis: redis,
				}, nil
			},
		})
	})
}
