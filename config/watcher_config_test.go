package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	baseDataProvider = []WatcherConfig{
		{"name1", 1},
		{"name2", 2},
		{"name3", 3},
	}
	tunedDataProvider = []WatcherConfig{
		{"name1", 1},
		{"name2", 3},
	}
	expectedConfigDataProvider = []WatcherConfig{
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