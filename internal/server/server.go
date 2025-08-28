package server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/hoppxi/zend/pkg/logger"
	"github.com/spf13/viper"
)

var (
	runtimeCfg = map[string]any{}
)

func InitConfig() {
	runtimeCfg = viper.AllSettings()
}

func ResolveServerAddress(flagAddr string) string {
	addr := flagAddr
	if addr == "" {
		addr = viper.GetString("server.addr")
	}

	if addr != "" {
		if _, _, err := net.SplitHostPort(addr); err != nil {
			if p, err2 := strconv.Atoi(addr); err2 == nil && p > 0 && p < 65536 {
				return fmt.Sprintf("127.0.0.1:%d", p)
			}
			logger.Log.Error("Invalid port: %s", addr)
			os.Exit(1)
		}
		return addr
	}

	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		logger.Log.Error("Failed to pick random port: %v", err)
		os.Exit(1)
	}
	defer l.Close()
	return l.Addr().String()
}

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		jsonEncode(runtimeCfg, w)
	case http.MethodPost:
		var incoming map[string]any
		if err := decodeJSONBody(r, &incoming); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// reset runtime config before merging
		runtimeCfg = viper.AllSettings()
		merge(runtimeCfg, incoming)
		jsonEncode(runtimeCfg, w)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func decodeJSONBody(r *http.Request, out any) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(out)
}

func jsonEncode(data any, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(data)
}

func merge(dst, src map[string]interface{}) {
	for k, v := range src {
		if vmap, ok := v.(map[string]interface{}); ok {
			if dmap, ok := dst[k].(map[string]interface{}); ok {
				merge(dmap, vmap)
				continue
			}
		}
		dst[k] = v
	}
}
