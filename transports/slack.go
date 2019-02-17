package transports

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

const transportSlackName = "slack"

type slack struct {
	httpClient *http.Client
	cfg        *config.SlackConfig
	logger     *logrus.Logger
}

func NewSlack(http *http.Client, cfg *config.SlackConfig, logger *logrus.Logger) Transport {
	logger.WithFields(logrus.Fields{
		"transport": transportSlackName,
	})

	return &slack{
		httpClient: http,
		cfg:        cfg,
		logger:     logger,
	}
}

func (obj slack) Send(msg command.Result) (err error) {
	obj.logger.WithFields(
		logrus.Fields{
			"msg_type": msg.GetType(),
		},
	).Debugf("Sending '%v' message.", msg.GetName())

	switch msg.GetType() {
	case command.ImageContent:
		return obj.sendImg(msg)
	default:
		return obj.sendText(msg)
	}
}
