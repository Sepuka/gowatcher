package env

import (
	"github.com/sepuka/gowatcher/definition/logger"
	"github.com/sepuka/gowatcher/services"
	"github.com/sirupsen/logrus"
	"os"
)

func GetCurrentHost() string {
	host, err := os.Hostname()
	if err != nil {
		log := services.Container.Get(logger.DefLogger).(logrus.FieldLogger)
		log.Warningf("Cannot detect current host: %v", err)
		host = "unknown host"
	}

	return host
}
