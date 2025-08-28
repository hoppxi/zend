package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hoppxi/zend/pkg/logger"
	"github.com/spf13/viper"
)

func RegisterRoutes(mux *http.ServeMux) {
	// Serve React frontend
	mux.Handle("/", http.FileServer(http.Dir(getReactDistPath())))
	mux.HandleFunc("/api/config", ConfigHandler)

	// Register music and images
	registerMusic(mux)
	registerImages(mux)
}

// Music handler: uses music.local.path or music.local.path_list
func registerMusic(mux *http.ServeMux) {
	if !viper.GetBool("music.enabled") {
		logger.Log.Info("Music is disabled")
		return
	}

	singlePath := viper.GetString("music.local.path")
	pathList := viper.GetStringSlice("music.local.path_list")

	if singlePath != "" {
		registerPathHandler(mux, "/api/music/", singlePath)
	} else if len(pathList) > 0 {
		validateAndRegisterPaths(mux, "/api/music/", pathList, "music")
	} else {
		logger.Log.Info("No music paths configured")
	}
}

// Image handler: uses image.path or image.path_list (no local)
func registerImages(mux *http.ServeMux) {
	if !viper.GetBool("image.enabled") {
		logger.Log.Info("Images are disabled")
		return
	}

	singlePath := viper.GetString("image.path")
	pathList := viper.GetStringSlice("image.path_list")

	if singlePath != "" {
		registerPathHandler(mux, "/api/image/", singlePath)
	} else if len(pathList) > 0 {
		validateAndRegisterPaths(mux, "/api/image/", pathList, "image")
	} else {
		logger.Log.Info("No image paths configured")
	}
}

// Validate that all paths are same type and register
func validateAndRegisterPaths(mux *http.ServeMux, route string, paths []string, key string) {
	allFiles, allDirs := true, true
	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			panic(err)
		}
		if info.IsDir() {
			allFiles = false
		} else {
			allDirs = false
		}
	}

	if !(allFiles || allDirs) {
		panic(fmt.Sprintf("%s.path_list must contain either all files or all directories", key))
	}

	for _, p := range paths {
		registerPathHandler(mux, route, p)
	}
}

func getReactDistPath() string {
	if dist := viper.GetString("dist"); dist != "" {
		return dist
	} else {
		return "string;"
	}
}


func registerPathHandler(mux *http.ServeMux, route, path string) {
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if info.IsDir() {
		mux.Handle(route, http.StripPrefix(route, http.FileServer(http.Dir(path))))
		logger.Log.Info("Serving directory %s on %s", path, route)
	} else {
		mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, path)
		})
		logger.Log.Info("Serving file %s on %s", path, route)
	}
}
