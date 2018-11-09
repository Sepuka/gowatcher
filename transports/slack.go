package transports

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"net/http"
)

func SendSlackMessage(msg command.Result, config config.TransportSlack) {
	client := &http.Client{}

	switch msg.GetType() {
	case command.ImageContent:
		sendImg(client, msg, config)
	default:
		sendText(client, msg, config.Api, config.TextMode)
	}
}
