package config

import (
	"reflect"
	"testing"
)

func TestResolveAlias(t *testing.T) {
	// Create temp directory with config
	tmpDir := t.TempDir()

	cfg := &ProjectConfig{
		Aliases: map[string]string{
			"db":  "up postgres redis",
			"api": "up api",
		},
	}

	if err := SaveProjectConfig(tmpDir, cfg); err != nil {
		t.Fatalf("Failed to save project config: %v", err)
	}

	tests := []struct {
		name     string
		args     []string
		expected []string
		isAlias  bool
	}{
		{
			name:     "resolve db alias",
			args:     []string{"db"},
			expected: []string{"up", "postgres", "redis"},
			isAlias:  true,
		},
		{
			name:     "resolve api alias",
			args:     []string{"api"},
			expected: []string{"up", "api"},
			isAlias:  true,
		},
		{
			name:     "not an alias",
			args:     []string{"up"},
			expected: []string{"up"},
			isAlias:  false,
		},
		{
			name:     "alias with extra args",
			args:     []string{"db", "-f"},
			expected: []string{"up", "postgres", "redis", "-f"},
			isAlias:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, isAlias, err := ResolveAlias(tmpDir, tt.args)
			if err != nil {
				t.Fatalf("ResolveAlias failed: %v", err)
			}

			if isAlias != tt.isAlias {
				t.Errorf("Expected isAlias=%v, got %v", tt.isAlias, isAlias)
			}

			if isAlias && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestResolveAlias_EmptyArgs(t *testing.T) {
	tmpDir := t.TempDir()

	result, isAlias, err := ResolveAlias(tmpDir, []string{})
	if err != nil {
		t.Fatalf("ResolveAlias failed: %v", err)
	}

	if isAlias {
		t.Error("Expected isAlias=false for empty args")
	}

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
}

func TestResolveAlias_NoConfig(t *testing.T) {
	tmpDir := t.TempDir()

	result, isAlias, err := ResolveAlias(tmpDir, []string{"up"})
	if err != nil {
		t.Fatalf("ResolveAlias failed: %v", err)
	}

	if isAlias {
		t.Error("Expected isAlias=false when no config exists")
	}

	if !reflect.DeepEqual(result, []string{"up"}) {
		t.Errorf("Expected original args, got %v", result)
	}
}

func TestSplitArgs(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"up postgres redis", []string{"up", "postgres", "redis"}},
		{"up api", []string{"up", "api"}},
		{"logs -f api", []string{"logs", "-f", "api"}},
		{`exec api "bash -c 'echo hello'"`, []string{"exec", "api", "bash -c 'echo hello'"}},
		{"  spaced  args  ", []string{"spaced", "args"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := splitArgs(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("splitArgs(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
