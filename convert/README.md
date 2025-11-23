
# convert

Comprehensive type conversion utilities for Go, including basic types, pointers, collections, and specialized conversions.

## Features

- Convert between basic types (string, bool, int, float)
- Pointer conversion utilities with generics
- Collection conversion (slice/map to pointer variants and vice versa)
- String manipulation utilities (snake_case, TitleCase, MD5 hashing)
- Binary conversion utilities (float64 to bytes)
- Big5 to UTF-8 encoding conversion

## Usage

### Basic Type Conversions

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/convert"
)

func main() {
    // Convert any value to string
    str := convert.ToString(123)
    fmt.Println(str) // Output: 123
    
    str = convert.ToString(true)
    fmt.Println(str) // Output: true
    
    // Convert to boolean
    b := convert.ToBool("true")
    fmt.Println(b) // Output: true
    
    b = convert.ToBool(1)
    fmt.Println(b) // Output: true
    
    b = convert.ToBool("")
    fmt.Println(b) // Output: false
    
    // Convert to int (returns any - nil if conversion fails)
    if i := convert.ToInt("123"); i != nil {
        fmt.Println(i) // Output: 123
    }
    
    // Convert to float
    if f := convert.ToFloat("123.45"); f != nil {
        fmt.Println(f) // Output: 123.45
    }
}
```

### Pointer Utilities

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/convert"
)

func main() {
    // Convert value to pointer
    value := 42
    ptr := convert.ToPtr(value)
    fmt.Println(*ptr) // Output: 42
    
    // Convert pointer to value (returns zero value if nil)
    result := convert.FromPtr(ptr)
    fmt.Println(result) // Output: 42
    
    // Handle nil pointer
    var nilPtr *int
    result = convert.FromPtr(nilPtr)
    fmt.Println(result) // Output: 0
}
```

### Collection Conversions

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/convert"
)

func main() {
    // Convert slice to pointer slice
    values := []int{1, 2, 3}
    ptrs := convert.SliceToPtrSlice(values)
    fmt.Println(*ptrs[0]) // Output: 1
    
    // Convert pointer slice back to values
    backToValues := convert.PtrSliceToSlice(ptrs)
    fmt.Println(backToValues) // Output: [1 2 3]
    
    // Convert map to pointer map
    valueMap := map[string]int{"a": 1, "b": 2}
    ptrMap := convert.MapToPtrMap(valueMap)
    fmt.Println(*ptrMap["a"]) // Output: 1
    
    // Convert pointer map back to values (skips nil pointers)
    backToMap := convert.PtrMapToMap(ptrMap)
    fmt.Println(backToMap) // Output: map[a:1 b:2]
}
```

### String Utilities

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/convert"
)

func main() {
    // Convert to snake_case
    snake := convert.SnakeCasedName("FooBarTest")
    fmt.Println(snake) // Output: foo_bar_test
    
    // Convert to TitleCase
    title := convert.TitleCasedName("foo_bar_test")
    fmt.Println(title) // Output: FooBarTest
    
    // MD5 hash
    hash := convert.MD5Hash("hello world")
    fmt.Println(hash) // Output: 5d41402abc4b2a76b9719d911017c592
}
```

### Binary Conversion

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/convert"
)

func main() {
    // Convert float64 to bytes (BigEndian)
    f := 123.456
    bytes := convert.Float64ToByte(f)
    fmt.Printf("Bytes: %v\n", bytes)
    
    // Convert bytes back to float64
    back := convert.ByteToFloat64(bytes)
    fmt.Printf("Float: %f\n", back) // Output: Float: 123.456000
}
```

### Encoding Conversion

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/convert"
)

func main() {
    // Convert Big5 encoded string to UTF-8
    big5Str := "\xa7A\xa6n" // "你好" in Big5
    utf8Str := convert.ConvertBig5ToUTF8(big5Str)
    fmt.Println(utf8Str) // Output: 你好
}
```
