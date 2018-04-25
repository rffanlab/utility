package oneprovider

import "github.com/pkg/errors"

type OneProvider struct {
	ApiKey string `json:"api_key"`
	ClientKey string `json:"client_key"`
}

func (c *OneProvider) Check() error {
	if c.ApiKey == "" || c.ClientKey == "" {
		return errors.New("没有设置APIkey 或者没有设置ClientKey")
	}
	return nil
}