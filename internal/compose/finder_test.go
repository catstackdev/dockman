package compose

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindComposeFile(t *testing.T) {
	// Create temp directory structure
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "subdir", "nested")
	if err := os.MkdirAll(subDir, 0o755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}

	// Create docker-compose.yml in root
	composeFile := filepath.Join(tmpDir, "docker-compose.yml")
	if err := os.WriteFile(composeFile, []byte("version: '3'"), 0o644); err != nil {
		t.Fatalf("Failed to write compose file: %v", err)
	}

	// Change to nested directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer func() { _ = os.Chdir(originalDir) }()
	if err := os.Chdir(subDir); err != nil {
		t.Fatalf("Failed to change to subdir: %v", err)
	}

	// Test finding file from subdirectory
	found, err := FindComposeFile()
	if err != nil {
		t.Fatalf("FindComposeFile failed: %v", err)
	}

	// Resolve symlinks for comparison (macOS /var -> /private/var)
	expected, _ := filepath.EvalSymlinks(composeFile)
	foundResolved, _ := filepath.EvalSymlinks(found)
	if foundResolved != expected {
		t.Errorf("Expected %s, got %s", expected, foundResolved)
	}
}

func TestFindComposeFile_NotFound(t *testing.T) {
	// Create empty temp directory
	tmpDir := t.TempDir()

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer func() { _ = os.Chdir(originalDir) }()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to tmpdir: %v", err)
	}

	// Should return error when no compose file found
	_, err = FindComposeFile()
	if err == nil {
		t.Error("Expected error when compose file not found")
	}
}

func TestFileExists(t *testing.T) {
	tmpDir := t.TempDir()

	// Test existing file
	existingFile := filepath.Join(tmpDir, "exists.txt")
	if err := os.WriteFile(existingFile, []byte("test"), 0o644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	if !fileExists(existingFile) {
		t.Error("Expected file to exist")
	}

	// Test non-existing file
	if fileExists(filepath.Join(tmpDir, "notexists.txt")) {
		t.Error("Expected file to not exist")
	}
}
