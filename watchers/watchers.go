package watchers

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/command/graph"
	"github.com/sepuka/gowatcher/config"
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

var (
	agents = map[string]func(chan<- command.Result, config.WatcherConfig){
		diskFreeAgentName: DiskFree,
		upTimeAgentName:   Uptime,
		whoAgentName:      Who,
		wAgentName:        W,
		laAgentName:       graph.LoadAvgGraph,
	}
)

func RunWatchers(c chan<- command.Result) {
	for _, cfg := range config.AppConfig.Watchers {
		start(c, cfg, getAgent(cfg.GetName()))
	}
}

func RunStatCollectors(c chan<- command.Result) {
	go stats.LoadAverage(c, services.Container.Get(services.KeyValue).(*store.RedisStore))
}

func start(c chan<- command.Result, config config.WatcherConfig, f func(chan<- command.Result, config.WatcherConfig)) {
	go f(c, config)
}

func getAgent(agentName string) func(chan<- command.Result, config.WatcherConfig) {
	if agent, ok := agents[agentName]; ok {
		return agent
	}
	panic(fmt.Sprint("Unknown watcher name ", agentName))
}
