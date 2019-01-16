package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
)

func Sha256Hmac(data, secret string) (enStr string) {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	enStr = hex.EncodeToString(h.Sum(nil))
	return
}

/**
按照ASCII码排序Map并且k=v&的形式进行连接，然后使用secret来进行加密
*/
func Sign(params map[string]string, secret string) (sign string) {
	var key []string
	var paramStr string
	for k, _ := range params {
		key = append(key, k)
	}
	sort.Strings(key)
	fmt.Println(key)
	for i, v := range key {
		if i == len(key)-1 {
			paramStr += v + "=" + params[v]
		} else {
			paramStr += v + "=" + params[v] + "&"
		}
	}
	return Sha256Hmac(paramStr, secret)
}
