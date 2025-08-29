package server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hoppxi/zend/internal/config"
	"github.com/hoppxi/zend/pkg/logger"
)

const (
	HugoFixedPort = 52125
	HugoDist      = "docs/public"
)

func StartHugoServer() {
	mux := http.NewServeMux()

	if _, err := os.Stat(HugoDist); err == nil {
		mux.Handle("/", http.FileServer(http.Dir(HugoDist)))
		logger.Log.Info("Hugo docs server serving at fixed port %d", HugoFixedPort)
	} else {
		logger.Log.Warn("Hugo build folder not found: %s", HugoDist)
	}

	mux.HandleFunc("/api/port", func(w http.ResponseWriter, r *http.Request) {
		configDir := config.GetConfigDir()
		portFile := filepath.Join(configDir, "server.port")

		data, err := os.ReadFile(portFile)
		if err != nil {
			http.Error(w, "server.port file not found", http.StatusInternalServerError)
			return
		}

		full := string(data)

		host, portStr, err := net.SplitHostPort(full)
		if err != nil {
			http.Error(w, "invalid server.port content", http.StatusInternalServerError)
			return
		}

		response := map[string]any{
			"port": portStr,
			"base": host,
			"full": full,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})


	addr := fmt.Sprintf("127.0.0.1:%d", HugoFixedPort)
	logger.Log.Info("Serving Hugo docs at http://%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		logger.Log.Error("Failed to start Hugo server: %v", err)
	}
}
