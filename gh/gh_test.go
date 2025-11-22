package gh

import (
	"os"
	"strings"
	"testing"
)

func TestSetOutput(t *testing.T) {
	t.Run("GITHUB_OUTPUT not set", func(t *testing.T) {
		if err := os.Setenv("GITHUB_OUTPUT", ""); err != nil {
			t.Fatal(err)
		}
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
		defer func() {
			if err := os.Remove(tempFile.Name()); err != nil {
				t.Fatal(err)
			}
		}()

		if err := os.Setenv("GITHUB_OUTPUT", tempFile.Name()); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := os.Unsetenv("GITHUB_OUTPUT"); err != nil {
				t.Fatal(err)
			}
		}()

		err = SetOutput(map[string]string{"key": "value"})
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		content, err := os.ReadFile(tempFile.Name())
		if err != nil {
			t.Fatalf("failed to read temp file: %v", err)
		}
		if !contains(string(content), "key=value\n") {
			t.Errorf(
				"expected file content to contain 'key=value\\n' but got '%s'",
				string(content),
			)
		}
	})

	t.Run("GITHUB_OUTPUT set but file write fails", func(t *testing.T) {
		if err := os.Setenv("GITHUB_OUTPUT", "/invalid/path"); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := os.Unsetenv("GITHUB_OUTPUT"); err != nil {
				t.Fatal(err)
			}
		}()

		err := SetOutput(map[string]string{"key": "value"})
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
		if !contains(err.Error(), "failed to open file") {
			t.Errorf(
				"expected error message to contain 'failed to open file' but got '%s'",
				err.Error(),
			)
		}
	})

	t.Run("multiline value with heredoc syntax", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "github_output")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer func() {
			if err := os.Remove(tempFile.Name()); err != nil {
				t.Fatal(err)
			}
		}()

		if err := os.Setenv("GITHUB_OUTPUT", tempFile.Name()); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := os.Unsetenv("GITHUB_OUTPUT"); err != nil {
				t.Fatal(err)
			}
		}()

		multilineValue := "line1\nline2\nline3"
		err = SetOutput(map[string]string{"multiline": multilineValue})
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		content, err := os.ReadFile(tempFile.Name())
		if err != nil {
			t.Fatalf("failed to read temp file: %v", err)
		}

		contentStr := string(content)
		// Check that it uses heredoc syntax
		if !strings.Contains(contentStr, "multiline<<") {
			t.Errorf("expected file content to use heredoc syntax but got '%s'", contentStr)
		}
		// Check that the multiline value is present
		if !strings.Contains(contentStr, "line1\nline2\nline3") {
			t.Errorf(
				"expected file content to contain the multiline value but got '%s'",
				contentStr,
			)
		}
		// Check that it has a delimiter
		if !strings.Contains(contentStr, "ghdelimiter") {
			t.Errorf("expected file content to contain delimiter but got '%s'", contentStr)
		}
	})

	t.Run("mixed single-line and multiline values", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "github_output")
		if err != nil {
			t.Fatalf("failed to create temp file: %v", err)
		}
		defer func() {
			if err := os.Remove(tempFile.Name()); err != nil {
				t.Fatal(err)
			}
		}()

		if err := os.Setenv("GITHUB_OUTPUT", tempFile.Name()); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := os.Unsetenv("GITHUB_OUTPUT"); err != nil {
				t.Fatal(err)
			}
		}()

		err = SetOutput(map[string]string{
			"single": "value",
			"multi":  "line1\nline2",
		})
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		content, err := os.ReadFile(tempFile.Name())
		if err != nil {
			t.Fatalf("failed to read temp file: %v", err)
		}

		contentStr := string(content)
		// Check single-line format
		if !strings.Contains(contentStr, "single=value\n") {
			t.Errorf("expected file content to contain 'single=value\\n' but got '%s'", contentStr)
		}
		// Check multiline format uses heredoc
		if !strings.Contains(contentStr, "multi<<") {
			t.Errorf(
				"expected file content to use heredoc syntax for multiline but got '%s'",
				contentStr,
			)
		}
		if !strings.Contains(contentStr, "line1\nline2") {
			t.Errorf(
				"expected file content to contain the multiline value but got '%s'",
				contentStr,
			)
		}
	})
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
