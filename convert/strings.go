package convert

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash for md5 hash string
func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
