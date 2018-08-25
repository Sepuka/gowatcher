package transports

import (
	"fmt"
	"github.com/gravitational/log"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/pack"
	"github.com/sepuka/gowatcher/watchers"
	"io"
	"net/http"
)

const (
	telegramPathTemplate = "%v/%v:%v/sendMessage"
	formatJson           = "application/json"
)

func SendTelegramMessage(data command.Result, config watchers.TransportTelegram) (resp *http.Response, err error) {
	telegramApi := config.Api
	url := fmt.Sprintf(telegramPathTemplate, telegramApi, config.BotId, config.Token)
	body := buildRequest(data, config)

	return http.Post(url, config.Format, body)
}

func SendUrgentMessage(data command.Result, config watchers.TransportTelegram) (resp *http.Response, err error) {
	urgent := config
	urgent.SilentNotify = false

	return SendTelegramMessage(data, urgent)
}

func buildRequest(data command.Result, config watchers.TransportTelegram) io.Reader {
	text := pack.FormatText(data, config.TextMode)
	d := map[string]interface{}{
		"chat_id":              config.ChatId,
		"text":                 text,
		"disable_notification": config.IsSilentNotify(),
		"parse_mode":           config.TextMode,
	}

	switch config.Format {
	case formatJson:
		return pack.Encode(d)
	default:
		log.Errorf("Unknown telegram buildRequest %v!", config.Format)
		panic("Bad telegram buildRequest")
	}
}
