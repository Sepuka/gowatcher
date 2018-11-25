package watchers

import (
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/services/store"
	"github.com/sepuka/gowatcher/stats"
)

const (
	diskFreeAgentName = "DiskFree"
	upTimeAgentName   = "Uptime"
	whoAgentName      = "Who"
	wAgentName        = "W"
	laAgentName       = "LoadAvgGraph"
)

type executive interface {
	exec()
}

var (
	workers = map[string]executive{
		diskFreeCommand: df,
		whoCommand:      who,
		wCommand:        w,
		uptimeCommand:   ut,
		laAgentName:     la,
	}
)

func RunWatchers() {
	for _, watcher := range workers {
		go watcher.exec()
	}
}

func RunStatCollectors() {
	go stats.LoadAverage(services.Container.Get(services.KeyValue).(*store.RedisStore))
}
