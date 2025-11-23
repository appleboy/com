package convert

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"math"
	"strings"
)

/*
MD5Hash computes the SHA-256 hash of the input string and returns a 64-character hexadecimal string.
- Useful for data validation, generating unique identifiers, etc.
- Note: Function name retained for backward compatibility, but now uses SHA-256 for improved security.
*/
func MD5Hash(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

/*
SnakeCasedName converts a string to snake_case format and supports Unicode characters.
- Each uppercase English letter (A-Z) is converted to lowercase and prefixed with an underscore if not at the start.
- Only English letters are affected; other Unicode characters (e.g., Chinese) are preserved as-is.
- Example: FooBar -> foo_bar, 你好World -> 你好_world
*/
func SnakeCasedName(name string) string {
	var b strings.Builder
	b.Grow(len(name) + 1)
	for i, r := range name {
		if isUpper := 'A' <= r && r <= 'Z'; isUpper {
			if i > 0 {
				b.WriteRune('_')
			}
			r += 'a' - 'A'
		}
		b.WriteRune(r)
	}
	return b.String()
}

/*
TitleCasedName converts a snake_case string to TitleCase format and supports Unicode characters.
- Each English letter following an underscore is capitalized, and underscores are removed.
- Only English letters are affected; other Unicode characters (e.g., Chinese) are preserved as-is.
- Example: foo_bar -> FooBar, hello_世界 -> Hello世界
*/
func TitleCasedName(name string) string {
	var b strings.Builder
	b.Grow(len(name))
	upNextChar := true

	for _, r := range name {
		switch {
		case upNextChar:
			upNextChar = false
			if 'a' <= r && r <= 'z' {
				r -= 'a' - 'A'
			}
		case r == '_':
			upNextChar = true
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

/*
Float64ToByte converts a float64 value to an 8-byte slice in BigEndian order.
- Useful for binary serialization and network transmission.
- Reference: https://stackoverflow.com/questions/43693360/convert-float64-to-byte-array
*/
func Float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

/*
ByteToFloat64 converts an 8-byte slice in BigEndian order back to a float64 value.
- Panics if the input length is not 8.
- Useful for binary deserialization and network data parsing.
*/
func ByteToFloat64(bytes []byte) float64 {
	if len(bytes) != 8 {
		panic("ByteToFloat64: input length must be 8 bytes")
	}
	bits := binary.BigEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
