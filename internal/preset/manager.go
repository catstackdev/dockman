package preset

import (
	"fmt"
	"os"

	"github.com/catstackdev/dockman/internal/config"
	"gopkg.in/yaml.v3"
)

// Manager handles preset operations
type Manager struct {
	config Config
}

// NewManager creates a new preset manager
func NewManager() (*Manager, error) {
	m := &Manager{}
	if err := m.load(); err != nil {
		return nil, err
	}
	return m, nil
}

// load reads presets from config file
func (m *Manager) load() error {
	// Check if presets file exists
	exists, err := config.PresetsFileExists()
	if err != nil {
		return err
	}

	if !exists {
		// Create default presets
		return m.createDefaultPresets()
	}

	// Read existing presets
	path, err := config.GetPresetsPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read presets: %w", err)
	}

	if err := yaml.Unmarshal(data, &m.config); err != nil {
		return fmt.Errorf("failed to parse presets: %w", err)
	}

	return nil
}

// createDefaultPresets creates a default presets file
func (m *Manager) createDefaultPresets() error {
	// Ensure config directory exists
	if err := config.EnsureConfigDir(); err != nil {
		return err
	}

	// Default presets
	m.config = Config{
		Presets: map[string]Preset{
			"dev": {
				Services:    []string{"postgres", "redis", "api"},
				Description: "Development environment",
			},
			"db": {
				Services:    []string{"postgres", "redis"},
				Description: "Database services only",
			},
		},
	}

	return m.save()
}

// save writes presets to config file
func (m *Manager) save() error {
	path, err := config.GetPresetsPath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(&m.config)
	if err != nil {
		return fmt.Errorf("failed to marshal presets: %w", err)
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("failed to write presets: %w", err)
	}

	return nil
}

// Get returns a preset by name
func (m *Manager) Get(name string) (*Preset, error) {
	preset, exists := m.config.Presets[name]
	if !exists {
		return nil, fmt.Errorf("preset '%s' not found", name)
	}
	return &preset, nil
}

// List returns all preset names
func (m *Manager) List() []string {
	names := make([]string, 0, len(m.config.Presets))
	for name := range m.config.Presets {
		names = append(names, name)
	}
	return names
}

// Exists checks if a preset exists
func (m *Manager) Exists(name string) bool {
	_, exists := m.config.Presets[name]
	return exists
}

// GetAll returns all presets
func (m *Manager) GetAll() map[string]Preset {
	return m.config.Presets
}
