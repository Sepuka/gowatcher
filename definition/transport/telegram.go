package transport

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	config2 "github.com/sepuka/gowatcher/definition/config"
	http2 "github.com/sepuka/gowatcher/definition/http"
	logger2 "github.com/sepuka/gowatcher/definition/logger"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/transports"
	"github.com/sirupsen/logrus"
	"net/http"
)

const DefTransportTelegram = "definition.transport.telegram"

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		var (
			cfg        *config.TelegramConfig
			logger     *logrus.Logger
			httpClient *http.Client
		)

		return builder.Add(di.Def{
			Name: DefTransportTelegram,
			Tags: []di.Tag{{
				Name: DefTransportTag,
			}},
			Build: func(ctn di.Container) (interface{}, error) {
				if err := ctn.Fill(logger2.DefLogger, &logger); err != nil {
					return nil, err
				}

				if err := ctn.Fill(http2.DefHttpClient, &httpClient); err != nil {
					return nil, err
				}

				if err := ctn.Fill(config2.DefConfigTelegram, &cfg); err != nil {
					return nil, err
				}

				return transports.NewTelegram(
					httpClient,
					cfg,
					logger,
				), nil
			},
		})
	})
}
