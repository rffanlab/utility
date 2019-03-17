package validate

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	VALIDATOR = "validator"

	REQUIRED = "required" //
	MIN      = "min"      // int值的最小值
	MAX      = "max"      // int值的最大值
	RANGE    = "range"    // int 值的范围
	MINSIZE  = "minsize"  // 切片的最小长度
	MAXSIZE  = "maxsize"  // 切片的最大长度
	LENGTH   = "length"   // 有效长度
	ISNUMBER = "isnumber" // 是一个有效的数值
	MATCH    = "match"    // 正则验证
	ISEMAIL  = "isemail"  // Email
	ISIPV4   = "isipv4"   // 是IPv4
	ISMOBILE = "ismobile" // 是手机号码
	ISPHONE  = "isphone"  // 是电话号码
)

/**
用来验证结构体,标签为validate 以";"为分隔符
required();range(1,100);size(100);match()
*/
func Validate(toValidate interface{}) (err error) {
	toValidateT := reflect.TypeOf(toValidate)
	toValidateV := reflect.ValueOf(toValidate)
	// 做一下结构体的验证
	switch {
	case toValidateT.Kind() == reflect.Struct:
		break
	case toValidateT.Kind() == reflect.Ptr && toValidateT.Elem().Kind() == reflect.Struct:
		toValidateT = toValidateT.Elem()
		toValidateV = toValidateV.Elem()
		break
	default:
		err = fmt.Errorf("%v 必须是一个结构体", toValidate)
		return
	}

	// 开始for循环进行每个标签的判断
	for i := 0; i < toValidateT.NumField(); i++ {
		// 获取tag标记
		value := toValidateV.Field(i).Interface()
		validateParam := GetTags(toValidateT.Field(i))
		for _, v := range validateParam {
			CheckAll(value, v)
		}
	}
	return
}

// 获取标记
func GetTags(tagMark reflect.StructField) (tagParams []string) {
	tag := tagMark.Tag.Get(VALIDATOR)
	tagParams = strings.Split(tag, ";")
	return
}

// 开始check检查
func CheckAll(toCheck interface{}, rule string) (stat bool) {
	if strings.HasPrefix(rule, REQUIRED) {
		stat = CheckRequired(toCheck)
	} else if strings.HasPrefix(rule, MIN) {

	} else if strings.HasPrefix(rule, MAX) {
	} else if strings.HasPrefix(rule, RANGE) {
	} else if strings.HasPrefix(rule, MINSIZE) {
	} else if strings.HasPrefix(rule, MAXSIZE) {
	} else if strings.HasPrefix(rule, LENGTH) {
	} else if strings.HasPrefix(rule, ISNUMBER) {
	} else if strings.HasPrefix(rule, MATCH) {
	} else if strings.HasPrefix(rule, ISEMAIL) {
	} else if strings.HasPrefix(rule, ISIPV4) {
	} else if strings.HasPrefix(rule, ISMOBILE) {
	} else if strings.HasPrefix(rule, ISPHONE) {

	}
	return
}

// 检查是否required
func CheckRequired(toCheck interface{}) (stat bool) {
	if toCheck == nil {
		return
	}
	if str, ok := toCheck.(string); ok {
		return len(strings.TrimSpace(str)) > 0
	}
	if _, ok := toCheck.(bool); ok {
		return true
	}
	if i, ok := toCheck.(int); ok {
		return i != 0
	}
	if i, ok := toCheck.(uint); ok {
		return i != 0
	}
	if i, ok := toCheck.(int8); ok {
		return i != 0
	}
	if i, ok := toCheck.(uint8); ok {
		return i != 0
	}
	if i, ok := toCheck.(int16); ok {
		return i != 0
	}
	if i, ok := toCheck.(uint16); ok {
		return i != 0
	}
	if i, ok := toCheck.(uint32); ok {
		return i != 0
	}
	if i, ok := toCheck.(int32); ok {
		return i != 0
	}
	if i, ok := toCheck.(int64); ok {
		return i != 0
	}
	if i, ok := toCheck.(uint64); ok {
		return i != 0
	}
	if t, ok := toCheck.(time.Time); ok {
		return !t.IsZero()
	}
	v := reflect.ValueOf(toCheck)
	if v.Kind() == reflect.Slice {
		return v.Len() > 0
	}
	return true
}
