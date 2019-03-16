package watchers

import (
	"github.com/sepuka/gowatcher/definition/logger"
	"github.com/sepuka/gowatcher/definition/watcher"
	"github.com/sepuka/gowatcher/services"
	"go.uber.org/zap"
)

type Executive interface {
	Exec()
}

var log           *zap.Logger

func RunWatchers() {
	var cnt int
	for _, wtchr := range services.GetByTag(watcher.DefWatcherTag) {
		go wtchr.(Executive).Exec()
		cnt++
	}
	services.Container.Fill(logger.DefLogger, &log)
	log.Info("Started watchers", zap.Int("number", cnt))
}
