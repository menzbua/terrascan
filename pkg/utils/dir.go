package utils

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"go.uber.org/zap"
)

// GetHomeDir returns the home directory path
func GetHomeDir() (terrascanDir string) {
	zap.S().Debug("looking up for the home directory path")

	terrascanDir, err := homedir.Dir()

	if err != nil {
		zap.S().Warnf("unable to determine the home directory: %v\n", err)
	}

	return
}

// GenerateTempDir generates a temporary directory
func GenerateTempDir() string {
	return filepath.Join(os.TempDir(), GenRandomString(6))
}
