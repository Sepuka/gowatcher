package logger

import (
	"errors"
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	services.Register(func(builder *di.Builder, cfg config.Configuration) error {
		LogLevel, err := logrus.ParseLevel(cfg.Logger.Level)
		if err != nil {
			return errors.New("cannot parse logger level")
		}

		fileLog, err := os.OpenFile(cfg.Logger.File, os.O_WRONLY|os.O_CREATE, 0640)
		if err != nil {
			return errors.New("cannot open or create log file")
		}

		return builder.Add(di.Def{
			Name: services.Logger,
			Build: func(ctn di.Container) (interface{}, error) {
				log := &logrus.Logger{
					Formatter:    new(logrus.TextFormatter),
					Hooks:        make(logrus.LevelHooks),
					Level:        LogLevel,
					ExitFunc:     os.Exit,
					ReportCaller: cfg.Logger.WithCaller,
				}
				log.SetOutput(fileLog)

				return log, nil
			},
		})
	})
}
