package utils

import (
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func TestOpenAndCloseFileInLoop(t *testing.T) {
	configPath := getConfigPath()

	OpenAndCloseFileInLoop(10, configPath)
}

func getConfigPath() string {
	_, currentFile, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(currentFile), "../..")
	configDir := path.Join(rootDir, "configs")
	configPath := filepath.Join(configDir, "config.json")

	return configPath
}
