package transports

import (
	"fmt"
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/pack"
	"github.com/sepuka/gowatcher/services"
	"github.com/sirupsen/logrus"
	"github.com/stevenroose/gonfig"
	"io"
	"net/http"
	"strconv"
)

const (
	telegramPathTemplate = "%v/%v:%v/sendMessage"
	transportTelegramName TransportName = "telegramm"
)

var telegramCfg TelegramConfig

type TelegramConfig struct {
	Api          string          `id:"Api"`
	SilentNotify bool            `id:"SilentNotify" default:true`
	TextMode     pack.FormatMode `id:"TextMode"`
	ChatId       string          `id:"ChatId"`
	BotId        string          `id:"BotId"`
	Token        string          `id:"Token"`
}

type Telegram struct {
	httpClient *http.Client
	cfg        *TelegramConfig
	logger     logrus.StdLogger
}

func (obj Telegram) Send(msg command.Result) (resp *http.Response, err error) {
	url := fmt.Sprintf(telegramPathTemplate, obj.cfg.Api, obj.cfg.BotId, obj.cfg.Token)
	body := obj.buildRequest(msg)

	return http.Post(url, contentTypeJson, body)
}

func (obj Telegram) GetName() TransportName {
	return transportTelegramName
}

func (obj Telegram) buildRequest(data command.Result) io.Reader {
	text := pack.FormatText(data, obj.cfg.TextMode)
	d := map[string]interface{}{
		"chat_id":              obj.cfg.ChatId,
		"text":                 text,
		"disable_notification": obj.isSilentNotify(),
		"parse_mode":           obj.cfg.TextMode,
	}

	return pack.Encode(d)
}

func (obj Telegram) isSilentNotify() string {
	return strconv.FormatBool(obj.cfg.SilentNotify)
}

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		cfg := params.Transports["telegram"].(map[string]interface{})
		gonfig.LoadMap(&telegramCfg, cfg, gonfig.Conf{})
		telegramCfg.TextMode = pack.GetTextMode(cfg["textMode"].(string))

		return builder.Add(di.Def{
			Name: services.Telegram,
			Build: func(ctn di.Container) (interface{}, error) {
				return &Telegram{
					&http.Client{},
					&telegramCfg,
					services.Container.Get(services.Logger).(*logrus.Logger),
				}, nil
			},
		})
	})
}
