package watchers

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"io"
	"github.com/gravitational/log"
	"os"
)

const (
	telegramPathTemplate = "%v/%v:%v/sendMessage"
	formatJson = "application/json"
	textModeHTML = "HTML"
	textModeMarkdown = "Markdown"
)

func SendMessage(data WatcherResult, config TelegramConfig) (resp *http.Response, err error) {
	telegramApi := config.Api
	url := fmt.Sprintf(telegramPathTemplate, telegramApi, config.BotId, config.Token)
	body := buildRequest(data, config)

	return http.Post(url, config.Format, body)
}

func SendUrgentMessage(data WatcherResult, config TelegramConfig) (resp *http.Response, err error) {
	urgent := config
	urgent.SilentNotify=false

	return SendMessage(data, urgent)
}

func buildRequest(data WatcherResult, config TelegramConfig) io.Reader {
	text := formatText(data, config)
	d := map[string]string{
		"chat_id": config.ChatId,
		"text": text,
		"disable_notification": config.IsSilentNotify(),
		"parse_mode": config.TextMode,
	}

	switch config.Format {
		case formatJson:
			out := new(bytes.Buffer)
			json.NewEncoder(out).Encode(d)
			return out
		default:
			log.Errorf("Unknown telegram buildRequest %v!", config.Format)
			panic("Bad telegram buildRequest")
	}
}

func formatText(data WatcherResult, config TelegramConfig) string {
	host := getCurrentHost()

	switch config.TextMode {
		case textModeHTML:
			return fmt.Sprintf(
				"<strong>%v</strong> <b>%v</b> says: <code>%s</code>",
					host,
					data.watcherName,
					data.GetText())
		case textModeMarkdown:
			return fmt.Sprintf(
				"%v *%v* says:\n ```%s```",
				host,
				data.watcherName,
				data.GetText())
		default:
			return fmt.Sprintf(
				"%v %v says: %s",
				host,
				data.watcherName,
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