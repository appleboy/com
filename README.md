# Common Functions

[![GoDoc](https://godoc.org/github.com/appleboy/com?status.svg)](https://godoc.org/github.com/appleboy/com) [![Build Status](http://drone.wu-boy.com/api/badges/appleboy/com/status.svg)](http://drone.wu-boy.com/appleboy/com) [![codecov](https://codecov.io/gh/appleboy/com/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/com) [![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/com)](https://goreportcard.com/report/github.com/appleboy/com) [![Sourcegraph](https://sourcegraph.com/github.com/appleboy/com/-/badge.svg)](https://sourcegraph.com/github.com/appleboy/com?badge)

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
