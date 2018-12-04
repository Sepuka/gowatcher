package transports

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/pack"
	"github.com/sepuka/gowatcher/services"
	"github.com/sirupsen/logrus"
	"github.com/stevenroose/gonfig"
	"log"
	"net/http"
)

const transportSlackName TransportName = "slack"

var slackCfg slackConfig

type slackConfig struct {
	Api           string          `id:"api" default:"https://slack.com/Api"`
	TextMode      pack.FormatMode `id:"textMode"`
	FileUploadUrl string          `id:"fileUploadUrl" default:"https://slack.com/Api/files.upload"`
	Token         string          `id:"token"`
}

type Slack struct {
	httpClient *http.Client
	cfg        *slackConfig
	logger     logrus.FieldLogger
}

func (obj Slack) Send(msg command.Result) (resp *http.Response, err error) {
	switch msg.GetType() {
	case command.ImageContent:
		return obj.sendImg(obj.httpClient, msg, obj.cfg)
	default:
		return obj.sendText(obj.httpClient, msg, obj.cfg.Api, obj.cfg.TextMode)
	}
}

func (obj Slack) GetName() TransportName {
	return transportSlackName
}

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		cfg := params.Transports["slack"].(map[string]interface{})
		err := gonfig.LoadMap(&slackCfg, cfg, gonfig.Conf{})
		if err != nil {
			log.Fatalf("Cannot instantiate slack configuration: %v", err)
			return err
		}

		slackCfg.TextMode = pack.GetTextMode(cfg["textMode"].(string))

		return builder.Add(di.Def{
			Name: services.Slack,
			Build: func(ctn di.Container) (interface{}, error) {
				return &Slack{
					&http.Client{},
					&slackCfg,
					services.Container.Get(services.Logger).(*logrus.Logger),
				}, nil
			},
		})
	})
}
