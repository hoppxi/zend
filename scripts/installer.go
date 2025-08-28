package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("=== Zend Cross-Platform Installer ===")

	// 1. Detect OS
	osName := runtime.GOOS
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// 2. Determine config and frontend paths
	var configDir string
	if osName == "windows" {
		appData := os.Getenv("APPDATA")
		if appData == "" {
			appData = filepath.Join(homeDir, "AppData", "Roaming")
		}
		configDir = filepath.Join(appData, "Zend")
	} else {
		xdg := os.Getenv("XDG_CONFIG_HOME")
		if xdg == "" {
			xdg = filepath.Join(homeDir, ".config")
		}
		configDir = filepath.Join(xdg, "zend")
	}

	frontendDist := filepath.Join(".", "web", "dist")
	zendBin := filepath.Join(".", "zend")

	// 3. Check Go and Node/npm
	checkCommand("go", "--version")
	checkCommand("node", "--version")
	checkCommand("npm", "--version")

	// 4. Build CLI
	fmt.Println("Building CLI...")
	runCommand("go", "mod", "tidy")
	runCommand("go", "build", "-o", zendBin)

	// 5. Build frontend
	fmt.Println("Building frontend...")
	runCommand("npm", "install", "--prefix", "web")
	runCommand("npm", "run", "build", "--prefix", "web")

	// 6. Create config if missing
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		fmt.Printf("Creating config directory: %s\n", configDir)
		err := os.MkdirAll(configDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	configFile := filepath.Join(configDir, "config.yaml")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Printf("Copying default config to: %s\n", configFile)
		runCommand("cp", "config/default.yaml", configFile)
	}

	// 7. Set ZEND_DIST environment variable for current session
	setZendDist(frontendDist, osName)

	fmt.Println("=== Installation completed ===")
	fmt.Printf("Run '%s' to start the server\n", zendBin)
	fmt.Println("Frontend assets are ready and $ZEND_DIST is set for this session.")
}

func checkCommand(cmd string, args ...string) {
	_, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Printf("Error: %s not found. Please install it.\n", cmd)
		os.Exit(1)
	}
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running command: %s %v\n", name, args)
		panic(err)
	}
}

func setZendDist(path, osName string) {
	fmt.Println("Setting ZEND_DIST environment variable...")

	switch osName {
	case "windows":
		fmt.Println("Use the following command in PowerShell:")
		fmt.Printf("  $env:ZEND_DIST='%s'\n", path)
	case "darwin", "linux":
		fmt.Println("You can export ZEND_DIST for this session with:")
		fmt.Printf("  export ZEND_DIST=%s\n", path)
	default:
		fmt.Printf("Set ZEND_DIST manually to: %s\n", path)
	}
}
