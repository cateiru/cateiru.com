package config_test

import (
	"testing"

	"github.com/cateiru/cateiru.com/src/config"
	"github.com/stretchr/testify/require"
)

func TestConfigInit(t *testing.T) {
	t.Run("created config", func(t *testing.T) {
		// Initialize config
		config.Init("test")

		// Set a config
		require.NotNil(t, config.Config)

		// If not set `mode` variable, set to set test mode in config.
		require.Equal(t, config.Config.Mode, "test")
	})

	t.Run("change mode local", func(t *testing.T) {
		config.Init("local")

		require.NotNil(t, config.Config)
		require.Equal(t, config.Config.Mode, "local")
	})
}
