package transports

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/pack"
	"net/http"
)

func (obj Slack) sendText(msg command.Result) (err error) {
	format := obj.cfg.TextMode
	url := obj.cfg.Api
	d := map[string]interface{}{
		"text": pack.FormatText(msg, format),
	}
	payload := pack.Encode(d)
	req, _ := http.NewRequest("POST", url, payload)

	sender := &loggedRequestSender{
		obj.httpClient,
		obj.logger,
		nil,
	}

	return sender.sendRequest(req)
}
