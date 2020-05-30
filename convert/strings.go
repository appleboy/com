package convert

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"strings"
	"unsafe"
)

// MD5Hash for md5 hash string
func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// BytesToStr converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StrToBytes turns a string into a []byte with 0 MemAllocs and 0 MemBytes.
// This is an unsafe operation and will lead to problems if the underlying bytes
// are changed.
func StrToBytes(s string) (b []byte) {
	if len(s) == 0 {
		return b
	}
	const max = 0x7fff0000 // 2147418112
	if len(s) > max {
		panic("string too large")
	}
	bytes := (*[max]byte)(
		unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data),
	)
	return bytes[:len(s):len(s)]
}

// SnakeCasedName convert String into Snake Case
// ex: FooBar -> foo_bar
func snakeCasedNameOld(name string) string {
	newstr := make([]rune, 0)
	for idx, chr := range name {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if idx > 0 {
				newstr = append(newstr, '_')
			}
			chr -= ('A' - 'a')
		}
		newstr = append(newstr, chr)
	}

	return string(newstr)
}

// SnakeCasedName convert String into Snake Case
// ex: FooBar -> foo_bar
func SnakeCasedName(name string) string {
	newstr := make([]byte, 0, len(name)+1)
	for i := 0; i < len(name); i++ {
		c := name[i]
		if isUpper := 'A' <= c && c <= 'Z'; isUpper {
			if i > 0 {
				newstr = append(newstr, '_')
			}
			c += 'a' - 'A'
		}
		newstr = append(newstr, c)
	}

	return BytesToStr(newstr)
}

// TitleCasedName convert String into title cased
// ex: foo_bar -> FooBar
func TitleCasedName(name string) string {
	newstr := make([]byte, 0, len(name))
	upNextChar := true

	name = strings.ToLower(name)

	for i := 0; i < len(name); i++ {
		c := name[i]
		switch {
		case upNextChar:
			upNextChar = false
			if 'a' <= c && c <= 'z' {
				c -= 'a' - 'A'
			}
		case c == '_':
			upNextChar = true
			continue
		}

		newstr = append(newstr, c)
	}

	return BytesToStr(newstr)
}
