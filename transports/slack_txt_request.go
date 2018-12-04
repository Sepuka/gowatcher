package transports

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/pack"
	"io/ioutil"
	"net/http"
)

func (obj Slack) sendText(httpClient *http.Client, msg command.Result, url string, format pack.FormatMode) (resp *http.Response, err error) {
	d := map[string]interface{}{
		"text": pack.FormatText(msg, format),
	}

	payload := pack.Encode(d)
	req, _ := http.NewRequest("POST", url, payload)

	resp, err = httpClient.Do(req)
	if err != nil {
		obj.logger.Errorf("Failed slack request %v", err)
		return nil, err
	}

	obj.logger.Debugf("Slack request to '%v' got '%v' status", url, resp.Status)
	_, respErr := ioutil.ReadAll(resp.Body)

	if respErr != nil {
		obj.logger.Error(respErr)
	}

	defer resp.Body.Close()

	return resp, err
}
