package convert

import (
	"bytes"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"unsafe"
)

var (
	strColon = []byte(":")
	strStar  = []byte("*")
)

func countParamsOld(path string) uint16 {
	var n uint
	for i := range []byte(path) {
		switch path[i] {
		case ':', '*':
			n++
		}
	}
	return uint16(n)
}

func countParamsNew(path string) uint16 {
	var n uint
	s := StrToBytes(path)
	n += uint(bytes.Count(s, strColon))
	n += uint(bytes.Count(s, strStar))
	return uint16(n)
}

var foo = strings.Repeat("/:param", 256)

func BenchmarkCountParamsOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countParamsOld(foo)
	}
}

func BenchmarkCountParamsNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countParamsNew(foo)
	}
}

var s = strings.Repeat("s", 1024)
var stringSink string
var byteSink []byte

func BenchmarkBytesToStrOld01(b *testing.B) {
	b.SetBytes(int64(len(s)))
	b.ReportAllocs()
	b.ResetTimer()

	byt := []byte(s)
	for i := 0; i < b.N; i++ {
		v := string(byt)
		stringSink = v
	}
}

func bytesToString(bytes []byte) (s string) {
	if len(bytes) == 0 {
		return s
	}
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sh.Data = uintptr(unsafe.Pointer(&bytes[0]))
	sh.Len = len(bytes)
	return s
}

func BenchmarkBytesToStrOld2(b *testing.B) {
	b.SetBytes(int64(len(s)))
	b.ReportAllocs()
	b.ResetTimer()

	byt := []byte(s)
	for i := 0; i < b.N; i++ {
		v := bytesToString(byt)
		stringSink = v
	}
}

func BenchmarkBytesToStrNew(b *testing.B) {
	b.SetBytes(int64(len(s)))
	b.ReportAllocs()
	b.ResetTimer()

	byt := []byte(s)
	for i := 0; i < b.N; i++ {
		v := BytesToStr(byt)
		stringSink = v
	}
}

// report: https://cloud.drone.io/appleboy/com/56/1/2
//
// BenchmarkBytesToStrOld01-48            20000000            337 ns/op    3038.42 MB/s      1024 B/op         1 allocs/op
// BenchmarkBytesToStrOld2-48           2000000000           2.15 ns/op  476609.23 MB/s         0 B/op         0 allocs/op
// BenchmarkBytesToStrNew-48            5000000000           1.07 ns/op  953243.02 MB/s         0 B/op         0 allocs/op

func BenchmarkStr2BytesOld01(b *testing.B) {
	b.SetBytes(int64(len(s)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := []byte(s)
		byteSink = v
	}
}

// strToBytesOld converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func strToBytesOld(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

func BenchmarkStr2BytesOld02(b *testing.B) {
	b.SetBytes(int64(len(s)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := strToBytesOld(s)
		byteSink = v
	}
}

func BenchmarkStr2BytesNew(b *testing.B) {
	b.SetBytes(int64(len(s)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v := StrToBytes(s)
		byteSink = v
	}
}

// Report: https://cloud.drone.io/appleboy/com/54/1/2
//
// BenchmarkStr2BytesOld01-48            20000000           340 ns/op   3010.68 MB/s     1024 B/op        1 allocs/op
// BenchmarkStr2BytesOld02-48          2000000000          3.60 ns/op 284659.92 MB/s        0 B/op        0 allocs/op
// BenchmarkStr2BytesNew-48            2000000000          2.15 ns/op 476653.64 MB/s        0 B/op        0 allocs/op

func BenchmarkConvertOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := []byte(s)
		_ = string(k)
	}
}

func BenchmarkConvertNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := StrToBytes(s)
		_ = BytesToStr(k)
	}
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func BenchmarkSnakeCasedNameRegex(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	s := strings.Repeat("FooBar", 32)
	for i := 0; i < b.N; i++ {
		_ = toSnakeCase(s)
	}
}

func BenchmarkSnakeCasedNameOld(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	s := strings.Repeat("FooBar", 32)
	for i := 0; i < b.N; i++ {
		_ = snakeCasedNameOld(s)
	}
}

func BenchmarkSnakeCasedNameNew(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	s := strings.Repeat("FooBar", 32)
	for i := 0; i < b.N; i++ {
		_ = SnakeCasedName(s)
	}
}

func BenchmarkTitleCasedName(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	s := strings.Repeat("foo_bar", 32)
	for i := 0; i < b.N; i++ {
		_ = TitleCasedName(s)
	}
}
