package transports

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/definition/logger"
	"go.uber.org/zap"
	"net/http"
)

const transportSlackName = "slack"

type slack struct {
	httpClient *http.Client
	cfg        *config.SlackConfig
	logger     logger.Logger
}

func NewSlack(http *http.Client, cfg *config.SlackConfig, logger logger.Logger) Transport {
	logger.With(zap.String("transport", transportSlackName))

	return &slack{
		httpClient: http,
		cfg:        cfg,
		logger:     logger,
	}
}

func (obj slack) Send(msg command.Result) (err error) {
	obj.
		logger.
		With(
		zap.String("msg_type", string(msg.GetType())),
		zap.String("msg", msg.GetName()),
	).Debug("Sending message.")

	switch msg.GetType() {
	case command.ImageContent:
		return obj.sendImg(msg)
	default:
		return obj.sendText(msg)
	}
}
