package convert

import (
	"bytes"
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
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func BenchmarkStr2BytesNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StrToBytes(s)
	}
}

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
