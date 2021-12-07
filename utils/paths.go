package utils

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// DestinationExists checks if a path or file exists
func DestinationExists(path string) bool {
	if _, err := os.Open(path); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false
		}
	}

	return true
}

// MakeDestination creates directory if it doesn't exist
func MakeDestination(path string) {
	dir := filepath.Dir(path)
	if !DestinationExists(dir) {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			Error.Fatal(err)
		}
	}
}

// FileExists checks if a file exists and returns fileinfo
func FileExists(file string) (fi fs.FileInfo, err error) {
	if fi, err = os.Stat(file); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return
}

// RootDir returns the application root directory
func RootDir() (dir string) {
	if dir = viper.GetString("defaults.workdir"); dir == "" {
		return "."
	}

	return viper.GetString("defaults.workdir")
}
