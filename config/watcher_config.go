package config

import (
	"time"
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