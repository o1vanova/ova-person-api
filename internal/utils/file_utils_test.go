package utils

import (
	"testing"
)

func TestOpenAndCloseFileInLoop(t *testing.T) {
	configPath := GetConfigPath()

	OpenAndCloseFileInLoop(10, configPath)
}
