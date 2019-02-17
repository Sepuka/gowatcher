package config

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"github.com/stevenroose/gonfig"
)

const DefConfigTelegram = "definition.config.telegram"

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		var (
			telegramCfg config.TelegramConfig
		)
		cfg := params.Transports["telegram"].(map[string]interface{})
		gonfig.LoadMap(&telegramCfg, cfg, gonfig.Conf{})

		return builder.Add(di.Def{
			Name: DefConfigTelegram,
			Build: func(ctn di.Container) (interface{}, error) {
				return &telegramCfg, nil
			},
		})
	})
}
