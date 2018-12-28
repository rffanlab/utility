package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Hmac(data, secret string) (enStr string) {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	enStr = hex.EncodeToString(h.Sum(nil))
	return
}
