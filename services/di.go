package services

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
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

func GetByTag(tag string) []interface{} {
	var defs []interface{}

	for _, def := range Container.Definitions() {
		for _, defTag := range def.Tags {
			if defTag.Name == tag {
				var fff interface{}
				if err := Container.Fill(def.Name, &fff); err != nil {
					panic(err)
				}
				defs = append(defs, fff)
			}
		}
	}

	return defs
}
