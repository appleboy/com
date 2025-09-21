
# gh

GitHub Actions output utilities for setting environment variables and outputs in CI/CD workflows.

## Features

- Set multiple GitHub Actions output variables in a single call
- Proper error handling for missing environment variables
- Safe file operations with error checking
- Designed specifically for GitHub Actions workflows

## Usage

### Basic Output Setting

```go
package main

import (
    "log"
    "github.com/appleboy/com/gh"
)

func main() {
    // Set single output variable
    data := map[string]string{
        "version": "1.2.3",
    }
    
    err := gh.SetOutput(data)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Multiple Outputs

```go
package main

import (
    "fmt"
    "log"
    "github.com/appleboy/com/gh"
)

func main() {
    // Set multiple output variables
    outputs := map[string]string{
        "build_status": "success",
        "version":      "v1.2.3",
        "commit_sha":   "abc123def456",
        "environment":  "production",
    }
    
    err := gh.SetOutput(outputs)
    if err != nil {
        log.Fatalf("Failed to set GitHub outputs: %v", err)
    }
    
    fmt.Println("GitHub Actions outputs set successfully!")
}
```

### In GitHub Actions Workflow

```yaml
# .github/workflows/example.yml
name: Example Workflow
on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v3
        with:
          go-version: 1.19
          
      - name: Run Go program that sets outputs
        run: go run main.go
        
      - name: Use the outputs
        run: |
          echo "Build status: ${{ steps.previous-step.outputs.build_status }}"
          echo "Version: ${{ steps.previous-step.outputs.version }}"
```

### Error Handling

```go
package main

import (
    "fmt"
    "log"
    "os"
    "github.com/appleboy/com/gh"
)

func main() {
    // Check if running in GitHub Actions
    if os.Getenv("GITHUB_ACTIONS") != "true" {
        fmt.Println("Not running in GitHub Actions, skipping output setting")
        return
    }
    
    outputs := map[string]string{
        "result": "success",
    }
    
    if err := gh.SetOutput(outputs); err != nil {
        log.Printf("Warning: Failed to set GitHub outputs: %v", err)
        // Continue execution - don't fail the entire workflow
    }
}
```

## API Reference

### `SetOutput(data map[string]string) error`

Writes key-value pairs to the GitHub Actions output environment file.

**Parameters:**

- `data`: Map of output variable names to their values

**Returns:**

- `error`: Error if operation fails, nil on success

**Behavior:**

- Reads the `GITHUB_OUTPUT` environment variable to get the output file path
- Appends each key-value pair to the file in the format `key=value\n`
- Returns error if `GITHUB_OUTPUT` is not set
- Returns error if file operations fail

**Environment Requirements:**

- Must be running in GitHub Actions environment
- `GITHUB_OUTPUT` environment variable must be set by GitHub Actions

**Example Output File Content:**

```txt
version=1.2.3
build_status=success
commit_sha=abc123def456
```

## Usage in GitHub Actions

This package is specifically designed for use within GitHub Actions workflows. The `GITHUB_OUTPUT` environment variable is automatically set by GitHub Actions and points to a temporary file that GitHub reads to capture workflow outputs.

### Setting Job Outputs

```go
// In your Go program running in GitHub Actions
outputs := map[string]string{
    "artifact_name": "app-v1.2.3.zip",
    "deploy_url": "https://deploy.example.com/v1.2.3",
}
gh.SetOutput(outputs)
```

These outputs can then be used in subsequent workflow steps or jobs:

```yaml
- name: Deploy
  run: |
    echo "Deploying ${{ steps.build.outputs.artifact_name }}"
    echo "Deploy URL: ${{ steps.build.outputs.deploy_url }}"
```

## Error Conditions

The function will return an error in the following cases:

1. **GITHUB_OUTPUT not set**: When not running in GitHub Actions environment
2. **File access errors**: When the output file cannot be opened or written to
3. **Write errors**: When writing to the output file fails

## Best Practices

1. **Environment Check**: Always check if running in GitHub Actions before calling
2. **Error Handling**: Handle errors gracefully to avoid workflow failures
3. **Output Naming**: Use descriptive, consistent naming for output variables
4. **Value Sanitization**: Ensure output values don't contain newlines or special characters

## Notes

- Output values are written as-is without escaping
- Each call appends to the existing output file
- The function is safe for concurrent use within a single workflow step
- Maximum output size is limited by GitHub Actions (typically 1MB per job)
