package common

import (
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"time"
	"fmt"
	"sort"
	"crypto/md5"
	"crypto/hmac"
	"io"
	"crypto/tls"
	"net/url"
	"errors"
)

const (
	// httpURL is for HTTP REST API URL.
	httpURL string = "http://gw.api.taobao.com/router/rest"
	// httpsURL is for HTTPS REST API URL.
	httpsURL string = "https://eco.taobao.com/router/rest"
)

type Client struct {
	AppKey string
	AppSecret string
	UseHttps bool
	Method string
	Sign_method string
	Format string
}

type SMSClient struct {
	Client Client
	Extend string
	SmsType string
	SmsFreeSignName string
	SmsParam string
	TemplateCode string
	RecieveMobile string
}
// 设置默认的appkey和secret
func (c *SMSClient)SetDefaultClient(appkey,appsecret string) *SMSClient {
	client := Client{
		AppKey:appkey,
		AppSecret:appsecret,
		UseHttps:true,
		Method:"alibaba.aliqin.fc.sms.num.send",
		Sign_method:"md5",
		Format:"json",
	}
	c.Client = client
	return c
}
// 设置发送类型
func (c *SMSClient) SetSMSType(theType string) (*SMSClient) {
	c.SmsType = theType
	return c
}
//设置签名
func (c *SMSClient) SetSMSFreeSignName(signName string) *SMSClient {
	c.SmsFreeSignName = signName	
	return c
}
// 设置模板编码
func (c *SMSClient) SetSMSTemplateCode(templatecode string) *SMSClient {
	c.TemplateCode = templatecode
	return c
}
//设置发送的手机号码
func (c *SMSClient) SetRecieveMobile(recieveMobile string) *SMSClient {
	c.RecieveMobile = recieveMobile
	return c
}

func (c *SMSClient) SetParams(params string) *SMSClient {
	c.SmsParam = params
	return c
}
// 发送方法
func (c *SMSClient)Send() (map[string]interface{},error) {
	if c.SmsType == "" || c.SmsFreeSignName == "" || c.TemplateCode == "" || c.RecieveMobile==""||c.SmsParam == ""{
		return nil,errors.New("请设置所有的参数")
	}
	params := make(map[string]string)
	params["sms_type"] = c.SmsType
	params["sms_free_sign_name"] = c.SmsFreeSignName
	params["sms_template_code"] = c.TemplateCode
	params["rec_num"] = c.RecieveMobile
	params["sms_param"] = c.SmsParam
	return c.Client.DoRequest(params)
}








// 设置通用的Map
func (c *Client)SetCommonParams() map[string]string {
	params := make(map[string]string)
	t := time.Now()
	params["method"] = c.Method
	params["format"] = c.Format
	params["v"] = "2.0"
	params["sign_method"] = c.Sign_method
	params["app_key"] = c.AppKey
	params["timestamp"] = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return params
}


// 生成全部参数
func (c *Client)SetRequestParams(requestParams,commonParams map[string]string) map[string]string {
	for k := range commonParams{
		requestParams[k] = commonParams[k]
	}
	return requestParams
}

// 给参数排序，并返回数值
func (c *Client)SortParamsToStr(params map[string]string) string {
	var keys []string
	for key := range params{
		keys = append(keys,key)
	}
	sort.Strings(keys)
	str := ""
	for _,k := range keys{
		str += k + params[k]
	}
	return str
}


// 签名算法
/*
* API算法的URL：http://open.taobao.com/docs/doc.htm?spm=a219a.7395905.0.0.hsp22E&articleId=101617&docType=1&treeId=1
*
*/
func (c *Client)SignMD5(params map[string]string) (string) {
	str := fmt.Sprintf("%s%s%s",c.AppSecret,c.SortParamsToStr(params),c.AppSecret)
	return fmt.Sprintf("%X",md5.Sum([]byte(str)))
}
/*
* HMAC 加密算法
*
*/
func (c *Client)SignHMAC(params map[string]string) (string) {
	str := c.SortParamsToStr(params)
	mac := hmac.New(md5.New,[]byte(c.AppSecret))
	mac.Write([]byte(str))
	return fmt.Sprintf("%X",mac.Sum(nil))
}

// 创建请求体

func (c *Client)MakeRequestBody(params map[string]string) (io.Reader,error) {
	values := url.Values{}
	for k,v := range params{
		values.Set(k,v)
	}
	return strings.NewReader(values.Encode()),nil
}
// 开始请求
/*
*  传入参数：一个包含所有参数的map
*  返回参数：返回一个map
*/
func (c *Client)DoRequest(params map[string]string) (map[string]interface{},error) {
	commonParams := c.SetCommonParams()
	requestParams := c.SetRequestParams(params,commonParams)
	if requestParams["sign_method"] == "md5" {
		sign := c.SignMD5(requestParams)
		requestParams["sign"] = sign
	}else if requestParams["sign_method"] == "hmac"{
		sign := c.SignHMAC(requestParams)
		requestParams["sign"] = sign
	}else {
		fmt.Errorf("签名方法配置错误")
	}
	values := url.Values{}
	for k,v := range requestParams{
		values.Set(k,v)
	}
	requestUrl := ""
	if c.UseHttps {
		requestUrl = httpsURL
	}else {
		requestUrl = httpURL
	}
	tr := &http.Transport{
		TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
	}
	requestClient := &http.Client{Transport:tr}
	requestBody,err := c.MakeRequestBody(requestParams)
	if err != nil{
		return nil,err
	}
	req,err := http.NewRequest("POST",requestUrl,requestBody)
	if err != nil{
		return nil,err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp,err := requestClient.Do(req)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil,err
	}
	responseStr := string(body)
	if strings.Contains(responseStr,"error"){

	}
	tbPwd := make(map[string]interface{})
	json.Unmarshal(body,&tbPwd)
	return tbPwd,nil
}

