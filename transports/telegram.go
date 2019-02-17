package transports

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/pack"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

const (
	telegramPathTemplate  = "%v/%v:%v/sendMessage"
	transportTelegramName = "telegram"
)

type Telegram struct {
	httpClient *http.Client
	cfg        *config.TelegramConfig
	logger     logrus.FieldLogger
}

func NewTelegram(http *http.Client, cfg *config.TelegramConfig, logger *logrus.Logger) Transport {
	logger.WithFields(logrus.Fields{
		"transport": transportTelegramName,
	})

	return &Telegram{
		httpClient: http,
		cfg:        cfg,
		logger:     logger,
	}
}

func (obj Telegram) Send(msg command.Result) (err error) {
	obj.logger.WithFields(
		logrus.Fields{
			"msg_type": msg.GetType(),
		},
	).Debugf("Sending '%v' message.", msg.GetName())

	url := fmt.Sprintf(telegramPathTemplate, obj.cfg.Api, obj.cfg.BotId, obj.cfg.Token)
	body := obj.buildRequest(msg)

	_, err = http.Post(url, contentTypeJson, body)

	return err
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
