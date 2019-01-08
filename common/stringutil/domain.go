package stringutil

import "regexp"

func IsDomain(domainStr string) (stat bool, err error) {
	patten := "^([a-zA-Z0-9]([a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z]{2,6}$"
	reg := regexp.MustCompile(patten)
	strs := reg.FindAllString(domainStr, -1)
	if len(strs) > 0 {
		stat = true
	}

	return
}
