package transport

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	configDef "github.com/sepuka/gowatcher/definition/config"
	httpDef "github.com/sepuka/gowatcher/definition/http"
	"github.com/sepuka/gowatcher/definition/logger"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/transports"
	"net/http"
)

const DefTransportSlack = "definition.transport.slack"

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefTransportSlack,
			Tags: []di.Tag{{
				Name: DefTransportTag,
			}},
			Build: func(ctn di.Container) (interface{}, error) {
				var (
					log        logger.Logger
					httpClient *http.Client
					cfg        *config.SlackConfig
				)

				if err := ctn.Fill(logger.DefLogger, &log); err != nil {
					return nil, err
				}

				if err := ctn.Fill(configDef.DefConfigSlack, &cfg); err != nil {
					return nil, err
				}

				if err := ctn.Fill(httpDef.DefHttpClient, &httpClient); err != nil {
					return nil, err
				}

				return transports.NewSlack(httpClient, cfg, log), nil
			},
		})
	})
}
