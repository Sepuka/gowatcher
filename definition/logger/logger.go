package logger

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"go.uber.org/zap"
	"os"
)

const (
	DefLogger = "definition.logger"
)

type Logger = *zap.Logger

func init() {
	services.Register(func(builder *di.Builder, cfg config.Configuration) (err error) {
		var logger Logger
		if cfg.Logger.IsProduction {
			logger, err = zap.NewProduction()
			if err != nil {
				return err
			}
		} else {
			logger, err = zap.NewDevelopment()
			if err != nil {
				return err
			}
		}
		defer logger.Sync()

		host, _ := os.Hostname()
		logger.With(zap.String("host", host))

		return builder.Add(di.Def{
			Name: DefLogger,
			Build: func(ctn di.Container) (interface{}, error) {
				return logger, nil
			},
		})
	})
}
