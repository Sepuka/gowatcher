package watchers

import (
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"fmt"
)

const (
	DiskFreeAgent = "DiskFree"
	UpTimeAgent   = "Uptime"
	WhoAgent      = "Who"
	WAgent        = "W"
)

var (
	baseConfigs config.WatcherConfigs = []config.WatcherConfig{
		*config.NewWatcherConfig(DiskFreeAgent, DfLoopInterval),
		*config.NewWatcherConfig(UpTimeAgent, UptimeLoopInterval),
		*config.NewWatcherConfig(WhoAgent, WhoLoopInterval),
		*config.NewWatcherConfig(WAgent, WLoopInterval),
	}
	agents = map[string]func(chan<- command.Result, config.WatcherConfig){
		DiskFreeAgent: DiskFree,
		UpTimeAgent: Uptime,
		WhoAgent: Who,
		WAgent: W,
	}
)

func RunWatchers(c chan<- command.Result) {
	for _, st := range baseConfigs {
		cfg := st.Merge(config.WatchersConfig)
		run(c, cfg, getAgent(st.GetName()))
	}
}

func run(c chan<- command.Result, config config.WatcherConfig, f func(chan<- command.Result, config.WatcherConfig)) {
	go f(c, config)
}

func getAgent(agentName string) func(chan<- command.Result, config.WatcherConfig) {
	if agent, ok := agents[agentName]; ok {
		return agent
	}
	panic(fmt.Sprint("Unknown watcher name ", agentName))
}