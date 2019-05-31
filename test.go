package main

import (
	"fmt"
	"utility/config"
)

type TestStruct struct {
	Name  string `json:"name"`
	Value int
}

type WenhaoExample struct {
	Wenhao string `json:"wenhao"`
}

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
	//sx,sy := robotgo.GetScreenSize()
	//robotgo.SaveCapture("F:/tmp/screen.png",0,0,sx,sy)
	//_, _, _ = img.ReadImg("F:/tmp/screen.png")
	//for k,v := range samePiexl{
	//	fmt.Println("2333")
	//	fmt.Println(k)
	//	fmt.Println(v)
	//}
	//stat, _ := stringutil.IsDomain("")
	//isip := common.IsIPv4("90.11.99.27")
	//fmt.Println(stat)
	//fmt.Println(isip)
	//out, err := command.RunCmd("systemctl", "status", "sshd")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(out)
	//lines := strings.Split(out, "\n")
	//for _, line := range lines {
	//	if strings.Contains(line, "active") {
	//		fmt.Println(line)
	//	}
	//}
	//compress.DeCompress("/root/mysql-5.6.41.tar.gz", "/root/")
	//var params map[string]string
	//params = make(map[string]string)
	//params["abnudsonfwe"] = "23333"
	//params["aadnaowenf"] = "43333"
	//sign := encryption.Sign(params,"2333")
	//fmt.Println(sign)
	//structT := TestStruct{
	//	"你好",
	//	233333,
	//}
	//mapT := common.ConvertStructToMapString(structT)
	//fmt.Println(mapT)
	//mapT2 := common.ConvertStructToMapStringWithTagName(structT, "json")
	//fmt.Println(mapT2)
	//fmt.Println(mapT["Value"])
	//fmt.Println(mapT2["Value"])
	//s := "2379h#fao9hj83#"
	//fmt.Println(strings.Index(s, "#"))
	//fmt.Println(s[:5])
	//params, err := config.ReadConfig("./test.json")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(params)

	var wenhao WenhaoExample
	data, err := config.ReadConfigByStruct(wenhao, "./test.json")
	fmt.Println(err)
	fmt.Println(data)

}
