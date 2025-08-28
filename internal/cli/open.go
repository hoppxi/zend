package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/hoppxi/zend/pkg/logger"
)

var overrideJSON string

// `zend open`
var OpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Open running Zend server in browser with temporary config overrides",
	Long:  "Opens the Zend homepage in the default browser and applies temporary config overrides from JSON.",
	Run:   runOpenCommand,
}

func init() {
	OpenCmd.Flags().StringVarP(&overrideJSON, "override", "o", "", "JSON object with temporary config overrides")
}

func runOpenCommand(cmd *cobra.Command, args []string) {
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		home, _ := os.UserHomeDir()
		configDir = filepath.Join(home, ".config")
	}
	portFile := filepath.Join(configDir, "zend", "server.port")

	data, err := os.ReadFile(portFile)
	if err != nil {
		logger.Log.Error("No running Zend server found.")
		os.Exit(1)
	}
	url := fmt.Sprintf("http://%s", string(data))

	if overrideJSON != "" {
		var tempConfig map[string]interface{}
		if err := json.Unmarshal([]byte(overrideJSON), &tempConfig); err != nil {
			logger.Log.Error("Invalid JSON for --override: %v", err)
			os.Exit(1)
		}

		configBytes, _ := json.Marshal(tempConfig)
		resp, err := http.Post(url+"/api/config", "application/json", bytes.NewBuffer(configBytes))
		if err != nil || resp.StatusCode != 200 {
			logger.Log.Error("Failed to apply overrides: %v", err)
		} else {
			logger.Log.Info("Applied temporary config overrides")
		}
	}

	if err := openBrowser(url); err != nil {
		logger.Log.Error("Failed to open browser: %v", err)
		os.Exit(1)
	}
	logger.Log.Success("Zend homepage opened in browser: %s", url)
}

// opens default browser
func openBrowser(url string) error {
	var cmdName string
	var args []string
	switch runtime.GOOS {
	case "darwin":
		cmdName = "open"
		args = []string{url}
	case "windows":
		cmdName = "cmd"
		args = []string{"/c", "start", url}
	default:
		cmdName = "xdg-open"
		args = []string{url}
	}
	return exec.Command(cmdName, args...).Start()
}
