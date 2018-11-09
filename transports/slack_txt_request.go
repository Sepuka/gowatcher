package transports

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/pack"
	"io/ioutil"
	"log"
	"net/http"
)

func sendText(httpClient *http.Client, msg command.Result, url string, format config.FormatMode) {
	d := map[string]interface{}{
		"text": pack.FormatText(msg, format),
	}

	payload := pack.Encode(d)
	req, _ := http.NewRequest("POST", url, payload)

	resp, _ := httpClient.Do(req)
	log.Printf("Slack request to '%v' got '%v' status", url, resp.Status)
	_, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
}
