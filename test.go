package main

import (
	"fmt"
	"utility/encryption"
)

func main() {
	//d := downloader.Downloader{}
	//d.SavePath = "D:/"
	//d.TargetUrl = "http://download.oracle.com/otn-pub/java/jdk/10.0.2+13/19aef61b38124481863b1413dce1855f/jdk-10.0.2_linux-x64_bin.tar.gz"
	//d.Cookie = "oraclelicense=accept-securebackup-cookie"
	//d.FullDownlod()
	//system.DetectOSType()
	//system.IsWindows()
	//system.IsLinux()
	//system.IsMacos()
	//result,err := command.RunCmd("netstat","-ntl")
	//if err != nil{
	//	logs.Error(err)
	//}
	//logs.Info(result)
	var params map[string]string
	params = make(map[string]string)
	params["23333"] = "42njoahufd"
	params["nowbnfewn"] = "23ui098hinfan"
	formatedStr, _ := encryption.FormatParams(params, "&")
	fmt.Println(formatedStr)

}
