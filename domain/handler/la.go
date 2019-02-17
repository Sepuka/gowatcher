package handler

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/parsers"
	"github.com/sepuka/gowatcher/stats"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	LoadAvgKeyName       = "loadavg1min"
	loadAvgLoopTime      = time.Second * 5
	loadAvgHistoryPeriod = time.Hour * 24
)

var (
	loadAvgKeysCount = int(loadAvgHistoryPeriod.Seconds() / loadAvgLoopTime.Seconds())
)

type loadAverageSnapshot struct {
	min1            float32
	min5            float32
	min15           float32
	activeProcesses uint16
	totalProcesses  uint16
	lasPid          uint16
}

type laResultHandler struct {
	c      chan<-command.Result
	store  stats.SliceStoreWriter
	logger logrus.FieldLogger
}

func NewLaResultHandler(transportChan chan<-command.Result, redis stats.SliceStoreWriter, logger *logrus.Logger) command.ResultHandler {
	return &laResultHandler{
		c: transportChan,
		store: redis,
		logger: logger,
	}
}

func (handler laResultHandler) Handle(result command.Result) {
	la := parse(result.GetContent())
	err := handler.store.Push(LoadAvgKeyName, la.min1)

	if err != nil {
		errMsg := fmt.Sprint("Statistics save error: loadavg cannot write key '", LoadAvgKeyName, "', error: ", err)
		handler.logger.Error(errMsg)
		handler.c <- command.NewResult("Load average statistics worker", errMsg, nil)
	}

	deleteOldKeys(handler.store)
}

func parse(result string) loadAverageSnapshot {
	cols := parsers.SplitPerCols(result, " ")
	processes := parsers.SplitPerCols(cols[3], "/")

	return loadAverageSnapshot{
		float32(parsers.FetchFloat(cols[0])),
		float32(parsers.FetchFloat(cols[1])),
		float32(parsers.FetchFloat(cols[2])),
		uint16(parsers.FetchInt(processes[0])),
		uint16(parsers.FetchInt(processes[1])),
		uint16(parsers.FetchInt(cols[4])),
	}
}

func deleteOldKeys(stack stats.SliceStoreWriter) {
	stack.Trim(LoadAvgKeyName, loadAvgKeysCount)
}
