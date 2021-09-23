package random

import (
	"math/rand"
	"time"
	"unsafe"
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

// ref: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// String supply rand string
func String(length int) string {
	return randStringBytesMaskImprSrcUnsafe(length)
}
