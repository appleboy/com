
# random

Cryptographically secure and fast random string and byte generation utilities with customizable character sets.

## Features

- Cryptographically secure random generation using `crypto/rand`
- Fast non-cryptographic random generation using `math/rand`
- Customizable character sets (Alphanumeric, Alphabetic, Numeric, Hex)
- Multiple generation methods optimized for different use cases
- Thread-safe operations

## Usage

### Cryptographically Secure Random Strings

```go
package main

import (
    "fmt"
    "log"
    "github.com/appleboy/com/random"
)

func main() {
    // Generate secure random string with custom charset
    str, err := random.StringWithCharset(16, random.Alphanumeric)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Secure random string: %s\n", str)
    
    // Generate with different charsets
    hexStr, _ := random.StringWithCharset(8, random.Hex)
    fmt.Printf("Hex string: %s\n", hexStr)
    
    numericStr, _ := random.StringWithCharset(6, random.Numeric)
    fmt.Printf("Numeric string: %s\n", numericStr)
}
```

### Fast Random Strings (Non-Cryptographic)

```go
package main

import (
    "fmt"
    "github.com/appleboy/com/random"
)

func main() {
    // Fast random string generation (not cryptographically secure)
    fastStr := random.String(10)
    fmt.Printf("Fast random string: %s\n", fastStr)
    
    // This is much faster but should not be used for security purposes
    for i := 0; i < 5; i++ {
        fmt.Printf("Fast #%d: %s\n", i+1, random.String(8))
    }
}
```

### Flexible Random String Generation

```go
package main

import (
    "fmt"
    "log"
    "github.com/appleboy/com/random"
)

func main() {
    // Use RandomString for flexible generation
    // Secure generation with custom charset
    secureStr, err := random.RandomString(12, random.Alphabetic, true)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Secure alphabetic: %s\n", secureStr)
    
    // Fast generation (ignores charset, always uses alphanumeric)
    fastStr, _ := random.RandomString(12, random.Hex, false)
    fmt.Printf("Fast generation: %s\n", fastStr)
    
    // Default charset when empty
    defaultStr, _ := random.RandomString(8, "", true)
    fmt.Printf("Default charset: %s\n", defaultStr)
}
```

### Custom Character Sets

```go
package main

import (
    "fmt"
    "log"
    "github.com/appleboy/com/random"
)

func main() {
    // Define custom character sets
    vowels := random.Charset("aeiouAEIOU")
    consonants := random.Charset("bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ")
    
    // Generate strings with custom charsets
    vowelStr, err := random.StringWithCharset(10, vowels)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Vowels only: %s\n", vowelStr)
    
    consonantStr, _ := random.StringWithCharset(10, consonants)
    fmt.Printf("Consonants only: %s\n", consonantStr)
    
    // Single character repeated
    singleChar := random.Charset("X")
    repeated, _ := random.StringWithCharset(5, singleChar)
    fmt.Printf("Repeated char: %s\n", repeated) // Output: XXXXX
}
```

### Random Bytes Generation

```go
package main

import (
    "fmt"
    "log"
    "github.com/appleboy/com/random"
)

func main() {
    // Note: randomBytes is not exported, but StringWithCharset uses it internally
    // For direct byte generation, you can use crypto/rand directly
    
    // Generate random string and convert to bytes if needed
    str, err := random.StringWithCharset(16, random.Hex)
    if err != nil {
        log.Fatal(err)
    }
    bytes := []byte(str)
    fmt.Printf("Random bytes: %v\n", bytes)
}
```

## API Reference

### Predefined Character Sets

```go
const (
    Alphanumeric Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    Alphabetic   Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    Numeric      Charset = "0123456789"
    Hex          Charset = "0123456789abcdef"
)
```

### Functions

#### `StringWithCharset(byteLen int, charset Charset) (string, error)`

Returns a cryptographically secure random string using the provided charset.

**Parameters:**

- `byteLen`: Length of the generated string in bytes
- `charset`: Character set to use for generation

**Returns:**

- `string`: Generated random string
- `error`: Error if random generation fails

**Security:** Uses `crypto/rand` - suitable for security-sensitive use cases.

#### `String(length int) string`

Returns a random string using a fast, non-cryptographically secure generator.

**Parameters:**

- `length`: Length of the generated string

**Returns:**

- `string`: Generated random string (alphanumeric characters only)

**Security:** Uses `math/rand` - NOT suitable for security-sensitive use cases.
**Performance:** Highly optimized for speed with bit masking techniques.

#### `RandomString(length int, charset Charset, secure bool) (string, error)`

Flexible random string generation with security and charset options.

**Parameters:**

- `length`: Length of the generated string
- `charset`: Character set to use (uses Alphanumeric if empty)
- `secure`: If true, uses cryptographically secure generation

**Returns:**

- `string`: Generated random string
- `error`: Error if secure generation fails (nil for non-secure)

**Behavior:**

- When `secure=true`: Uses `StringWithCharset()` with the provided charset
- When `secure=false`: Uses fast generation, ignores charset parameter
- Empty charset defaults to `Alphanumeric`

## Performance Characteristics

### Cryptographically Secure (`StringWithCharset`)

- **Security:** ✅ Cryptographically secure
- **Speed:** Slower due to system entropy usage
- **Use Cases:** Passwords, tokens, security keys, session IDs

### Fast Generation (`String`)

- **Security:** ❌ Not cryptographically secure
- **Speed:** Very fast with optimized bit masking
- **Use Cases:** Test data, temporary identifiers, non-security contexts

## Security Considerations

1. **Always use cryptographically secure functions** (`StringWithCharset`, `RandomString` with `secure=true`) for:
   - Passwords and passphrases
   - API keys and tokens
   - Session identifiers
   - Cryptographic nonces
   - Any security-sensitive random data

2. **Fast generation is acceptable for**:
   - Test data generation
   - Temporary file names
   - Non-security related identifiers
   - Performance-critical scenarios where security isn't required

## Thread Safety

- All functions are thread-safe
- Fast generation uses a mutex-protected `math/rand.Source`
- Secure generation uses `crypto/rand` which is inherently thread-safe

## Examples by Use Case

```go
// Password generation (secure)
password, _ := random.StringWithCharset(12, random.Alphanumeric)

// API token (secure)
token, _ := random.StringWithCharset(32, random.Hex)

// Test data (fast)
testID := random.String(8)

// Numeric PIN (secure)
pin, _ := random.StringWithCharset(4, random.Numeric)
```
