package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hoppxi/zend/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func HandleGet(cmd *cobra.Command) bool {
	getKey, _ := cmd.Flags().GetString("get")
	if getKey == "" {
		return false
	}

	val := viper.Get(getKey)
	if val == nil {
		logger.Log.Error("Key '%s' not found", getKey)
		os.Exit(1)
	}

	fmt.Printf("%v\n", val)
	return true
}

func HandleSet(cmd *cobra.Command) bool {
	setKV, _ := cmd.Flags().GetString("set")
	if setKV == "" {
		return false
	}

	parts := strings.SplitN(setKV, "=", 2)
	if len(parts) != 2 {
		logger.Log.Error("Invalid format for --set. Use key=value")
		os.Exit(1)
	}

	key, value := parts[0], parts[1]
	viper.Set(key, value)

	if err := viper.WriteConfig(); err != nil {
		logger.Log.Error("Failed to write config: %v", err)
		os.Exit(1)
	}

	logger.Log.Success("Updated %s in config", key)
	return true
}

func HandleApply(cmd *cobra.Command, sourcePath string) bool {
	applyFlag, _ := cmd.Flags().GetBool("apply")
	if !applyFlag || sourcePath == "" {
		return false
	}

	defaultPath := GetDefaultConfigPath()
	input, err := os.ReadFile(sourcePath)
	if err != nil {
		logger.Log.Error("Failed to read source config: %v", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(filepath.Dir(defaultPath), 0755); err != nil {
		logger.Log.Error("Failed to create default config directory: %v", err)
		os.Exit(1)
	}

	if err := os.WriteFile(defaultPath, input, 0644); err != nil {
		logger.Log.Error("Failed to apply config: %v", err)
		os.Exit(1)
	}

	logger.Log.Success("Applied %s as the default config", sourcePath)
	return true
}

func HandleValidate(cmd *cobra.Command) bool {
	validateFlag, _ := cmd.Flags().GetBool("validate")
	if !validateFlag {
		return false
	}
	Validate()
	return true
}

func OutputConfigJSON() {
	allSettings := viper.AllSettings()
	jsonData, err := json.MarshalIndent(allSettings, "", "  ")
	if err != nil {
		logger.Log.Error("Failed to encode config as JSON: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
