package main

import (
	"utility"
	"fmt"
	"strings"
)

func main()  {



	//for {
	//	fmt.Println(utility.RandVerifyCode(5))
	//}


	//fmt.Println(utility.GetCurrentDirectory())
	//v  := utility.Validator{}
	//v.IDMustBePositiveInteger(-1)
	//fmt.Println(v.Status)
	//fmt.Println(v.ErrMsg)


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

	fmt.Println(36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36*36)
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
			if strings.Contains(str,"出发业务流控"){
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
