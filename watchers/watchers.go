package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"fmt"
)

const (
	diskFreeAgentName = "DiskFree"
	upTimeAgentName   = "Uptime"
	whoAgentName      = "Who"
	wAgentName        = "W"
)

var (
	baseConfigs = []config.WatcherConfig{
		*config.NewWatcherConfig(diskFreeAgentName, DfLoopInterval),
		*config.NewWatcherConfig(upTimeAgentName, UptimeLoopInterval),
		*config.NewWatcherConfig(whoAgentName, WhoLoopInterval),
		*config.NewWatcherConfig(wAgentName, WLoopInterval),
	}
	agents = map[string]func(chan<- command.Result, config.WatcherConfig){
		diskFreeAgentName: DiskFree,
		upTimeAgentName:   Uptime,
		whoAgentName:      Who,
		wAgentName:        W,
	}
)

func RunWatchers(c chan<- command.Result) {
	for _, baseConfig := range baseConfigs {
		preparedConfig := baseConfig.Merge(config.WatchersConfig)
		start(c, preparedConfig, getAgent(baseConfig.GetName()))
	}
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