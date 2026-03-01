你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Common Functions

[![Lint and Testing](https://github.com/appleboy/com/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/com/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/com/actions/workflows/security.yml/badge.svg)](https://github.com/appleboy/com/actions/workflows/security.yml)
[![GoDoc](https://godoc.org/github.com/appleboy/com?status.svg)](https://godoc.org/github.com/appleboy/com)
[![codecov](https://codecov.io/gh/appleboy/com/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/com)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/com)](https://goreportcard.com/report/github.com/appleboy/com)

**Common Functions** is an open source collection of utility functions designed to simplify and accelerate Go development. This library provides robust, reusable helpers for common programming tasks such as random value generation, array and slice manipulation, file operations, and data type conversions. By centralizing frequently needed logic, it helps Go developers write cleaner, more efficient, and maintainable code across a wide range of projects.

## Package Usage

### array

Check if a value exists in a slice.

```go
import "github.com/appleboy/com/array"

found := array.Contains([]int{1, 2, 3}, 2) // true
```

### bytesconv

Zero-allocation conversion between string and []byte.

```go
import "github.com/appleboy/com/bytesconv"

b := bytesconv.StrToBytes("hello")
s := bytesconv.BytesToStr([]byte{'w', 'o', 'r', 'l', 'd'})
```

### convert

String case conversion, MD5 hashing, and float/byte conversion.

```go
import "github.com/appleboy/com/convert"

snake := convert.SnakeCasedName("FooBar") // "foo_bar"
title := convert.TitleCasedName("foo_bar") // "FooBar"
hash := convert.MD5Hash("data")
b := convert.Float64ToByte(3.14)
f := convert.ByteToFloat64(b)
```

### file

File and directory utilities.

```go
import "github.com/appleboy/com/file"

isDir, _ := file.IsDir("/tmp")
isFile, _ := file.IsFile("/tmp/file.txt")
_ = file.Copy("src.txt", "dst.txt")
_ = file.Remove("/tmp/old")
```

### gh

Set GitHub Actions output variables.

```go
import "github.com/appleboy/com/gh"

_ = gh.SetOutput(map[string]string{"key": "value"})
```

### random

Generate random strings for various use cases.

```go
import "github.com/appleboy/com/random"

s, _ := random.StringWithCharset(16, random.Alphanumeric) // secure random string
fast := random.randStringBytesMaskImprSrcUnsafe(16) // fast, not secure
```

### trace

Measure and log function execution time.

```go
import "github.com/appleboy/com/trace"

trace.ExecuteTime("myTask", func() {
    // code to measure
})
```
