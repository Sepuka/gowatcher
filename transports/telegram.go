package transports

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/definition/logger"
	"github.com/sepuka/gowatcher/pack"
	"go.uber.org/zap"
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
	logger     logger.Logger
}

func NewTelegram(http *http.Client, cfg *config.TelegramConfig, logger logger.Logger) Transport {
	logger.With(zap.String("transport", transportTelegramName))

	return &Telegram{
		httpClient: http,
		cfg:        cfg,
		logger:     logger,
	}
}

func (obj Telegram) Send(msg command.Result) (err error) {
	obj.
		logger.
		With(
			zap.String("msg_type", string(msg.GetType())),
			zap.String("msg", msg.GetName()),
		).
		Debug("Sending message.")

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
