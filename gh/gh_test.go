package gh

import (
	"os"
	"strings"
	"testing"
)

// testEnv sets up a temporary file and environment for testing
func testEnv(t *testing.T) (string, func()) {
	t.Helper()
	tempFile, err := os.CreateTemp("", "github_output")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	if err := os.Setenv("GITHUB_OUTPUT", tempFile.Name()); err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		if err := os.Unsetenv("GITHUB_OUTPUT"); err != nil {
			t.Error(err)
		}
		if err := os.Remove(tempFile.Name()); err != nil {
			t.Error(err)
		}
	}

	return tempFile.Name(), cleanup
}

// readOutputFile reads the content of the output file
func readOutputFile(t *testing.T, filePath string) string {
	t.Helper()
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read temp file: %v", err)
	}
	return string(content)
}

// assertNoError asserts that error is nil
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}
}

// assertError asserts that error is not nil
func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("expected an error but got nil")
	}
}

// assertErrorMessage asserts that error message matches expected
func assertErrorMessage(t *testing.T, err error, expected string) {
	t.Helper()
	if err.Error() != expected {
		t.Errorf("expected error message '%s' but got '%s'", expected, err.Error())
	}
}

// assertContains asserts that string contains substring
func assertContains(t *testing.T, s, substr string) {
	t.Helper()
	if !strings.Contains(s, substr) {
		t.Errorf("expected content to contain '%s' but got '%s'", substr, s)
	}
}

func TestSetOutputNotSet(t *testing.T) {
	if err := os.Setenv("GITHUB_OUTPUT", ""); err != nil {
		t.Fatal(err)
	}
	err := SetOutput(map[string]string{"key": "value"})
	assertError(t, err)
	assertErrorMessage(t, err, "GITHUB_OUTPUT is not set")
}

func TestSetOutputSuccess(t *testing.T) {
	filePath, cleanup := testEnv(t)
	defer cleanup()

	err := SetOutput(map[string]string{"key": "value"})
	assertNoError(t, err)

	content := readOutputFile(t, filePath)
	if !contains(content, "key=value\n") {
		t.Errorf("expected file content to contain 'key=value\\n' but got '%s'", content)
	}
}

func TestSetOutputFileWriteFails(t *testing.T) {
	if err := os.Setenv("GITHUB_OUTPUT", "/invalid/path"); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Unsetenv("GITHUB_OUTPUT"); err != nil {
			t.Fatal(err)
		}
	}()

	err := SetOutput(map[string]string{"key": "value"})
	assertError(t, err)
	if !contains(err.Error(), "failed to open file") {
		t.Errorf(
			"expected error message to contain 'failed to open file' but got '%s'",
			err.Error(),
		)
	}
}

func TestSetOutputMultiline(t *testing.T) {
	filePath, cleanup := testEnv(t)
	defer cleanup()

	multilineValue := "line1\nline2\nline3"
	err := SetOutput(map[string]string{"multiline": multilineValue})
	assertNoError(t, err)

	content := readOutputFile(t, filePath)
	assertContains(t, content, "multiline<<")
	assertContains(t, content, "line1\nline2\nline3")
	assertContains(t, content, "ghdelimiter")
}

func TestSetOutputMixed(t *testing.T) {
	filePath, cleanup := testEnv(t)
	defer cleanup()

	err := SetOutput(map[string]string{
		"single": "value",
		"multi":  "line1\nline2",
	})
	assertNoError(t, err)

	content := readOutputFile(t, filePath)
	assertContains(t, content, "single=value\n")
	assertContains(t, content, "multi<<")
	assertContains(t, content, "line1\nline2")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
