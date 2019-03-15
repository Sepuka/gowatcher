package env

import (
	"github.com/sepuka/gowatcher/definition/logger"
	"github.com/sepuka/gowatcher/services"
	"go.uber.org/zap"
	"os"
)

func GetCurrentHost() string {
	host, err := os.Hostname()
	if err != nil {
		log := services.Container.Get(logger.DefLogger).(logger.Logger)
		log.Error("Cannot detect current host: %v", zap.Error(err))
		host = "unknown host"
	}

	return host
}
