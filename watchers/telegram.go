package watchers

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"io"
	"github.com/gravitational/log"
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
	switch config.Transport.TextMode {
		case textModeHTML:
			return fmt.Sprintf("watcher <b>%v</b> says:<br> <strong>%s</strong>", data.watcherName, data.GetText())
		case textModeMarkdown:
			return fmt.Sprintf("watcher *%v* says:\n ```%s```", data.watcherName, data.GetText())
		default:
			return fmt.Sprintf("watcher %v says: %s", data.watcherName, data.GetText())
	}
}