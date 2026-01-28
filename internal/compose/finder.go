package compose

import (
	"fmt"
	"os"
	"path/filepath"
)

var composeFileNames = []string{
	"docker-compose.yml",
	"docker-compose.yaml",
	"compose.yml",
	"compose.yaml",
}

// FindComposeFile searches for docker-compose file
// 1. Current directory
// 2. Walk up parent directories
// 3. Return error if not found
func FindComposeFile() (string, error) {
	// Start from current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	// Search current and parent directories
	dir := currentDir
	for {
		// Check all possible compose file names
		for _, filename := range composeFileNames {
			path := filepath.Join(dir, filename)
			if fileExists(path) {
				return path, nil
			}
		}

		// Move to parent directory
		parent := filepath.Dir(dir)

		// Stop if we reached root
		if parent == dir {
			break
		}

		dir = parent
	}

	return "", fmt.Errorf("docker-compose file not found (searched from %s to root)", currentDir)
}

// fileExists checks if a file exists
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// GetProjectDir returns the directory containing docker-compose file
func GetProjectDir(composePath string) string {
	return filepath.Dir(composePath)
}
