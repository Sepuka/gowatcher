package watchers

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/command/graph"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/services/store"
	"github.com/sepuka/gowatcher/stats"
	"github.com/sirupsen/logrus"
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
		fnc, err := getAgent(cfg.GetName())

		if err != nil {
			log := services.Container.Get(services.Logger).(logrus.FieldLogger)
			log.Error(err)
			continue
		}

		go fnc(c, cfg)
	}
}

func RunStatCollectors(c chan<- command.Result) {
	go stats.LoadAverage(c, services.Container.Get(services.KeyValue).(*store.RedisStore))
}

func getAgent(agentName string) (func(chan<- command.Result, config.WatcherConfig), error) {
	if agent, ok := agents[agentName]; ok {
		return agent, nil
	}

	return nil, fmt.Errorf("invalid agent name %s", agentName)
}
