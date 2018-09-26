package stats

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/parsers"
	"log"
	"strconv"
	"time"
)

type loadAverage struct {
	min1            float32
	min5            float32
	min15           float32
	activeProcesses uint16
	totalProcesses  uint16
	lasPid          uint16
}

const (
	loadAvgPath          = "/proc/loadavg"
	loadAvgLoopTime      = time.Second * 5
	loadAvgKeyName       = "loadavg1min"
	loadAvgHistoryPeriod = time.Hour * 24
)

var (
	loadAvgKeysCount = int(loadAvgHistoryPeriod.Seconds() / loadAvgLoopTime.Seconds())
)

func LoadAverage(c chan<- command.Result) {
	handler := &laResultHandler{
		c,
		config.KeyValueStore,
	}
	command.ReadFileLoop(loadAvgPath, loadAvgLoopTime, handler)
}

type laResultHandler struct {
	c     chan<- command.Result
	redis StackStore
}

func (handler laResultHandler) Handle(result command.Result) {
	la := parse(result.GetText())
	err := config.KeyValueStore.Push(loadAvgKeyName, la.min1)

	if err != nil {
		errMsg := fmt.Sprint("Statistics save error: loadavg cannot write key '", loadAvgKeyName, "', error: ", err)
		log.Println(errMsg)
		handler.c <- command.NewResult("Load average statistics worker", errMsg, nil)
	}

	clearOldKeys(handler.redis)
}

func parse(result string) loadAverage {
	cols := parsers.SplitPerCols(result, " ")
	processes := parsers.SplitPerCols(cols[3], "/")

	return loadAverage{
		float32(fetchFloat(cols[0])),
		float32(fetchFloat(cols[1])),
		float32(fetchFloat(cols[2])),
		uint16(fetchInt(processes[0])),
		uint16(fetchInt(processes[1])),
		uint16(fetchInt(cols[4])),
	}
}

func fetchInt(value string) uint64 {
	res, err := strconv.ParseUint(parsers.TrimSpaces(value), 10, 0)

	if err != nil {
		log.Println("Cannot read value ", value, " as int, error: ", err)
		return 0
	}

	return res
}

func fetchFloat(value string) float64 {
	res, err := strconv.ParseFloat(parsers.TrimSpaces(value), 64)

	if err != nil {
		log.Println("Cannot read value ", value, " as float, error: ", err)
		return 0
	}

	return res
}

func clearOldKeys(stack StackStore) {
	stack.Trim(loadAvgKeyName, loadAvgKeysCount)
}
