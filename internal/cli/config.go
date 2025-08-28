package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hoppxi/zend/internal/config"
	"github.com/hoppxi/zend/pkg/logger"
)

//`zend config`
var ConfigCmd = &cobra.Command{
	Use:   "config [file]",
	Short: "Manage Zend configuration",
	Long:  "View, modify, validate, or apply Zend YAML configuration files",
	Args:  cobra.MaximumNArgs(1),
	Run:   runConfigCommand,
}

func init() {
	ConfigCmd.Flags().BoolP("print", "p", false, "Print current/default config in JSON")
	ConfigCmd.Flags().BoolP("validate", "v", false, "Validate config file syntax")
	ConfigCmd.Flags().StringP("get", "g", "", "Get a config key value (dot notation supported)")
	ConfigCmd.Flags().StringP("set", "s", "", "Set a config key=value (dot notation supported)")
	ConfigCmd.Flags().BoolP("apply", "a", false, "Apply a config file as the default (not for Home Manager)")
}

func runConfigCommand(cmd *cobra.Command, args []string) {
	filePath := ""

	// Determine config file path
	if len(args) == 1 {
		filePath = args[0]
		viper.SetConfigFile(filePath)
	} else {
		filePath = config.GetDefaultConfigPath()
		viper.SetConfigFile(filePath)
	}

	viper.SetConfigType("yaml")

	// Ensure the config file exists for --set or --apply
	setKV, _ := cmd.Flags().GetString("set")
	applyFlag, _ := cmd.Flags().GetBool("apply")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if setKV != "" || applyFlag {
			if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
				logger.Log.Error("Failed to create config directory: %v", err)
				os.Exit(1)
			}
			if err := os.WriteFile(filePath, []byte{}, 0644); err != nil {
				logger.Log.Error("Failed to create config file: %v", err)
				os.Exit(1)
			}
		} else {
			logger.Log.Error("Config file not found: %s", filePath)
			os.Exit(1)
		}
	}

	// Read existing config (ignore empty file errors if --set)
	_ = viper.ReadInConfig()

	// Process flags in order of priority
	if config.HandleSet(cmd) { return }
	if config.HandleApply(cmd, filePath) { return }
	if config.HandleGet(cmd) { return }
	if config.HandleValidate(cmd) { return }

	// --print outputs clean JSON
	printFlag, _ := cmd.Flags().GetBool("print")
	if printFlag {
		config.OutputConfigJSON()
		return
	}

	// Default: just print the config path
	fmt.Println(viper.ConfigFileUsed())
}
