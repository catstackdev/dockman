package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetConfigDir(t *testing.T) {
	// Save original HOME
	originalHome := os.Getenv("HOME")
	defer func() { _ = os.Setenv("HOME", originalHome) }()

	tmpDir := t.TempDir()
	if err := os.Setenv("HOME", tmpDir); err != nil {
		t.Fatalf("Failed to set HOME: %v", err)
	}

	configDir, err := GetConfigDir()
	if err != nil {
		t.Fatalf("GetConfigDir failed: %v", err)
	}

	expected := filepath.Join(tmpDir, ConfigDirName)
	if configDir != expected {
		t.Errorf("Expected %s, got %s", expected, configDir)
	}
}

func TestGetPresetsPath(t *testing.T) {
	originalHome := os.Getenv("HOME")
	defer func() { _ = os.Setenv("HOME", originalHome) }()

	tmpDir := t.TempDir()
	if err := os.Setenv("HOME", tmpDir); err != nil {
		t.Fatalf("Failed to set HOME: %v", err)
	}

	presetsPath, err := GetPresetsPath()
	if err != nil {
		t.Fatalf("GetPresetsPath failed: %v", err)
	}

	expected := filepath.Join(tmpDir, ConfigDirName, PresetsFileName)
	if presetsPath != expected {
		t.Errorf("Expected %s, got %s", expected, presetsPath)
	}
}

func TestEnsureConfigDir(t *testing.T) {
	originalHome := os.Getenv("HOME")
	defer func() { _ = os.Setenv("HOME", originalHome) }()

	tmpDir := t.TempDir()
	if err := os.Setenv("HOME", tmpDir); err != nil {
		t.Fatalf("Failed to set HOME: %v", err)
	}

	err := EnsureConfigDir()
	if err != nil {
		t.Fatalf("EnsureConfigDir failed: %v", err)
	}

	configDir := filepath.Join(tmpDir, ConfigDirName)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		t.Error("Config directory was not created")
	}

	// Test idempotency - calling again should not error
	err = EnsureConfigDir()
	if err != nil {
		t.Fatalf("EnsureConfigDir (second call) failed: %v", err)
	}
}

func TestPresetsFileExists(t *testing.T) {
	originalHome := os.Getenv("HOME")
	defer func() { _ = os.Setenv("HOME", originalHome) }()

	tmpDir := t.TempDir()
	if err := os.Setenv("HOME", tmpDir); err != nil {
		t.Fatalf("Failed to set HOME: %v", err)
	}

	// Initially should not exist
	exists, err := PresetsFileExists()
	if err != nil {
		t.Fatalf("PresetsFileExists failed: %v", err)
	}
	if exists {
		t.Error("Expected presets file to not exist initially")
	}

	// Create config dir and presets file
	if err := EnsureConfigDir(); err != nil {
		t.Fatalf("Failed to ensure config dir: %v", err)
	}
	presetsPath, _ := GetPresetsPath()
	if err := os.WriteFile(presetsPath, []byte("presets: {}"), 0o644); err != nil {
		t.Fatalf("Failed to write presets file: %v", err)
	}

	// Now should exist
	exists, err = PresetsFileExists()
	if err != nil {
		t.Fatalf("PresetsFileExists failed: %v", err)
	}
	if !exists {
		t.Error("Expected presets file to exist after creation")
	}
}

func TestLoadProjectConfig(t *testing.T) {
	tmpDir := t.TempDir()

	// Test loading non-existent config (should return defaults)
	cfg, err := LoadProjectConfig(tmpDir)
	if err != nil {
		t.Fatalf("LoadProjectConfig failed: %v", err)
	}
	if cfg == nil {
		t.Fatal("Expected non-nil config")
	}

	// Test loading existing config
	testCfg := &ProjectConfig{
		ComposeFile:   "docker-compose.dev.yml",
		DefaultPreset: "dev",
		Aliases: map[string]string{
			"db": "up postgres",
		},
		AutoPull: true,
	}
	if err := SaveProjectConfig(tmpDir, testCfg); err != nil {
		t.Fatalf("Failed to save project config: %v", err)
	}

	loaded, err := LoadProjectConfig(tmpDir)
	if err != nil {
		t.Fatalf("LoadProjectConfig failed: %v", err)
	}

	if loaded.ComposeFile != testCfg.ComposeFile {
		t.Errorf("ComposeFile mismatch: got %s, want %s", loaded.ComposeFile, testCfg.ComposeFile)
	}
	if loaded.DefaultPreset != testCfg.DefaultPreset {
		t.Errorf("DefaultPreset mismatch: got %s, want %s", loaded.DefaultPreset, testCfg.DefaultPreset)
	}
	if loaded.AutoPull != testCfg.AutoPull {
		t.Errorf("AutoPull mismatch: got %v, want %v", loaded.AutoPull, testCfg.AutoPull)
	}
}

func TestSaveProjectConfig(t *testing.T) {
	tmpDir := t.TempDir()

	cfg := &ProjectConfig{
		ComposeFile:   "custom-compose.yml",
		DefaultPreset: "production",
		Aliases: map[string]string{
			"start": "up -d",
			"stop":  "down",
		},
	}

	err := SaveProjectConfig(tmpDir, cfg)
	if err != nil {
		t.Fatalf("SaveProjectConfig failed: %v", err)
	}

	// Verify file was created
	configPath := filepath.Join(tmpDir, ProjectConfigName)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Error("Config file was not created")
	}

	// Verify content can be loaded back
	loaded, err := LoadProjectConfig(tmpDir)
	if err != nil {
		t.Fatalf("LoadProjectConfig failed: %v", err)
	}

	if loaded.ComposeFile != cfg.ComposeFile {
		t.Errorf("ComposeFile mismatch after save/load")
	}
}
