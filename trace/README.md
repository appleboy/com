
# trace

Simple execution time measurement and logging utilities for Go applications, designed for performance monitoring and debugging.

## Features

- Measure and log function execution time
- Millisecond precision timing
- Simple, clean logging output
- Zero dependencies (uses standard library only)
- Minimal performance overhead

## Usage

### Basic Execution Time Measurement

```go
package main

import (
    "time"
    "github.com/appleboy/com/trace"
)

func main() {
    // Measure execution time of a function
    trace.ExecuteTime("database query", func() {
        // Simulate database operation
        time.Sleep(100 * time.Millisecond)
    })
    // Output: [database query] elapsed=100ms
}
```

### Multiple Operations

```go
package main

import (
    "time"
    "github.com/appleboy/com/trace"
)

func main() {
    trace.ExecuteTime("step 1: initialization", func() {
        time.Sleep(50 * time.Millisecond)
    })
    
    trace.ExecuteTime("step 2: processing", func() {
        time.Sleep(200 * time.Millisecond)
    })
    
    trace.ExecuteTime("step 3: cleanup", func() {
        time.Sleep(30 * time.Millisecond)
    })
    
    // Output:
    // [step 1: initialization] elapsed=50ms
    // [step 2: processing] elapsed=200ms
    // [step 3: cleanup] elapsed=30ms
}
```

### Real-World Examples

```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/appleboy/com/trace"
)

func fetchUserData(userID string) map[string]interface{} {
    var result map[string]interface{}
    
    trace.ExecuteTime("fetch user "+userID, func() {
        // Simulate API call
        resp, err := http.Get(fmt.Sprintf("https://api.example.com/users/%s", userID))
        if err != nil {
            return
        }
        defer resp.Body.Close()
        
        body, _ := ioutil.ReadAll(resp.Body)
        json.Unmarshal(body, &result)
    })
    
    return result
}

func processData(data []string) []string {
    var processed []string
    
    trace.ExecuteTime("data processing", func() {
        for _, item := range data {
            // Simulate processing
            processed = append(processed, "processed_"+item)
        }
    })
    
    return processed
}

func main() {
    // Measure API calls
    user1 := fetchUserData("123")
    user2 := fetchUserData("456")
    
    // Measure data processing
    data := []string{"item1", "item2", "item3"}
    processed := processData(data)
    
    fmt.Printf("Users: %v, %v\n", user1, user2)
    fmt.Printf("Processed: %v\n", processed)
}
```

### Nested Measurements

```go
package main

import (
    "time"
    "github.com/appleboy/com/trace"
)

func complexOperation() {
    trace.ExecuteTime("complex operation", func() {
        trace.ExecuteTime("  - sub-operation 1", func() {
            time.Sleep(100 * time.Millisecond)
        })
        
        trace.ExecuteTime("  - sub-operation 2", func() {
            time.Sleep(150 * time.Millisecond)
        })
        
        trace.ExecuteTime("  - sub-operation 3", func() {
            time.Sleep(75 * time.Millisecond)
        })
    })
}

func main() {
    complexOperation()
    
operation 1] elapsed=100ms
    // [  - sub-operation 2] elapsed=150ms
    // [  - sub-operation 3] elapsed=75ms
    // [complex operation] elapsed=325ms
}
```

### Performance Benchmarking

```go
package main

import (
    "github.com/appleboy/com/trace"
)

func algorithmA(data []int) int {
    var result int
    trace.ExecuteTime("Algorithm A", func() {
        for _, v := range data {
            result += v * v
        }
    })
    return result
}

func algorithmB(data []int) int {
    var result int
    trace.ExecuteTime("Algorithm B", func() {
        for i := 0; i < len(data); i++ {
            result += data[i] * data[i]
        }
    })
    return result
}

func main() {
    data := make([]int, 1000000)
    for i := range data {
        data[i] = i
    }
    
    // Compare performance
    resultA := algorithmA(data)
    resultB := algorithmB(data)
    
    // Output will show timing comparison:
    // [Algorithm A] elapsed=2ms
    // [Algorithm B] elapsed=3ms
}
```

## API Reference

### `ExecuteTime(title string, fn func())`

Executes the provided function and logs the execution time.

**Parameters:**

- `title`: A descriptive name for the operation being measured
- `fn`: The function to execute and measure

**Behavior:**

- Records start time using `time.Now()`
- Executes the provided function
- Calculates elapsed time using `time.Since()`
- Logs the result in the format: `[title] elapsed=NNms`

**Output Format:**

```txt
[title] elapsed=123ms
```

**Logging:**

- Uses Go's standard `log` package
- Output includes timestamp and other standard log formatting
- Precision: Milliseconds (uses `time.Duration.Milliseconds()`)

## Use Cases

### Development and Debugging

- Identify performance bottlenecks
- Monitor function execution times during development
- Debug slow operations

### Performance Monitoring

- Track operation performance in production
- Compare algorithm performance
- Monitor API response times

### Optimization

- Before/after performance comparisons
- Identify optimization opportunities
- Validate performance improvements

## Best Practices

1. **Descriptive Titles**: Use clear, descriptive titles for easy identification

   ```go
   trace.ExecuteTime("user authentication", func() { ... })
   trace.ExecuteTime("database migration step 3", func() { ... })
   ```

2. **Granular Measurements**: Measure specific operations rather than large blocks

   ```go
   // Good - specific operations
   trace.ExecuteTime("parse JSON", func() { ... })
   trace.ExecuteTime("validate data", func() { ... })
   
   // Less useful - too broad
   trace.ExecuteTime("entire application", func() { ... })
   ```

3. **Consistent Naming**: Use consistent naming conventions for related operations

   ```go
   trace.ExecuteTime("db: insert user", func() { ... })
   trace.ExecuteTime("db: update profile", func() { ... })
   trace.ExecuteTime("db: delete session", func() { ... })
   ```

## Performance Impact

The `trace.ExecuteTime` function has minimal performance overhead:

- Time measurement: ~100-500 nanoseconds
- Logging: Depends on log configuration
- Function call overhead: Negligible

For production use, consider:

- Using build tags to conditionally include tracing
- Implementing sampling for high-frequency operations
- Custom logging destinations for performance data

## Integration with Logging Systems

The package uses Go's standard `log` package, which can be configured to write to different outputs:

```go
// Redirect to file
logFile, _ := os.Create("performance.log")
log.SetOutput(logFile)

// Customize format
log.SetFlags(log.LstdFlags | log.L
    // Output:
    // [  - sub-

shortfile)

// Use with structured logging libraries
trace.ExecuteTime("operation", func() {
    // Your operation here
})
```

This allows integration with existing logging infrastructure and monitoring systems.
