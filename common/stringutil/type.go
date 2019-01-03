package stringutil

import (
	"fmt"
	"strings"
)

type KeyValuePair struct {
	Key       string
	Value     string
	Separator string
}

/**
获取键值对的String值
*/
func (c *KeyValuePair) GetLine() string {
	sep := c.Separator
	if sep == "" {
		sep = " "
	}
	return fmt.Sprintf("%s%s%s", c.Key, c.Separator, c.Value)
}

/**
该方法仅支持字符串中包含该分隔符如果没有分隔符以及分隔符并没有分割出东西来，将会返回nil
*/
func ParseLine(line, sep string) (kvp KeyValuePair) {
	if strings.Contains(line, sep) {
		tmp := strings.Split(line, sep)
		if tmp != nil && len(tmp) > 0 {
			kvp.Key = tmp[0]
			kvp.Value = tmp[1]
			kvp.Separator = sep
		}
	}
	return
}
