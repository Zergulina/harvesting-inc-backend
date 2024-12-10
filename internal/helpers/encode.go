package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func EncodeSha256(data string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
