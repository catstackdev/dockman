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
	os.MkdirAll(subDir, 0o755)

	// Create docker-compose.yml in root
	composeFile := filepath.Join(tmpDir, "docker-compose.yml")
	os.WriteFile(composeFile, []byte("version: '3'"), 0o644)

	// Change to nested directory
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)
	os.Chdir(subDir)

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

	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)
	os.Chdir(tmpDir)

	// Should return error when no compose file found
	_, err := FindComposeFile()
	if err == nil {
		t.Error("Expected error when compose file not found")
	}
}

func TestFileExists(t *testing.T) {
	tmpDir := t.TempDir()

	// Test existing file
	existingFile := filepath.Join(tmpDir, "exists.txt")
	os.WriteFile(existingFile, []byte("test"), 0o644)

	if !fileExists(existingFile) {
		t.Error("Expected file to exist")
	}

	// Test non-existing file
	if fileExists(filepath.Join(tmpDir, "notexists.txt")) {
		t.Error("Expected file to not exist")
	}
}
