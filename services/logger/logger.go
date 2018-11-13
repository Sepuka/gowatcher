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

func (obj Logger) Debugf(fmt string, v ...interface{}) {
	obj.logf(config.LogLevelDebug, fmt, v...)
}

func (obj Logger) Errorf(fmt string, v ...interface{}) {
	obj.logf(config.LogLevelDefault, fmt, v...)
}

func (obj Logger) Error(v ...interface{}) {
	obj.logger.Print(v...)
}

func (obj Logger) Fatalf(fmt string, v ...interface{}) {
	obj.logger.Printf(fmt, v...)
	os.Exit(1)
}

func (obj Logger) Log(v ...interface{}) {
	obj.logger.Print(v...)
}

func (obj Logger) logf(level config.LogLevel, fmt string, v ...interface{}) {
	if level == obj.level {
		obj.logger.Printf(fmt, v...)
	}
}

func init() {
	services.Register(func(builder *di.Builder, params config.Configuration) error {
		return builder.Add(di.Def{
			Name: services.Logger,
			Build: func(ctn di.Container) (interface{}, error) {
				return &Logger{log.New(os.Stderr, "", log.LstdFlags), params.Logger.Level}, nil
			},
		})
	})
}
