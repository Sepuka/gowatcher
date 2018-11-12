package logger

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
	level  config.LogLevel
}

func (logger Logger) Debug(v ...interface{}) {
	logger.Log(config.LogLevelDebug, v)
}

func (logger Logger) Log(level config.LogLevel, v ...interface{}) {
	if level == logger.level {
		logger.logger.Println(v)
	}
}

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: services.LoggerComponent,
			Build: func(ctn di.Container) (interface{}, error) {
				return &Logger{log.New(os.Stderr, "", log.LstdFlags), params.Logger.Level}, nil
			},
		})
	})
}
