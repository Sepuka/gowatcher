package transports

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/pack"
	"net/http"
)

func (obj slack) sendText(msg command.Result) (err error) {
	url := obj.cfg.Api
	d := map[string]interface{}{
		"text": pack.FormatText(msg, obj.cfg.TextMode),
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
