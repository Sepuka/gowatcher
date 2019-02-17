package watcher

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/definition/handler"
	"github.com/sepuka/gowatcher/definition/store"
	"github.com/sepuka/gowatcher/domain"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/stats"
)

const (
	DefCollectorLoadAverage = "definition.collector.la"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefCollectorLoadAverage,
			Tags: []di.Tag{{
				Name: DefWatcherTag,
			}},
			Build: func(ctn di.Container) (interface{}, error) {
				var (
					laHandler command.ResultHandler
					redis stats.SliceStoreWriter
				)

				if err := services.Container.Fill(handler.DefHandlerLoadAverage, &laHandler); err != nil {
					return nil, err
				}

				if err := services.Container.Fill(store.DefStoreRedis, &redis); err != nil {
					return nil, err
				}

				return &domain.LaCollector{
					Handler: laHandler,
					Store: redis,
				}, nil
			},
		})
	})
}
