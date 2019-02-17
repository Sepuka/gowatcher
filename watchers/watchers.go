package watchers

import (
	"github.com/sepuka/gowatcher/definition/watcher"
	"github.com/sepuka/gowatcher/services"
)

type Executive interface {
	Exec()
}

func RunWatchers() {
	for _, wtchr := range services.GetByTag(watcher.DefWatcherTag) {
		go wtchr.(Executive).Exec()
	}
}
