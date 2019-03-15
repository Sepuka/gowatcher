package handler

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/definition/logger"
	"github.com/sepuka/gowatcher/definition/store"
	"github.com/sepuka/gowatcher/definition/transport"
	"github.com/sepuka/gowatcher/domain/handler"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/stats"
)

const (
	DefHandlerLoadAverage = "definition.handler.la"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefHandlerLoadAverage,
			Build: func(ctn di.Container) (interface{}, error) {
				var (
					transportChan chan<- command.Result
					redis stats.SliceStoreWriter
					log logger.Logger
				)

				if err := services.Container.Fill(transport.DefTransportChan, &transportChan); err != nil {
					return nil, err
				}

				if err := services.Container.Fill(store.DefStoreRedis, &redis); err != nil {
					return nil, err
				}

				if err := ctn.Fill(logger.DefLogger, &log); err != nil {
					return nil, err
				}

				return handler.NewLaResultHandler(transportChan, redis, log), nil
			},
		})
	})
}
