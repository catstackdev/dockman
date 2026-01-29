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

// ResolveAlias checks if the first argument is an alias and expands it
// Returns the resolved args, whether it was an alias, and any error
func ResolveAlias(projectDir string, args []string) ([]string, bool, error) {
	if len(args) == 0 {
		return args, false, nil
	}

	config, err := LoadProjectConfig(projectDir)
	if err != nil {
		return nil, false, err
	}

	// Check if first arg is an alias
	if config.Aliases == nil {
		return args, false, nil
	}

	aliasCmd, exists := config.Aliases[args[0]]
	if !exists {
		return args, false, nil
	}

	// Split alias command into parts
	aliasParts := splitArgs(aliasCmd)

	// Append any extra args passed after the alias
	if len(args) > 1 {
		aliasParts = append(aliasParts, args[1:]...)
	}

	return aliasParts, true, nil
}

// splitArgs splits a command string into arguments
func splitArgs(cmd string) []string {
	var args []string
	var current string
	inQuote := false
	quoteChar := rune(0)

	for _, r := range cmd {
		switch {
		case r == '"' || r == '\'':
			if inQuote && r == quoteChar {
				inQuote = false
				quoteChar = 0
			} else if !inQuote {
				inQuote = true
				quoteChar = r
			} else {
				current += string(r)
			}
		case r == ' ' && !inQuote:
			if current != "" {
				args = append(args, current)
				current = ""
			}
		default:
			current += string(r)
		}
	}

	if current != "" {
		args = append(args, current)
	}

	return args
}
