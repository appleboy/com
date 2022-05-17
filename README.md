# Common Functions

[![Run Tests](https://github.com/appleboy/com/actions/workflows/go.yml/badge.svg)](https://github.com/appleboy/com/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/appleboy/com?status.svg)](https://godoc.org/github.com/appleboy/com)
[![codecov](https://codecov.io/gh/appleboy/com/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/com)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/com)](https://goreportcard.com/report/github.com/appleboy/com)

This is an open source project for commonly used functions for the [Go programming language](https://golang.org/).

## Feature

* [x] Random
* [x] Array
* [x] File
* [x] Convert

## Benchmrk

```sh
goos: linux
goarch: amd64
pkg: github.com/appleboy/com/array
cpu: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz
BenchmarkArrayInMap
BenchmarkArrayInMap-2           411962          8343 ns/op        5224 B/op           8 allocs/op
BenchmarkArrayInSlice
BenchmarkArrayInSlice-2        4165724           863.8 ns/op           0 B/op           0 allocs/op
BenchmarkIn
BenchmarkIn-2                  4610620           776.3 ns/op        1792 B/op           1 allocs/op
BenchmarkInArray
BenchmarkInArray-2              388922          9177 ns/op        1624 B/op         101 allocs/op
PASS
ok      github.com/appleboy/com/array    16.040s
goos: linux
goarch: amd64
pkg: github.com/appleboy/com/convert
cpu: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz
BenchmarkCountParamsOld
BenchmarkCountParamsOld-2             2575500          1400 ns/op           0 B/op           0 allocs/op
BenchmarkCountParamsNew
BenchmarkCountParamsNew-2            33431834           108.5 ns/op           0 B/op           0 allocs/op
BenchmarkBytesToStrOld01
```
