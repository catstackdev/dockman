package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const ProjectConfigName = ".dockman.yml"

// ProjectConfig holds project-specific configuration
type ProjectConfig struct {
	ComposeFile   string            `yaml:"compose_file,omitempty"`
	DefaultPreset string            `yaml:"default_preset,omitempty"`
	Aliases       map[string]string `yaml:"aliases,omitempty"`
	AutoPull      bool              `yaml:"auto_pull,omitempty"`
}

// LoadProjectConfig loads config from project directory
func LoadProjectConfig(projectDir string) (*ProjectConfig, error) {
	configPath := filepath.Join(projectDir, ProjectConfigName)

	// Check if config exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Return default config if not found
		return &ProjectConfig{}, nil
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	// Parse YAML
	var config ProjectConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &config, nil
}

// SaveProjectConfig saves config to project directory
func SaveProjectConfig(projectDir string, config *ProjectConfig) error {
	configPath := filepath.Join(projectDir, ProjectConfigName)

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0o644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}
