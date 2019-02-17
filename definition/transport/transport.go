package transport

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
)

const (
	DefTransportTag = "transports"
	DefTransportChan = "definition.transport.chan"
)

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefTransportChan,
			Build: func(ctn di.Container) (interface{}, error) {
				return make(chan command.Result), nil
			},
		})
	})
}
