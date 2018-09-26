package config

import (
	"time"
)

const (
	watcherNameSection = "name"
	watcherLoopSection = "loop"
	minLoop            = time.Second
	maxLoop            = time.Hour * 24
)

func NewWatcherConfig(name string, loop time.Duration) *WatcherConfig {
	return &WatcherConfig{name, loop}
}

type WatcherConfig struct {
	name string
	loop time.Duration
}

func (setting WatcherConfig) GetName() string {
	return setting.name
}

func (setting WatcherConfig) GetLoop() time.Duration {
	return setting.loop
}

func (baseConfig WatcherConfig) Merge(tunedConfig []WatcherConfig) WatcherConfig {
	for _, tuned := range tunedConfig {
		if tuned.GetName() == baseConfig.GetName() {
			// не просто возвращать, а делать мерж того, что отличается
			return tuned
		}
	}
	return baseConfig
}

func initWatcherConfigs() {
	for _, watcher := range config.Watchers {
		cfg := WatcherConfig{
			watcher[watcherNameSection].(string),
			toValidTime(watcher[watcherLoopSection].(float64)),
		}
		WatchersConfig = append(WatchersConfig, cfg)
	}
}

func toValidTime(loopTime float64) time.Duration {
	cfgLoopTime := time.Duration(loopTime) * time.Second

	if cfgLoopTime < minLoop {
		return minLoop
	}
	if cfgLoopTime > maxLoop {
		return maxLoop
	}

	return cfgLoopTime
}
