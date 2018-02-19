package watchers

import (
	"fmt"
	"net/http"
	"io"
)

const (
	telegramApi          = "https://api.telegram.org"
	telegramPathTemplate = "%v/%v:%v/sendMessage"
)

func SendMessage(payload io.Reader, config Configuration) (resp *http.Response, err error) {
	url := fmt.Sprintf(telegramPathTemplate, telegramApi, config.BotId, config.Token)

	return http.Post(url, "application/json", payload)
}
