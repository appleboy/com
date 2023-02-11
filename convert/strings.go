package convert

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"math"
	"strings"

	"github.com/appleboy/com/bytesconv"
)

// MD5Hash for md5 hash string
func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
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

	return bytesconv.BytesToStr(newstr)
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

	return bytesconv.BytesToStr(newstr)
}

// Float64ToByte convert float64 to byte
// ref: https://stackoverflow.com/questions/43693360/convert-float64-to-byte-array
func Float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

// ByteToFloat64 convert byte to float64
func ByteToFloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
