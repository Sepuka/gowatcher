package transports

import (
	"github.com/sepuka/gowatcher/watchers"
	"net/http"
	"github.com/sepuka/gowatcher/pack"
)

func SendSlackMessage(msg watchers.WatcherResult, config watchers.TransportSlack) {
	d := map[string]interface{}{
		"text": pack.FormatText(msg, config.TextMode),
	}
	client := &http.Client{}
	payload := pack.Encode(d)
	req, _ := http.NewRequest("POST", config.Api, payload)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}

