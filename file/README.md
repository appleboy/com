
# file

File and directory operations utilities for Go, providing a comprehensive set of functions for path validation, file manipulation, and size formatting.

## Features

- Path validation (check if path is file or directory)
- File and directory removal
- Safe file copying with existence checks
- Human-readable file size formatting
- Error handling with detailed messages

## Usage

### Check Path Types

```go
package main

import (
    "fmt"
    "log"
    "github.com/appleboy/com/file"
)

func main() {
    // Check if path is a directory
    isDir, err := file.IsDir("./testdata")
    if err != nil {
        log.Printf("Error checking directory: %v", err)
    } else {
        fmt.Printf("Is directory: %t\n", isDir)
    }
    
    // Check if path is a file
    isFile, err := file.IsFile("./main.go")
    if err != nil {
        log.Printf("Error checking file: %v", err)
    } else {
        fmt.Printf("Is file: %t\n", isFile)
    }
}
```

### File Operations

```go
package main

import (
    "log"
    "github.com/appleboy/com/file"
)

func main() {
    // Copy a file (fails if destination exists)
    err := file.Copy("source.txt", "destination.txt")
    if err != nil {
        log.Printf("Copy failed: %v", err)
    }
    
    // Remove file or directory (including all contents)
    err = file.Remove("temp_file.txt")
    if err != nil {
        log.Printf("Remove failed: %v", err)
    }
    
    // Remove directory with all contents
    err = file.Remove("temp_directory")
    if err != nil {
        log.Printf("Remove failed: %v", err)
    }
}
```

### File Size Formatting

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/file"
)

func main() {
    // Format various file sizes
    fmt.Println(file.FormatSize(512))        // Output: 512 B
    fmt.Println(file.FormatSize(1024))       // Output: 1.0 KB
    fmt.Println(file.FormatSize(1536))       // Output: 1.5 KB
    fmt.Println(file.FormatSize(1048576))    // Output: 1.0 MB
    fmt.Println(file.FormatSize(1073741824)) // Output: 1.0 GB
}
```

### Complete Example

```go
package main

import (
    "fmt"
    "log" 
    "os"
    "github.com/appleboy/com/file"
)

func main() {
    // Create a test file
    testFile := "example.txt"
    err := os.WriteFile(testFile, []byte("Hello, World!"), 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    // Check if it's a file
    if isFile, err := file.IsFile(testFile); err == nil && isFile {
        fmt.Printf("%s is a file\n", testFile)
        
        // Get file info for size
        if info, err := os.Stat(testFile); err == nil {
            size := file.FormatSize(info.Size())
            fmt.Printf("File size: %s\n", size)
        }
        
        // Copy the file
        copyFile := "example_copy.txt"
        if err := file.Copy(testFile, copyFile); err != nil {
            log.Printf("Copy failed: %v", err)
        } else {
            fmt.Printf("Copied %s to %s\n", testFile, copyFile)
        }
        
        // Cleanup
        file.Remove(testFile)
        file.Remove(copyFile)
    }
}


## API Reference

### `IsDir(dir string) (bool, error)`

Returns true if the given path is a directory.

**Parameters:**
- `dir`: Path to check

**Returns:**
- `bool`: True if path is a directory, false if it's a file
- `error`: Error if path doesn't exist or other issues occur

### `IsFile(filePath string) (bool, error)`

Returns true if the given path is a file.

**Parameters:**
- `filePath`: Path to check

**Returns:**
- `bool`: True if path is a file, false if it's a directory
- `error`: Error if path doesn't exist or other issues occur

### `Remove(filePath string) error`

Removes the file or directory at the given path, including any children if it's a directory.

**Parameters:**
- `filePath`: Path to remove

**Returns:**
- `error`: Error if removal fails, nil if successful or path doesn't exist

**Notes:**
- Uses `os.RemoveAll()` internally
- Safe to call on non-existent paths (returns nil)

### `Copy(src, dst string) error`

Copies a regular file from src to dst. Fails if destination already exists.

**Parameters:**
- `src`: Source file path
- `dst`: Destination file path

**Returns:**
- `error`: Error if copy fails

**Notes:**
- Only works with regular files (not directories or special files)
- Uses `io.Copy()` for efficient transfer
- Fails if destination already exists
- Preserves file content but not metadata

### `FormatSize(bytes int64) string`

Returns a human-readable string for a file size in bytes.

**Parameters:**
- `bytes`: File size in bytes

**Returns:**
- `string`: Formatted size (e.g., "1.2 MB", "512 B")

**Examples:**
- `0` → `"0 B"`
- `1024` → `"1.0 KB"`
- `1536` → `"1.5 KB"`
- `1048576` → `"1.0 MB"`
- `1073741824` → `"1.0 GB"`

**Notes:**
- Uses 1024 as the unit base
- Supports up to Exabytes (EB)
- Always shows one decimal place for units KB and above
