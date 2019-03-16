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
		var (
			logger Logger
			zapCfg zap.Config
		)

		if cfg.Logger.IsProduction {
			zapCfg = zap.NewProductionConfig()
		} else {
			zapCfg = zap.NewDevelopmentConfig()

		}

		zapCfg.OutputPaths = append(zapCfg.OutputPaths, cfg.Logger.File)
		logger, err = zapCfg.Build()
		if err != nil {
			return err
		}
		host, _ := os.Hostname()
		logger.With(zap.String("host", host))

		return builder.Add(di.Def{
			Name: DefLogger,
			Build: func(ctn di.Container) (interface{}, error) {
				return logger, nil
			},
			Close: func(obj interface{}) error {
				return obj.(Logger).Sync()
			},
		})
	})
}
