package compose

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewExecutorWithFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a compose file
	composeFile := filepath.Join(tmpDir, "docker-compose.yml")
	if err := os.WriteFile(composeFile, []byte("version: '3'\nservices:\n  web:\n    image: nginx"), 0o644); err != nil {
		t.Fatalf("Failed to write compose file: %v", err)
	}

	// Test with existing file
	executor, err := NewExecutorWithFile(composeFile)
	if err != nil {
		t.Fatalf("NewExecutorWithFile failed: %v", err)
	}

	if executor.ComposeFile != composeFile {
		t.Errorf("Expected ComposeFile %s, got %s", composeFile, executor.ComposeFile)
	}

	if executor.ProjectDir != tmpDir {
		t.Errorf("Expected ProjectDir %s, got %s", tmpDir, executor.ProjectDir)
	}
}

func TestNewExecutorWithFile_NotFound(t *testing.T) {
	_, err := NewExecutorWithFile("/nonexistent/docker-compose.yml")
	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
}

func TestExecutor_GetInfo(t *testing.T) {
	tmpDir := t.TempDir()
	composeFile := filepath.Join(tmpDir, "docker-compose.yml")
	if err := os.WriteFile(composeFile, []byte("version: '3'"), 0o644); err != nil {
		t.Fatalf("Failed to write compose file: %v", err)
	}

	executor, _ := NewExecutorWithFile(composeFile)
	info := executor.GetInfo()

	if info == "" {
		t.Error("GetInfo returned empty string")
	}

	// Should contain project dir and compose file
	if !contains(info, tmpDir) {
		t.Errorf("GetInfo should contain project dir: %s", info)
	}
	if !contains(info, composeFile) {
		t.Errorf("GetInfo should contain compose file: %s", info)
	}
}

func TestGetProjectDir(t *testing.T) {
	tests := []struct {
		composePath string
		expected    string
	}{
		{"/home/user/project/docker-compose.yml", "/home/user/project"},
		{"/var/app/compose.yml", "/var/app"},
		{"./docker-compose.yml", "."},
	}

	for _, tt := range tests {
		t.Run(tt.composePath, func(t *testing.T) {
			result := GetProjectDir(tt.composePath)
			if result != tt.expected {
				t.Errorf("GetProjectDir(%s) = %s, want %s", tt.composePath, result, tt.expected)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
