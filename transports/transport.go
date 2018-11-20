package transports

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"net/http"
)

const (
	contentTypeJson = "application/json"
)

type TransportName string

type Transport interface {
	GetName() TransportName
	Send(msg command.Result) (resp *http.Response, err error)
}

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: services.Transports,
			Build: func(ctn di.Container) (interface{}, error) {
				return []Transport{
					ctn.Get(services.Telegram).(*Telegram),
					ctn.Get(services.Slack).(*Slack),
				}, nil
			},
		})
	})
}
