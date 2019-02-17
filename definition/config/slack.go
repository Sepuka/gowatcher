package config

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"github.com/stevenroose/gonfig"
)

const DefConfigSlack = "definition.config.slack"

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		var (
			slackCfg config.SlackConfig
		)
		cfg := params.Transports["slack"].(map[string]interface{})
		if err := gonfig.LoadMap(&slackCfg, cfg, gonfig.Conf{}); err != nil {
			return err
		}

		return builder.Add(di.Def{
			Name: DefConfigSlack,
			Build: func(ctn di.Container) (interface{}, error) {
				return &slackCfg, nil
			},
		})
	})
}
