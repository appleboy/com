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
cpu: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
BenchmarkArrayInMap
BenchmarkArrayInMap-2     	  360637	      8577 ns/op	    5224 B/op	       8 allocs/op
BenchmarkArrayInSlice
BenchmarkArrayInSlice-2   	 8538693	       426.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkIn
BenchmarkIn-2             	 3864145	       936.2 ns/op	    1792 B/op	       1 allocs/op
BenchmarkInArray
BenchmarkInArray-2        	  304135	     11783 ns/op	    1624 B/op	     101 allocs/op
PASS
ok  	github.com/appleboy/com/array	15.543s
goos: linux
goarch: amd64
pkg: github.com/appleboy/com/convert
cpu: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
BenchmarkCountParamsOld
BenchmarkCountParamsOld-2        	  886718	      3453 ns/op	       0 B/op	       0 allocs/op
BenchmarkCountParamsNew
BenchmarkCountParamsNew-2        	30483649	       116.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytesToStrOld01
BenchmarkBytesToStrOld01-2       	17070554	       202.1 ns/op	5067.66 MB/s	    1024 B/op	       1 allocs/op
BenchmarkBytesToStrOld2
BenchmarkBytesToStrOld2-2        	1000000000	         1.694 ns/op	604583.42 MB/s	       0 B/op	       0 allocs/op
BenchmarkBytesToStrNew
BenchmarkBytesToStrNew-2         	1000000000	         0.7197 ns/op	1422849.47 MB/s	       0 B/op	       0 allocs/op
BenchmarkStr2BytesOldLong
BenchmarkStr2BytesOldLong-2      	17513360	       202.7 ns/op	5052.15 MB/s	    1024 B/op	       1 allocs/op
BenchmarkStr2BytesNewSLong
BenchmarkStr2BytesNewSLong-2     	1000000000	         2.537 ns/op	403604.72 MB/s	       0 B/op	       0 allocs/op
BenchmarkStr2BytesOldShort
BenchmarkStr2BytesOldShort-2     	15229582	       202.6 ns/op	   4.94 MB/s	    1024 B/op	       1 allocs/op
BenchmarkStr2BytesNewShort
BenchmarkStr2BytesNewShort-2     	1000000000	         2.525 ns/op	 396.12 MB/s	       0 B/op	       0 allocs/op
BenchmarkConvertOld
BenchmarkConvertOld-2            	 9095806	       402.2 ns/op	    2048 B/op	       2 allocs/op
BenchmarkConvertNew
BenchmarkConvertNew-2            	1000000000	         0.3389 ns/op	       0 B/op	       0 allocs/op
BenchmarkSnakeCasedNameRegex
BenchmarkSnakeCasedNameRegex-2   	   98422	     35881 ns/op	    4802 B/op	      80 allocs/op

```
