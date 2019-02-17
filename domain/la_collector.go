package domain

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/stats"
	"time"
)

var (
	loadAvgPath          = "/proc/loadavg"
	loadAvgLoopTime      = time.Second * 5
)

type LaCollector struct {
	Handler command.ResultHandler
	Store   stats.SliceStoreWriter
}

func (obj LaCollector) Exec() {
	command.ReadFileLoop(loadAvgPath, loadAvgLoopTime, obj.Handler)
}
