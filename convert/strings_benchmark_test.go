package convert

import (
	"bytes"
	"regexp"
	"strings"
	"testing"
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

func BenchmarkBytesToStrOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string([]byte(s))
	}
}

func BenchmarkBytesToStrNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BytesToStr([]byte(s))
	}
}

func BenchmarkStr2BytesOld(b *testing.B) {
	b.SetBytes(int64(len(s)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := []byte(s)
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

// BenchmarkStr2BytesOld-8          6584626                  181 ns/op       5655.81 MB/s        1024 B/op          1 allocs/op
// BenchmarkStr2BytesNew-8         646368142                1.89 ns/op     541225.49 MB/s           0 B/op          0 allocs/op

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
