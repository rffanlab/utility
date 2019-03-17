package common

import (
	"fmt"
	"reflect"
)

/**
转换的类型为map[string]string
根据结构体的字段名来进行结构体的解析
*/
func ConvertStructToMapString(theStruct interface{}) map[string]string {
	v := reflect.ValueOf(theStruct)
	rmap := make(map[string]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		rmap[v.Type().Field(i).Name] = fmt.Sprintf("%v", v.Field(i).Interface())
	}
	return rmap
}

/**
转换的类型为map[string]string
根据标签内的名字来进行标签名和Map的转换如果标签没有，则使用字段名作为标签进行转换
如果tag的入参为默认的空值，则不做任何操作直接以字段名进行Map的解析
*/
func ConvertStructToMapStringWithTagName(theStruct interface{}, tag string) map[string]string {
	if tag == "" {
		return ConvertStructToMapString(theStruct)
	}
	v := reflect.ValueOf(theStruct)
	rmap := make(map[string]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		key := ""
		if v.Type().Field(i).Tag.Get(tag) == "" {
			key = v.Type().Field(i).Name
		} else {
			key = v.Type().Field(i).Tag.Get(tag)
		}
		rmap[key] = fmt.Sprintf("%v", v.Field(i).Interface())
	}
	return rmap
}

func ConvertMapToStruct(sourceMap interface{}, targetStruct interface{}) {

}
