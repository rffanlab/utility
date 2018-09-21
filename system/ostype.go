package system

import (
	"fmt"
	"runtime"
	"strings"
	"utility/common"
)

/**
windows
linux
darwin 苹果
 */

 /**
 系统类型
  */
type OSType struct {
	Type string
	Version string
	Distribution string
	Arch string
}

/**
解析LinuxRelease文件
 */
func ParseOSRealeaseFile()(map[string]string,error)  {
	var params map[string]string
	params = make(map[string]string)
	lines,err :=common.ReadLines("/etc/os-release")
	if err != nil {
		return nil,err
	}
	for _,value := range  lines {
		strs := strings.Split(value,"=")
		if len(strs)>1 {
			params[strs[0]]=strs[1]
		}
	}
	return params,nil
}


/**
检测系统类型
目前Windows只能检测到是window不能检测更深入的版本，例如win7 win10等
Linux能够检测到发行版本，原理是通过Linux的os-release文件来进行确认的。
macos没办法确认
 */
func DetectOSType() (OSType,error) {
	var ostype OSType
	ostype.Arch = runtime.GOARCH
	ostype.Type = runtime.GOOS
	if ostype.Type == "linux"{
		params,err := ParseOSRealeaseFile()
		if err != nil {
			return ostype,err
		}
		ostype.Distribution = strings.Replace(params["ID"],"\"","",-1)
		ostype.Version = strings.Replace(params["VERSION_ID"],"\"","\"",-1)
	}
	return ostype,nil
}

/**
获取系统线程数量
 */
func GetThreadNum() int {
	return runtime.NumCPU()
}

func IsWindows() bool {
	ostype,err:= DetectOSType()
	if err != nil {
		return false
	}
	fmt.Print(ostype.Type)
	if ostype.Type == "windows" {
		return true
	}
	return false
}

func IsLinux() bool {
	ostype,err:= DetectOSType()
	if err != nil {
		return false
	}
	fmt.Print(ostype.Type)
	if ostype.Type == "linux" {
		return true
	}
	return false
}

func IsMacos() bool {
	ostype,err:= DetectOSType()
	if err != nil {
		return false
	}
	fmt.Print(ostype.Type)
	if ostype.Type == "darwin" {
		return true
	}
	return false
}


