package services

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
)

const (
	Logger        = "logger"
	Slack         = "transport.slack"
	Telegram      = "transport.telegram"
	Transports    = "transport.all"
	TransportChan = "transport.chan"
	KeyValue      = "store.key_value"
)

var (
	Container  di.Container
	components []creatorFn
)

type (
	creatorFn func(builder *di.Builder, cfg config.Configuration) error
)

func Register(fn creatorFn) {
	components = append(components, fn)
}

func Build(params config.Configuration) {
	builder, err := di.NewBuilder()
	if err != nil {
		panic(err)
	}

	for _, fnc := range components {
		fnc(builder, params)
	}

	Container = builder.Build()
}
