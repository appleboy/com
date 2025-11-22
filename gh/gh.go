package gh

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
)

func SetOutput(data map[string]string) error {
	githubOutput := os.Getenv("GITHUB_OUTPUT")
	if githubOutput == "" {
		return errors.New("GITHUB_OUTPUT is not set")
	}

	file, err := os.OpenFile(githubOutput, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", githubOutput, err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			// You can log or handle the error here if needed
			_ = cerr
		}
	}()

	for k, v := range data {
		if strings.Contains(v, "\n") {
			// Use heredoc syntax for multiline values
			delimiter := generateDelimiter()
			_, err = fmt.Fprintf(file, "%s<<%s\n%s\n%s\n", k, delimiter, v, delimiter)
			if err != nil {
				return fmt.Errorf("failed to write to file %s: %w", githubOutput, err)
			}
		} else {
			// Use simple format for single-line values
			_, err = fmt.Fprintf(file, "%s=%s\n", k, v)
			if err != nil {
				return fmt.Errorf("failed to write to file %s: %w", githubOutput, err)
			}
		}
	}

	return nil
}

// generateDelimiter generates a unique delimiter for multiline values
func generateDelimiter() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		// Fallback to a simple delimiter if random generation fails
		return "ghdelimiter"
	}
	return "ghdelimiter_" + hex.EncodeToString(b)
}
