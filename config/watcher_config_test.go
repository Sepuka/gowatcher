package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	baseDataProvider WatcherConfigs = []WatcherConfig{
		{"name1", 1},
		{"name2", 2},
		{"name3", 3},
	}
	tunedDataProvider WatcherConfigs = []WatcherConfig{
		{"name1", 1},
		{"name2", 3},
	}
	expectedConfigDataProvider WatcherConfigs = []WatcherConfig{
		{"name1", 1},
		{"name2", 3},
		{"name3", 3},
	}
)

func TestTuneConfig(t *testing.T) {
	for i, baseConfig := range baseDataProvider {
		expected := expectedConfigDataProvider[i]
		assert.Equal(t, expected, baseConfig.Merge(tunedDataProvider))
	}
}