# array

Utility functions for working with Go slices, particularly for checking element existence.

## Features

- Generic functions that work with any comparable type
- Efficient O(n) time complexity for slice operations
- Zero external dependencies

## Usage

### Check if an element exists in a slice

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/array"
)

func main() {
    // Works with any comparable type
    intSlice := []int{1, 2, 3, 4, 5}
    exists := array.Contains(intSlice, 3)
    fmt.Println("Contains 3?", exists) // Output: Contains 3? true
    
    stringSlice := []string{"hello", "world", "go"}
    exists = array.Contains(stringSlice, "go")
    fmt.Println("Contains 'go'?", exists) // Output: Contains 'go'? true
    
    // Returns false if not found
    exists = array.Contains(intSlice, 99)
    fmt.Println("Contains 99?", exists) // Output: Contains 99? false
}
```

## API Reference

### `Contains[T comparable](slice []T, key T) bool`

Checks if a given key of any comparable type exists within a slice.

**Parameters:**
- `slice`: A slice of any comparable type T
- `key`: An element of type T to search for within the slice

**Returns:**
- `bool`: True if the key is found in the slice, false otherwise

**Time Complexity:** O(n), where n is the length of the slice.

**Notes:**
- Suitable for small to medium slices or infrequent lookups
- For large slices with frequent lookups, consider using a map for better performance
- Works with any comparable type (int, string, float64, etc.)
