package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	defaultConfig "github.com/hoppxi/zend/config"
	"github.com/hoppxi/zend/pkg/logger"
)


func GetConfigDir() string {
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			logger.Log.Error("Unable to detect home directory: %v", err)
			os.Exit(1)
		}
		configDir = filepath.Join(home, ".config")
	}
	zendConfigDir := filepath.Join(configDir, "zend")
	if _, err := os.Stat(zendConfigDir); os.IsNotExist(err) {
		if err := os.MkdirAll(zendConfigDir, 0755); err != nil {
			logger.Log.Error("Failed to create config directory: %v", err)
			os.Exit(1)
		}
	}
	return zendConfigDir
}

func EnsureConfig() {
	configDir := GetConfigDir()
	configFile := filepath.Join(configDir, "config.yaml")

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// Copy the full default config (with comments) to the user's config directory
		if err := os.WriteFile(configFile, defaultConfig.DefaultConfigData, 0644); err != nil {
			logger.Log.Error("Failed to create default config: %v", err)
			os.Exit(1)
		}
		logger.Log.Success("Created default config at %s", configFile)
	}

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Error("Failed to read config: %v", err)
		os.Exit(1)
	}
	logger.Log.Info("Loaded config from %s", configFile)
}

func GetDefaultConfigPath() string {
	return filepath.Join(GetConfigDir(), "config.yaml")
}
