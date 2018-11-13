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

		return builder.Add(di.Def{
			Name: services.Logger,
			Build: func(ctn di.Container) (interface{}, error) {
				return &logrus.Logger{
					Out:          os.Stderr,
					Formatter:    new(logrus.TextFormatter),
					Hooks:        make(logrus.LevelHooks),
					Level:        LogLevel,
					ExitFunc:     os.Exit,
					ReportCaller: false,
				}, nil
			},
		})
	})
}
