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

func SendMessage(data WatcherResult, config Configuration) (resp *http.Response, err error) {
	telegramApi := config.Transport.Api
	url := fmt.Sprintf(telegramPathTemplate, telegramApi, config.BotId, config.Token)
	body := buildRequest(data, config)

	return http.Post(url, config.Transport.Format, body)
}

func SendUrgentMessage(data WatcherResult, config Configuration) (resp *http.Response, err error) {
	urgent := config
	urgent.Transport.SilentNotify=false

	return SendMessage(data, urgent)
}

func buildRequest(data WatcherResult, config Configuration) io.Reader {
	text := formatText(data, config)
	d := map[string]string{
		"chat_id": config.ChatId,
		"text": text,
		"disable_notification": config.Transport.IsSilentNotify(),
		"parse_mode": config.Transport.TextMode,
	}

	switch config.Transport.Format {
		case formatJson:
			out := new(bytes.Buffer)
			json.NewEncoder(out).Encode(d)
			return out
		default:
			log.Errorf("Unknown telegram buildRequest %v!", config.Transport.Format)
			panic("Bad telegram buildRequest")
	}
}

func formatText(data WatcherResult, config Configuration) string {
	host := getCurrentHost()

	switch config.Transport.TextMode {
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