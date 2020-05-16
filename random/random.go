package random

import (
	"math/rand"
)

type (
	// Charset is string type
	Charset string
)

const (
	// Alphanumeric contain Alphabetic and Numeric
	Alphanumeric Charset = Alphabetic + Numeric
	// Alphabetic is \w+ \W
	Alphabetic Charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Numeric is number list
	Numeric Charset = "0123456789"
	// Hex is Hexadecimal
	Hex Charset = Numeric + "abcdef"
)

func randomBytes(n int) []byte {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return bytes
}

// StringWithCharset support rand string you defined
func StringWithCharset(byteLen int, charset Charset) string {
	bytes := randomBytes(byteLen)
	length := len(charset)
	for i, b := range bytes {
		bytes[i] = charset[b%byte(length)]
	}

	return string(bytes)
}

// String supply rand string
func String(length int) string {
	return StringWithCharset(length, Alphanumeric)
}
