package gh

import (
	"os"
	"testing"
)

func TestSetOutput(t *testing.T) {
	t.Run("GITHUB_OUTPUT not set", func(t *testing.T) {
		os.Setenv("GITHUB_OUTPUT", "")
		err := SetOutput(map[string]string{"key": "value"})
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
		if err.Error() != "GITHUB_OUTPUT is not set" {
			t.Errorf("expected error message 'GITHUB_OUTPUT is not set' but got '%s'", err.Error())
		}
	})

	t.Run("GITHUB_OUTPUT set and file write successful", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "github_output")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(tempFile.Name())

		os.Setenv("GITHUB_OUTPUT", tempFile.Name())
		defer os.Unsetenv("GITHUB_OUTPUT")

		err = SetOutput(map[string]string{"key": "value"})
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		content, err := os.ReadFile(tempFile.Name())
		if err != nil {
			t.Fatalf("failed to read temp file: %v", err)
		}
		if !contains(string(content), "key=value\n") {
			t.Errorf("expected file content to contain 'key=value\\n' but got '%s'", string(content))
		}
	})

	t.Run("GITHUB_OUTPUT set but file write fails", func(t *testing.T) {
		os.Setenv("GITHUB_OUTPUT", "/invalid/path")
		defer os.Unsetenv("GITHUB_OUTPUT")

		err := SetOutput(map[string]string{"key": "value"})
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
		if !contains(err.Error(), "failed to open file") {
			t.Errorf("expected error message to contain 'failed to open file' but got '%s'", err.Error())
		}
	})
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
