package transports

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/pack"
	"github.com/sepuka/gowatcher/services"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func sendText(httpClient *http.Client, msg command.Result, url string, format pack.FormatMode) (resp *http.Response, err error) {
	log := services.Container.Get(services.Logger).(*logrus.Logger)
	d := map[string]interface{}{
		"text": pack.FormatText(msg, format),
	}

	payload := pack.Encode(d)
	req, _ := http.NewRequest("POST", url, payload)

	resp, err = httpClient.Do(req)
	if err != nil {
		log.Errorf("Failed slack request %v", err)
		return nil, err
	}

	log.Debugf("Slack request to '%v' got '%v' status", url, resp.Status)
	_, respErr := ioutil.ReadAll(resp.Body)

	if respErr != nil {
		log.Error(respErr)
	}

	defer resp.Body.Close()

	return resp, err
}
