你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Common Functions

[![GoDoc](https://godoc.org/github.com/appleboy/com?status.svg)](https://godoc.org/github.com/appleboy/com) 
[![Build Status](https://cloud.drone.io/api/badges/appleboy/com/status.svg)](https://cloud.drone.io/appleboy/com)
[![codecov](https://codecov.io/gh/appleboy/com/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/com) 
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/com)](https://goreportcard.com/report/github.com/appleboy/com) 

This is an open source project for commonly used functions for the [Go programming language](https://golang.org/).

## Feature

* [x] Random
* [x] Array
* [x] File
* [x] Convert

## Benchmrk

```
$ go test -v -benchmem -run=^$ -bench=^Benchmark ./array/
goos: darwin
goarch: amd64
pkg: github.com/appleboy/com/array
BenchmarkArrayMap-8               200000              9735 ns/op            5654 B/op          9 allocs/op
BenchmarkArraySlice-8            2000000               663 ns/op               0 B/op          0 allocs/op
BenchmarkIn-8                    1000000              1137 ns/op            1792 B/op          1 allocs/op
BenchmarkInArray-8                100000             14337 ns/op            1632 B/op        101 allocs/op
PASS
ok      github.com/appleboy/com/array   6.780
```
