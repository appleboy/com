# bytesconv

High-performance zero-allocation conversion between strings and byte slices for Go.

## Features

- Zero memory allocation conversions
- Support for Go 1.19+ and Go 1.20+ with optimized implementations
- Significant performance improvement over standard conversions
- Uses unsafe operations for maximum efficiency

## Usage

### Convert string to byte slice

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/bytesconv"
)

func main() {
    s := "hello world"
    b := bytesconv.StrToBytes(s)
    fmt.Printf("%v\n", b) // Output: [104 101 108 108 111 32 119 111 114 108 100]
    
    // Zero allocation - much faster than []byte(s)
    fmt.Printf("Length: %d\n", len(b)) // Output: Length: 11
}
```

### Convert byte slice to string

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/bytesconv"
)

func main() {
    b := []byte{'H', 'e', 'l', 'l', 'o'}
    s := bytesconv.BytesToStr(b)
    fmt.Println(s) // Output: Hello
    
    // Zero allocation - much faster than string(b)
}
```

### Round-trip conversion

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/bytesconv"
)

func main() {
    original := "Hello, 世界!"
    
    // String -> []byte -> string
    bytes := bytesconv.StrToBytes(original)
    back := bytesconv.BytesToStr(bytes)
    
    fmt.Println("Original:", original)
    fmt.Println("Round-trip:", back)
    fmt.Println("Equal:", original == back) // Output: Equal: true
}
```

## API Reference

### `StrToBytes(s string) []byte`

Converts a string to a byte slice without memory allocation.

**Parameters:**
- `s`: Input string to convert

**Returns:**
- `[]byte`: Byte slice representation of the string

### `BytesToStr(b []byte) string`

Converts a byte slice to a string without memory allocation.

**Parameters:**
- `b`: Input byte slice to convert

**Returns:**
- `string`: String representation of the byte slice

## Implementation Details

- **Go 1.20+**: Uses `unsafe.Slice()` and `unsafe.String()` for optimal performance
- **Go 1.19**: Uses manual unsafe pointer manipulation for compatibility
- **Build Tags**: Automatically selects the appropriate implementation based on Go version
- **Safety**: While using unsafe operations, the functions are safe when used correctly

## Performance

These functions provide significant performance benefits over standard conversions:
- `StrToBytes()` vs `[]byte(s)`: No allocation, much faster
- `BytesToStr()` vs `string(b)`: No allocation, much faster

## Warnings

- The returned byte slice from `StrToBytes()` shares memory with the original string
- The returned string from `BytesToStr()` shares memory with the original byte slice
- Do not modify the returned byte slice from `StrToBytes()` as it may cause undefined behavior
- These functions are safe for read-only operations and temporary conversions
