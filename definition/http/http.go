package http

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"net/http"
)

const DefHttpClient = "definition.http"

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefHttpClient,
			Build: func(ctn di.Container) (interface{}, error) {
				return &http.Client{}, nil
			},
		})
	})
}
