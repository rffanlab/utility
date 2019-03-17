package encryption

import (
	"fmt"
	"sort"
)

func FormatParams(params map[string]string, sep string) (formatedStr string, err error) {
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		if i == (len(keys) - 1) {
			formatedStr += fmt.Sprintf("%s=%s", k, params[k])
		} else {
			formatedStr += fmt.Sprintf("%s=%s%s", k, params[k], sep)
		}
	}
	return
}

/**
生成签名 sha256Hmac
*/
func MakeSign(params map[string]string, secret string) (result string, err error) {
	formatStr, err := FormatParams(params, "&")
	if err != nil {
		return
	}
	result = Sha256Hmac(formatStr, secret)
	return
}

/**
  生成签名 md5
*/
