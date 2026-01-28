package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigDirName   = ".dockman"
	PresetsFileName = "presets.yaml"
)

// GetConfigDir returns the config directory path (~/.dockman)
func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, ConfigDirName), nil
}

// GetPresetsPath returns the presets file path
func GetPresetsPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, PresetsFileName), nil
}

// EnsureConfigDir creates the config directory if it doesn't exist
func EnsureConfigDir() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configDir, 0o755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	return nil
}

// PresetsFileExists checks if presets file exists
func PresetsFileExists() (bool, error) {
	path, err := GetPresetsPath()
	if err != nil {
		return false, err
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
