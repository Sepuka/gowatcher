package stats

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/parsers"
	"github.com/sepuka/gowatcher/services"
	"github.com/sirupsen/logrus"
	"time"
)

type loadAverageSnapshot struct {
	min1            float32
	min5            float32
	min15           float32
	activeProcesses uint16
	totalProcesses  uint16
	lasPid          uint16
}

const (
	LoadAvgKeyName       = "loadavg1min"
	loadAvgPath          = "/proc/loadavg"
	loadAvgLoopTime      = time.Second * 5
	loadAvgHistoryPeriod = time.Hour * 24
)

var (
	loadAvgKeysCount = int(loadAvgHistoryPeriod.Seconds() / loadAvgLoopTime.Seconds())
)

func LoadAverage(c chan<- command.Result, writer SliceStoreWriter) {
	handler := &laResultHandler{
		c,
		writer,
		services.Container.Get(services.Logger).(logrus.FieldLogger),
	}
	command.ReadFileLoop(loadAvgPath, loadAvgLoopTime, handler)
}

type laResultHandler struct {
	c      chan<- command.Result
	store  SliceStoreWriter
	logger logrus.FieldLogger
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

func deleteOldKeys(stack SliceStoreWriter) {
	stack.Trim(LoadAvgKeyName, loadAvgKeysCount)
}
