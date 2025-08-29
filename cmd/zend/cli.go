package zend

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/hoppxi/zend/internal/cli"
	"github.com/hoppxi/zend/internal/config"
	"github.com/hoppxi/zend/internal/server"
	"github.com/hoppxi/zend/pkg/logger"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:     "zend",
		Short:   "Zend - customizable browser homepage",
		Long:    "Zend is a lightweight customizable browser home-page server with CLI integration.",
		Version: "0.1.0",
		Run: func(cmd *cobra.Command, args []string) {
			start := time.Now() // measure startup

			addrFlag, _ := cmd.Flags().GetString("addr")
			configDir := config.GetConfigDir()
			portFile := filepath.Join(configDir, "server.port")

			if cmd.Flags().Changed("addr") && addrFlag == "" {
				if data, err := os.ReadFile(portFile); err == nil {
					logger.Log.Info("Current server address: %s", string(data))
					return
				} else {
					logger.Log.Error("No running Zend server found.")
					os.Exit(1)
				}
			}

			config.EnsureConfig()
			server.InitConfig()

			addr := server.ResolveServerAddress(addrFlag)

			if err := os.WriteFile(portFile, []byte(addr), 0644); err != nil {
				logger.Log.Error("Failed to write port file: %v", err)
				os.Exit(1)
			}

			go server.StartHugoServer()

			host, port, err := net.SplitHostPort(addr)
			if err != nil {
				log.Fatalf("Invalid address: %v", err)
			}

			// calculate elapsed ms
			elapsed := time.Since(start).Milliseconds()

			// colors
			arrow := color.New(color.FgGreen, color.Bold).Sprint("➜")
			header := color.New(color.FgHiCyan, color.Bold).SprintFunc()
			label := color.New(color.FgWhite, color.Bold).SprintFunc()
			url := color.New(color.FgHiYellow).SprintFunc()

			// formatted output
			fmt.Printf("\n%s %s\n\n", header("Using Port -"), color.HiYellowString(port))
			fmt.Printf("  %s v%s  %s\n\n", header("Zend"), cmd.Version, color.HiGreenString("started in %d ms", elapsed))
			fmt.Printf("  %s  %s:   %s\n", arrow, label("Local"), url(fmt.Sprintf("http://%s:%s/", host, port)))
			fmt.Printf("  %s  %s: use `--addr` to get the port\n", arrow, label("Network"))
			fmt.Printf("  %s  %s: use `config` to get config path\n\n", arrow, label("Config"))

			server.RegisterRoutes(http.DefaultServeMux)

			if err := http.ListenAndServe(addr, nil); err != nil {
				logger.Log.Error("Failed to start Zend API server: %v", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(cli.OpenCmd)
	rootCmd.AddCommand(cli.ConfigCmd)
	rootCmd.Flags().StringP("addr", "a", "", "Address or port to run the server on (default random)")

	if err := rootCmd.Execute(); err != nil {
		logger.Log.Error("Error: %v", err)
		os.Exit(1)
	}
}
