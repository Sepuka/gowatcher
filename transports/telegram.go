package transports

import (
	"fmt"
	"net/http"
	"io"
	"github.com/gravitational/log"
	"os"
	"github.com/sepuka/gowatcher/watchers"
	"github.com/sepuka/gowatcher/pack"
)

const (
	telegramPathTemplate = "%v/%v:%v/sendMessage"
	formatJson = "application/json"
	textModeHTML = "HTML"
	textModeMarkdown = "Markdown"
)

func SendMessage(data watchers.WatcherResult, config watchers.TransportTelegram) (resp *http.Response, err error) {
	telegramApi := config.Api
	url := fmt.Sprintf(telegramPathTemplate, telegramApi, config.BotId, config.Token)
	body := buildRequest(data, config)

	return http.Post(url, config.Format, body)
}

func SendUrgentMessage(data watchers.WatcherResult, config watchers.TransportTelegram) (resp *http.Response, err error) {
	urgent := config
	urgent.SilentNotify=false

	return SendMessage(data, urgent)
}

func buildRequest(data watchers.WatcherResult, config watchers.TransportTelegram) io.Reader {
	text := formatText(data, config)
	d := map[string]interface{}{
		"chat_id": config.ChatId,
		"text": text,
		"disable_notification": config.IsSilentNotify(),
		"parse_mode": config.TextMode,
	}

	switch config.Format {
		case formatJson:
			return pack.Encode(d)
		default:
			log.Errorf("Unknown telegram buildRequest %v!", config.Format)
			panic("Bad telegram buildRequest")
	}
}

func formatText(data watchers.WatcherResult, config watchers.TransportTelegram) string {
	host := getCurrentHost()

	switch config.TextMode {
		case textModeHTML:
			return fmt.Sprintf(
				"<strong>%v</strong> <b>%v</b> says: <code>%s</code>",
					host,
					data.GetName(),
					data.GetText())
		case textModeMarkdown:
			return fmt.Sprintf(
				"%v *%v* says:\n ```%s```",
				host,
				data.GetName(),
				data.GetText())
		default:
			return fmt.Sprintf(
				"%v %v says: %s",
				host,
				data.GetName(),
				data.GetText())
	}
}

func getCurrentHost() string {
	host, err := os.Hostname()
	if err != nil {
		log.Warningf("Cannot detect current host: %v", err)
		host = "unknown host"
	}

	return host
}