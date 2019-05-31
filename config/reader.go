package config

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
	"utility/common"
)

// 读取配置文件
func ReadConfig(filePath string) (params map[string]interface{}, err error) {
	params = make(map[string]interface{})
	if common.Exist(filePath) {
		suffix, _ := common.GetFileSuffix(filePath)
		if suffix == "json" {
			// 开始使用JSON的方式进行读取 ，此种方式支持不同级的参数的重名
			bytes, err2 := ioutil.ReadFile(filePath)
			if err2 != nil {
				err = err2
				return
			}
			err = json.Unmarshal(bytes, &params)
		} else if suffix == "xml" {
			// 开始使用xml的方式进行读取 ，此种方式支持不同级的参数的重名
			bytes, err2 := ioutil.ReadFile(filePath)
			if err2 != nil {
				err = err2
				return
			}
			err = xml.Unmarshal(bytes, &params)
		} else if suffix == "yml" {
			bytes, err2 := ioutil.ReadFile(filePath)
			if err2 != nil {
				err = err2
				return
			}
			err = yaml.Unmarshal(bytes, &params)
		} else {
			// 开始使用正常的方式进行逐行读取 此种方式，并不支持参数重名
			lines, err2 := common.ReadLines(filePath)
			if err2 != nil {
				err = err2
				return
			}
			// 开始逐行进行参数的读取
			for _, v := range lines {
				line := strings.TrimSpace(v)
				// 开始读取#的开始
				if !strings.HasPrefix(line, "#") {
					// 开始读取# 前面的值
					annimationIndex := strings.Index(line, "#")
					if annimationIndex != -1 {
						line = line[:annimationIndex]
					}
					// 开始
					line = strings.TrimSpace(line)
					tmps := strings.Split(line, "=")
					if len(tmps) > 1 {
						params[tmps[0]] = tmps[1]
					}
				}
			}

		}
	} else {
		err = errors.New("Path Not Exist，路径不存在")
	}
	return
}

func ReadConfigByStruct(data interface{}, filePath string) (interface{}, error) {
	if common.Exist(filePath) {
		suffix, _ := common.GetFileSuffix(filePath)
		if suffix == "json" {
			fmt.Println("开始json的解析了")
			bytes, err := ioutil.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(bytes, &data)
			fmt.Println(data)
			return data, err
		} else if suffix == "xml" {
			bytes, err := ioutil.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			err = xml.Unmarshal(bytes, &data)
			return data, err
		} else if suffix == "yml" {
			bytes, err := ioutil.ReadFile(filePath)
			if err != nil {
				return nil, err
			}
			err = yaml.Unmarshal(bytes, &data)
			return data, err
		} else {
			// 开始通过反射来进行
			// 多次循环来进行数据的读取

			return data, nil
		}
	} else {
		return nil, errors.New("文件不存在")
	}
}
