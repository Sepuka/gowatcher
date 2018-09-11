package transports

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/pack"
	"net/http"
	"github.com/sepuka/gowatcher/config"
)

func SendSlackMessage(msg command.Result, config config.TransportSlack) {
	d := map[string]interface{}{
		"text": pack.FormatText(msg, config.TextMode),
	}
	client := &http.Client{}
	payload := pack.Encode(d)
	req, _ := http.NewRequest("POST", config.Api, payload)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}
