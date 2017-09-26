package utility

import (
	"image"
	"os"
	"github.com/astaxie/beego/logs"
	"image/png"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode"
	"sort"
	"crypto/sha1"
	"io"
	"strings"
	"fmt"
	"go-requests"
	"io/ioutil"
	"encoding/json"
	"errors"
	"net/url"
)

type Wechat struct {
	Appkey string
	AppSecret string
	Token string
}


type WechatResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid string `json:"openid"`
	Scope string `json:"scope"`
}

/******************************************
*           微信验证和处理相关            *
*                                         *
*******************************************/



func (c *Wechat) CheckConfigSet() (error) {
	if c.Appkey == "" {
		return errors.New("请设置微信APPKEY")
	}
	if c.AppSecret == ""{
		return errors.New("请设置微信APPSecret")
	}
	if c.Token == "" {
		return errors.New("请设置微信TOKEN")
	}
	return nil
}



//  方法：微信验证接口
/*
*   传入参数：
*   @Param:signatur Type:string
*   @Param:timestamp Type:string
*   @Param:nonce Type:string
*   返回参数：
*   @Param:bool Explain:返回确认是否正确
*/
func (c *Wechat) Verify(signature,timestamp,nonce string) (bool) {
	sign := c.MakeSignatureWith(timestamp,nonce)
	if signature == sign{
		return true
	}else {
		return false
	}
}

func (c *Wechat) MakeSignatureWith(timestamp,nonce string) string {
	err := c.CheckConfigSet()
	if err != nil{
		logs.Error(err)
	}
	sl := []string{c.Token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

//  方法：发送消息接口
/*
*   传入参数：
*   @Param: Type:
*   @Param: Type:
*   返回参数：
*   @Param: Explain:
*   @Param: Explain:
*/
func (c *Wechat) SendWechatMessageToClient() {

}

//  方法：自动回复接口
/*
*   传入参数：
*   @Param: Type:
*   @Param: Type:
*   返回参数：
*   @Param: Explain:
*   @Param: Explain:
*/
func (c *Wechat) AutoResponse() {

}

//// 方法：生成验证二维码
///*
//*  传入参数：
//*  @Param:codeName Type:string The File it store
//*  @Param:returnUrl Type:Not Url Encoded string
//*  @Param:msg Type:string 用来设置state的value（可能设置的是userkey或者其他）
//*  返回参数：
//*  @Param: Type:
//*  @Param: Type:
//*/
//func (c *Wechat) MakeVerifyQRCode(codeName,returnUrl,msg string) (string,error) {
//	currentDirectory := GetCurrentDirectory()
//	fullPath := currentDirectory + beego.AppConfig.String("qrcodepath")+"/"+codeName
//	encodedUrl := strings.ToLower(url.QueryEscape(returnUrl))
//	logs.Info("编码后的url是：",encodedUrl)
//	// TODO redirect_url 错了,如果是服务号则可通过，订阅号不支持此权限
//	urlmsg := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd328b9395790be6e&redirect_uri="+encodedUrl+"&response_type=code&scope=snsapi_userinfo&state="+msg+"#wechat_redirect"
//	c.MakeQRCode(urlmsg,fullPath)
//	return beego.AppConfig.String("qrcodepath")+"/"+codeName,nil
//}
func (c *Wechat) MakeVerifyQrCode(codePath, returnUrl, msg string) (string, error) {
	encodedUrl := strings.ToLower(url.QueryEscape(returnUrl))
	logs.Info("编码后的url是：",encodedUrl)
	urlmsg := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wxd328b9395790be6e&redirect_uri="+encodedUrl+"&response_type=code&scope=snsapi_userinfo&state="+msg+"#wechat_redirect"
	err := c.MakeQRCode(urlmsg,codePath)
	if err != nil{
		logs.Info("生成二维码时出错，请查看下面的错误内容")
		logs.Error(err)
	}
	return codePath,nil
}


func (c *Wechat) GetUserWechatInfo(withCode string) (WechatResponse,error) {
	var wr WechatResponse
	err := c.CheckConfigSet()
	if err != nil{
		return wr,err
	}
	getAccess_tokenUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid="+c.Appkey+"&secret="+c.AppSecret+"&code="+ withCode + "&grant_type=authorization_code"
	r := go_requests.Requests{}
	theio,err := r.Get(getAccess_tokenUrl,nil)
	if err != nil{
		return wr,err
	}
	theBody,err := ioutil.ReadAll(theio)
	if err != nil{
		return wr,err
	}
	json.Unmarshal(theBody,&wr)
	return wr,nil
}









/******************************************
*              生成二维码相关             *
*                                         *
*******************************************/


func (c *Wechat) WritePng(filename string, img image.Image) (error){
	file,err := os.Create(filename)
	if err != nil{
		logs.Info("生成写入图片错误")
		logs.Error(err)
		return err
	}
	err2 := png.Encode(file,img)
	if err2 != nil{
		logs.Info("转换为png失败")
		logs.Error(err2)
		return err
	}
	file.Close()
	return nil
}

func (c *Wechat)MakeQRCode(message string,fileName string) (error) {
	code,err := qr.Encode(message,qr.L,qr.Auto)
	if err != nil{
		return err
	}
	code,err = barcode.Scale(code,300,300)
	err = c.WritePng(fileName,code)
	if err != nil{
		return err
	}
	return nil
}
