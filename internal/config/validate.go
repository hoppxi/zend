package config

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hoppxi/zend/pkg/logger"
	"github.com/spf13/viper"
)

// validatePosition checks if position is a valid keyword or percentage pair.
func validatePosition(pos string) bool {
	if pos == "" {
		return true // allow empty (uses default)
	}
	validKeywords := map[string]bool{
		"top": true, "bottom": true, "left": true, "right": true,
		"top-left": true, "top-right": true, "bottom-left": true, "bottom-right": true,
		"center": true,
	}
	if validKeywords[strings.ToLower(pos)] {
		return true
	}
	// Check "(x%, y%)" format, e.g., "(50%, 30%)"
	percentRegex := regexp.MustCompile(`^\(\s*\d{1,3}%\s*,\s*\d{1,3}%\s*\)$`)
	return percentRegex.MatchString(pos)
}

// validatePath checks if a given file or directory exists (if it's not empty).
func validatePath(p string) bool {
	if p == "" {
		return true // empty is allowed in some configs
	}
	if strings.HasPrefix(p, "http://") || strings.HasPrefix(p, "https://") {
		return true // URL paths are valid
	}
	absPath, err := filepath.Abs(p)
	if err != nil {
		return false
	}
	_, err = os.Stat(absPath)
	return err == nil
}

// validateColor checks for valid hex color (#rrggbb or #rrggbbaa).
func validateColor(color string) bool {
	if color == "" {
		return true
	}
	hexRegex := regexp.MustCompile(`^#(?:[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$`)
	return hexRegex.MatchString(color) || strings.HasPrefix(color, "transparent") || strings.HasPrefix(color, "rgb")
}

// validateAPIKey ensures non-empty API key.
func validateAPIKey(key string) bool {
	return strings.TrimSpace(key) != ""
}

func Validate() {
	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Error("Config validation failed: unable to read config file: %v", err)
		os.Exit(1)
	}
	
	dist := viper.GetString("dist")
	if dist != "" && !validatePath(dist) {
		logger.Log.Error("Invalid 'dist' path: %s", dist)
		os.Exit(1)
	}
	
	randomUse := viper.GetString("random.use")
	if randomUse != "" && randomUse != "color" && randomUse != "image" {
		logger.Log.Error("Invalid 'random.use' value: %s (must be 'color' or 'image')", randomUse)
		os.Exit(1)
	}

	imagePath := viper.GetString("image.path")
	imageList := viper.GetStringSlice("image.path_list")
	if imagePath != "" && !validatePath(imagePath) {
		logger.Log.Error("Invalid image.path: %s", imagePath)
		os.Exit(1)
	}
	for _, p := range imageList {
		if !validatePath(p) {
			logger.Log.Error("Invalid image.path_list entry: %s", p)
			os.Exit(1)
		}
	}

	solidColor := viper.GetString("solid.color")
	if solidColor != "" && !validateColor(solidColor) {
		logger.Log.Error("Invalid solid.color value: %s (must be hex or named color)", solidColor)
		os.Exit(1)
	}

	transitionType := viper.GetString("general.transition.type")
	validTransitions := map[string]bool{"fade": true, "slide": true, "zoom": true, "grow": true}
	if transitionType != "" && !validTransitions[strings.ToLower(transitionType)] {
		logger.Log.Error("Invalid transition type: %s (must be fade, slide, zoom, grow)", transitionType)
		os.Exit(1)
	}

	engine := viper.GetString("search_bar.engine")
	validEngines := map[string]bool{"google": true, "bing": true, "brave": true, "duckduckgo": true}
	if engine != "" && !validEngines[strings.ToLower(engine)] {
		logger.Log.Error("Invalid search_bar.engine: %s (must be google, bing, brave, duckduckgo)", engine)
		os.Exit(1)
	}

	if !validatePosition(viper.GetString("clock.position")) {
		logger.Log.Error("Invalid clock.position: %s (must be keyword or (x%%, y%%))", viper.GetString("clock.position"))
		os.Exit(1)
	}

	if viper.GetBool("weather.enabled") {
		apiKey := viper.GetString("weather.api_key")
		if !validateAPIKey(apiKey) {
			logger.Log.Error("Missing or invalid weather.api_key (required when weather.enabled = true)")
			os.Exit(1)
		}
		if !validatePosition(viper.GetString("weather.position")) {
			logger.Log.Error("Invalid weather.position: %s", viper.GetString("weather.position"))
			os.Exit(1)
		}
	}

	if viper.GetBool("music.enabled") {
		if !validatePosition(viper.GetString("music.position")) {
			logger.Log.Error("Invalid music.position: %s", viper.GetString("music.position"))
			os.Exit(1)
		}
		if viper.GetBool("music.local.enabled") {
			musicPath := viper.GetString("music.local.path")
			if musicPath == "" && len(viper.GetStringSlice("music.local.path_list")) == 0 {
				logger.Log.Error("music.local.enabled = true but no 'path' or 'path_list' provided")
				os.Exit(1)
			}
			if musicPath != "" && !validatePath(musicPath) {
				logger.Log.Error("Invalid music.local.path: %s", musicPath)
				os.Exit(1)
			}
			for _, mp := range viper.GetStringSlice("music.local.path_list") {
				if !validatePath(mp) {
					logger.Log.Error("Invalid music.local.path_list entry: %s", mp)
					os.Exit(1)
				}
			}
		}
		if !validatePosition(viper.GetString("music.visualizer.position")) {
			logger.Log.Error("Invalid music.visualizer.position: %s", viper.GetString("music.visualizer.position"))
			os.Exit(1)
		}
	}

	logger.Log.Success("Config is valid!")
}
