package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os/exec"
	"strings"
	"timer/utility"
)

func main()  {

	cmd := exec.Command("D:\\nginx-1.15.8\\nginx-1.15.8\\nginx.exe","-t")
	cmd.Dir="D:\\nginx-1.15.8\\nginx-1.15.8"
	stdout,err := cmd.StdoutPipe()
	stderr,err := cmd.StderrPipe()
	if err != nil {
		fmt.Print(err)
	}
	defer stderr.Close()
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	opBytes,err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(opBytes))
	opBytes,err = ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("-----------------开始打印错误-------------")
	fmt.Println(string(opBytes))

	// 短信发送在这里
	//SendSMSExample()

	//MakrRandUserkey()
	//fmt.Println(rand.Intn(100))

	//md51,_ := utility.LargeFilemd5("D:/codes/go/src/jingcheng/store/audi.mp4")
	//md52,_ := utility.Filemd5("D:/codes/go/src/jingcheng/store/audi.mp4")
	//
	//fmt.Println(md51)
	//fmt.Println(md52)

	//fmt.Println(utility.IsRelativePath("2333"))

	//fmt.Println(36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36)
	//fmt.Println(UrlEncodeString("fileid=1"))
	//WeChatAccessToken()
	//fmt.Println(utility.GetTimestamp())
	//fmt.Println(utility.CompareTimestampNow(1406418130))

	//fmt.Println(utility.RandStr(15))
	//fmt.Println(utility.CompareStrToSaltEncryptedStr("face4337197","cc6c3060b062b59d2df8d3cb382069f0x5dy"))


	//fmt.Println(utility.Today())
	//t,_ := utility.TransferDateFromStringToTime("2017年10月12日 下午2点20")
	//tm := utility.TimeToTimestamp(t)
	//fmt.Println(utility.CompareTimestampNow(tm))
	//var t2 time.Time
	//var t3 time.Time
	//fmt.Println(t2)
	//fmt.Println(t2 == t3)
	//fmt.Println(utility.TransferDateFromStringToTime("2017-10-13 13:38"))
	//fmt.Println(utility.TimeNowForSecond())


	//ipinfo,err := utility.GetIPinfo("158.69.251.119")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(ipinfo)
	//dirs,err := ioutil.ReadDir("D:/codes/gopath/src/utility/")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//for _,v := range dirs{
	//	fmt.Println(v.Name())
	//}


	// 开始做获取Nginx的最新稳定版本的配置
	//nginx := nginx2.Nginx{}
	//fmt.Println(nginx.GetDownloadLink())
	//downloadLink := nginx.GetDownloadLink()
	//d := downloader.Downloader{
	//	TargetUrl:downloadLink,
	//}
	//d.LiteDown()


	// 开始Python相关工具的测试
	//pythonHelper := python.Python{}
	//linklist,err := pythonHelper.GetVersionList()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _,value := range linklist{
	//	fmt.Println(pythonHelper.MakeDownloadLinkByVer(value))
	//}

	//err := utility.AppendLine("\n23333","./testfile")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//for i:=0;i<4;i++{
	//	system.Ping("www.163.com")
	//	time.Sleep(time.Duration(2)*time.Second)
	//}



}

// 发送短信 模板，由于重写了方法，所以方法的使用有点类似JS的链式写法。
// 终结为Send()方法，该方法会检测前面的参数是否设置正确。
// 如果设置正确，则直接发送数据
func SendSMSExample() {
	sms := utility.SMSClient{}
	result,err := sms.SetDefaultClient("","").
		SetRecieveMobile("").
		SetSMSFreeSignName("").
		SetSMSType("normal").
		SetSMSTemplateCode("").
		SetParams("{\"name\":\"\",\"appname\":\"\",\"code\":\"\"}").
		Send()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	if _,ok := result["error_response"];ok{
		if str,ok := result["error_response"].(string);ok{
			if strings.Contains(str,"触发业务流控"){
				fmt.Println(str)
			}else {
				fmt.Println(str)
			}
		}
	}else {
		fmt.Println("发送成功")
	}
}

func MakrRandUserkey() {
	for  {
		fmt.Println(utility.RandVerifyCode(6))
	}
	//fmt.Println(utility.MakeUserkey())
}

func UrlEncodeString(theStr string) string {
	return url.QueryEscape(theStr)
}

func WeChatAccessToken() {
	wc := utility.Wechat{
		Appkey:"",
		AppSecret:"",
	}
	ac,err := wc.GetAccessToken()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(ac)
	stat,err := wc.IsSubscribe("",ac)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(stat)
}