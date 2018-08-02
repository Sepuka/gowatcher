package transports

import (
	"github.com/sepuka/gowatcher/watchers"
	"net/http"
	"bytes"
	"encoding/json"
)

func SendMessage2(msg watchers.WatcherResult, config watchers.TransportSlack) {
	d := map[string]string{
		"text": msg.GetText(),
	}
	out := new(bytes.Buffer)
	json.NewEncoder(out).Encode(d)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", config.Api, out)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}