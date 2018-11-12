package services

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
)

const LoggerComponent = "logger"

var (
	Container  di.Container
	components []creatorFn
)

type (
	creatorFn func(builder *di.Builder, params config.Configuration) error
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
