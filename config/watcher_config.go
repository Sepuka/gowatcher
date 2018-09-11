package config

import (
	"time"
)

type WatcherConfigs []WatcherConfig

func (ws WatcherConfigs) Tune(defaultConfig WatcherConfig) WatcherConfig {
	for _, s := range ws {
		if s.name == defaultConfig.GetName() {
			return WatcherConfig{s.name, defaultConfig.GetLoop()}
		}
	}
	return defaultConfig
}

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
	return time.Duration(setting.loop)
}

func (baseConfig WatcherConfig) Merge(tunedConfig WatcherConfigs) WatcherConfig {
	for _, tuned := range tunedConfig {
		if tuned.GetName() == baseConfig.GetName() {
			// не просто возвращать, а делать мерж того, что отличается
			return tuned
		}
	}
	return baseConfig
}