package preset

import (
	"os"
	"testing"

	"github.com/catstackdev/dockman/internal/config"
)

func TestPresetManager(t *testing.T) {
	// Setup temp config directory
	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)
	defer os.Unsetenv("HOME")

	// Ensure config dir exists
	config.EnsureConfigDir()

	// Create manager
	mgr, err := NewManager()
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	// Test default presets were created
	presets := mgr.GetAll()
	if len(presets) == 0 {
		t.Error("Expected default presets to be created")
	}

	// Test Get
	devPreset, err := mgr.Get("dev")
	if err != nil {
		t.Errorf("Failed to get 'dev' preset: %v", err)
	}
	if len(devPreset.Services) == 0 {
		t.Error("Dev preset should have services")
	}

	// Test Exists
	if !mgr.Exists("dev") {
		t.Error("Expected 'dev' preset to exist")
	}
	if mgr.Exists("nonexistent") {
		t.Error("Expected 'nonexistent' preset to not exist")
	}

	// Test List
	names := mgr.List()
	if len(names) == 0 {
		t.Error("Expected preset names to be returned")
	}
}
