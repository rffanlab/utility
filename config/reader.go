package config

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"utility/common"
)

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

/**
我们约定俗称以 . 作为分隔符
*/
func GetParamByKey(key string, params map[string]interface{}) (result string) {
	pointIndex := strings.Index(key, ".")
	strs := strings.Split(key, ".")
	if len(strs) <= 0 {
		return
	}
	for k, v := range params {
		fmt.Println(k, v)
		if k == strs[0] {
			if len(strs) > 1 {
				r := reflect.ValueOf(v)
				if r.Kind() == reflect.Map {
					p := ConvertInterfaceToMap(v)
					return GetParamByKey(key[pointIndex+1:], p)
				} else {

					return
				}
			} else {
				r, err := ConvertInterfaceToString(v)
				result = r
				if err != nil {
					result = ""
				}
				return
			}
		} else {

		}
	}
	return
}

func ConvertInterfaceToMap(i interface{}) (result map[string]interface{}) {
	result = make(map[string]interface{})
	r := reflect.ValueOf(i)
	if r.Kind() == reflect.Map {
		for _, key := range r.MapKeys() {
			valueI := key.Interface()
			value, err := ConvertInterfaceToString(valueI)
			if err != nil {
				fmt.Println(err.Error())
			}
			result[value] = r.MapIndex(key).Interface()
		}
	}
	return result
}

/**
由于不会奇葩到用其他东西，例如数值尤其是小数点做值，因此只做这几点的转换
*/
func ConvertInterfaceToString(i interface{}) (result string, err error) {
	r := reflect.ValueOf(i)
	if r.Kind() == reflect.String {
		result = i.(string)
	} else if r.Kind() == reflect.Int {
		value := i.(int)
		result = strconv.Itoa(value)
	} else if r.Kind() == reflect.Float32 {
		result = fmt.Sprintf("%f", i)
	} else if r.Kind() == reflect.Float64 {
		result = fmt.Sprintf("%f", i)
	} else {
		err = errors.New(fmt.Sprintf("Unsupport kind to convert %r", r.Kind()))
	}
	return
}
