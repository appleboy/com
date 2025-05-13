package gh

import (
	"errors"
	"fmt"
	"os"
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
		_, err = fmt.Fprintf(file, "%s=%s\n", k, v)
		if err != nil {
			return fmt.Errorf("failed to write to file %s: %w", githubOutput, err)
		}
	}

	return nil
}
