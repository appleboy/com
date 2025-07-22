package random

import (
	"crypto/rand"
	mathrand "math/rand"
	"sync"
	"time"

	"github.com/appleboy/com/bytesconv"
)

type (
	// Charset represents a set of characters to use for random string generation.
	Charset string
)

const (
	// Alphanumeric contains all alphabetic and numeric characters (A-Z, a-z, 0-9).
	Alphanumeric Charset = Alphabetic + Numeric
	// Alphabetic contains all uppercase and lowercase English letters (A-Z, a-z).
	Alphabetic Charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Numeric contains all decimal digits (0-9).
	Numeric Charset = "0123456789"
	// Hex contains all hexadecimal digits (0-9, a-f).
	Hex Charset = Numeric + "abcdef"
)

/*
randomBytes returns a slice of n cryptographically secure random bytes.
It returns an error if the system's secure random number generator fails.
*/
func randomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

/*
StringWithCharset returns a cryptographically secure random string of the given byte length,
using the provided charset. The randomness is suitable for security-sensitive use cases.
Returns an error if the system's secure random number generator fails.
*/
func StringWithCharset(byteLen int, charset Charset) (string, error) {
	bytes, err := randomBytes(byteLen)
	if err != nil {
		return "", err
	}
	length := len(charset)
	for i, b := range bytes {
		bytes[i] = charset[b%byte(length)]
	}

	return bytesconv.BytesToStr(bytes), nil
}

// ref: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	src = mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	mu  sync.Mutex
)

/*
randStringBytesMaskImprSrcUnsafe returns a random string of the given length using math/rand.
This function is optimized for speed but is NOT suitable for cryptographic or security-sensitive use cases.
*/
func randStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	mu.Lock()
	defer mu.Unlock()
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

	return bytesconv.BytesToStr(b)
}

/*
String returns a random string of the given length using a fast, non-cryptographically secure generator.
For security-sensitive use cases, use StringWithCharset instead.
*/
func String(length int) string {
	return randStringBytesMaskImprSrcUnsafe(length)
}

/*
RandomString returns a random string of the given length using the provided charset.
If secure is true, it uses a cryptographically secure random generator (returns error on failure).
If secure is false, it uses a fast, non-cryptographically secure generator (never returns error).
If charset is empty, Alphanumeric is used.
*/
func RandomString(length int, charset Charset, secure bool) (string, error) {
	if charset == "" {
		charset = Alphanumeric
	}
	if secure {
		return StringWithCharset(length, charset)
	}
	// Fast, insecure method ignores charset and always uses letterBytes
	return randStringBytesMaskImprSrcUnsafe(length), nil
}
