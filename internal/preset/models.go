package preset

// Config holds all presets
type Config struct {
	Presets map[string]Preset `yaml:"presets"`
}

// Preset defines a group of services
type Preset struct {
	Services    []string          `yaml:"services"`
	Env         map[string]string `yaml:"env,omitempty"`         // Optional environment variables
	Description string            `yaml:"description,omitempty"` // Optional description
}
