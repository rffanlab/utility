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

type OSType struct {
	Type string
	Version string
	Distribution string
	Arch string
}

func ParseOSRealeaseFile()(map[string]string,error)  {
	var params map[string]string
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


fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)


	return ostype,nil
}

func GetThreadNum() int {
	return runtime.NumCPU()
}


