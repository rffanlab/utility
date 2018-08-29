package common

// 写于 2017-09-26 预计半年内是生效的。

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
	"utility/request"
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
*        微信公众号相关结构体             *
*                                         *
*******************************************/

// 通用错误结构体
type WechatErrMsg struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
}

// 获取access_token的结构体
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
}
// 错误结构体
type AccessTokenErrResponse struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
}

// 获取用户基本信息结构体 相关介绍，请查看URL：https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140839
type WechatUserinfo struct {
	Subscribe int `json:"subscribe"`
	Openid string `json:"openid"`
	Nickname string `json:"nickname"`
	Sex int `json:"sex"`
	Language string `json:"language"`
	City string `json:"city"`
	Province string `json:"province"`
	Country string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	SubscribeTime int `json:"subscribe_time"`
	Unionid string `json:"unionid"`
	Remark string `json:"remark"`
	Groupid int `json:"groupid"`
	TagidList []int `json:"tagid_list"`
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
	encodeMsg := strings.ToLower(url.QueryEscape(msg))
	logs.Info("编码后的url是：",encodedUrl)
	urlmsg := "https://open.weixin.qq.com/connect/oauth2/authorize?appid="+c.Appkey+"&redirect_uri="+encodedUrl+"&response_type=code&scope=snsapi_userinfo&state="+encodeMsg+"#wechat_redirect"
	err := c.MakeQRCode(urlmsg,codePath)
	if err != nil{
		logs.Info("生成二维码时出错，请查看下面的错误内容")
		logs.Error(err)
	}
	return codePath,nil
}

// 方法：获取用户基本信息
/*
*  传入参数：
*  @Param:code Type:string Comment:通过微信登录来获得的code
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func (c *Wechat) GetUserWechatInfo(withCode string) (WechatResponse,error) {
	var wr WechatResponse
	err := c.CheckConfigSet()
	if err != nil{
		return wr,err
	}
	getAccess_tokenUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid="+c.Appkey+"&secret="+c.AppSecret+"&code="+ withCode + "&grant_type=authorization_code"
	r := request.Requests{}
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

// 方法：获取微信公众号AccessToken
/*
*  传入参数：
*  @Param:appkey Type:string Comment:就是腾讯的appid 请在结构体初始化的时候就直接赋值
*  @Param:appsecret Type:string Comment:请在结构体初始化的时候就直接赋值
*  返回参数：
*  @Param: Type:string Comment:AccessToken
*  @Param: Type:string Comment:错误
*/
func (c *Wechat) GetAccessToken() (string,error) {

	requestUrl := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="+c.Appkey+"&secret="+c.AppSecret
	r := request.Requests{}
	theio,err := r.Get(requestUrl,nil)
	if err != nil{
		fmt.Println(err)
	}
	theBody,err := ioutil.ReadAll(theio)
	if err != nil{
		fmt.Println(err)
	}
	theStr := string(theBody)
	if strings.Contains(theStr,"errcode"){
		fmt.Println("错了")
		var acerr AccessTokenErrResponse
		json.Unmarshal([]byte(theStr),&acerr)
		return "",errors.New(acerr.Errmsg)
	}else {
		fmt.Println("没错")
		var ac AccessTokenResponse
		json.Unmarshal([]byte(theStr),&ac)
		return ac.AccessToken,nil
	}
}



// 方法：检查用户是否订阅
/*
*  传入参数：
*  @Param:openid Type:string Comment:
*  @Param:access_token Type:string Comment:注意该access_token 不是网页授权登录的access_token,而是微信公众号接口全局唯一凭证
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func (c *Wechat) IsSubscribe(openid,access_token string) (bool, error) {
	requestUrl := "https://api.weixin.qq.com/cgi-bin/user/info?access_token="+access_token+"&openid="+openid+"&lang=zh_CN "
	r := request.Requests{}
	theio,err := r.Get(requestUrl,nil)
	if err != nil{
		return false,err
	}
	theBody,err := ioutil.ReadAll(theio)
	if err != nil {
		return false,err
	}
	if strings.Contains(string(theBody),"errcode"){
		var err WechatErrMsg
		json.Unmarshal(theBody,&err)
		return false,errors.New(err.Errmsg)
	}else {
		var accessToken WechatUserinfo
		json.Unmarshal(theBody,&accessToken)
		if accessToken.Subscribe == 1 {
			return true,nil
		}else {
			return false,nil
		}
	}

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
