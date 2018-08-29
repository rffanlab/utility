package request

import (
	"net/http"
	"crypto/tls"
	"fmt"
	"net/url"
	"io"
	"strings"
	"golang.org/x/net/proxy"
)

type Requests struct {
	Url string
	UserAgent string 
	StatusCode int
	Proxy string
	Headers map[string]string
}


//
func (c *Requests)SetUserAgent(useragent string) {
	c.UserAgent = useragent
}

func (c *Requests)setHeaders(headerParams map[string]string )  {
	c.Headers = headerParams
}

// 传入参数：params 必须是string的map

func (c *Requests)Get(theUrl string,params map[string]string) (io.Reader,error) {
	if theUrl == "" {
		theUrl = c.Url
	}
	var tr *http.Transport
	if c.Proxy != "" {
		dialSocksProxy, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
		if err != nil {
			fmt.Println("Error connecting to proxy:", err)
		}
		fmt.Println("已经过了检测了")
		tr = &http.Transport{
			Dial: dialSocksProxy.Dial,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
	}else {
		tr = &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{Transport:tr}
	realUrl := ""
	if params != nil{
		readerString := ""
		for k,v := range params{
			if readerString == ""{
				readerString = fmt.Sprintf("%s=%s",k,url.QueryEscape(v)) //进行URL编码的参数
			}else {
				readerString = fmt.Sprintf("%s&%s=%s",readerString,k,url.QueryEscape(v))
			}
		}
		realUrl = fmt.Sprintf("%s?%s",theUrl,readerString)
	}else {
		realUrl = theUrl
	}
	req,err := http.NewRequest("GET",realUrl,nil)
	if err != nil{
		return nil,err
	}
	if c.UserAgent != ""{
		req.Header.Set("User-Agent",c.UserAgent)
	}else{
		req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	}

	if c.Headers != nil {
		for key,value := range c.Headers {
			req.Header.Set(key,value)
		}
	}



	resp,err := client.Do(req)
	if err != nil{
		return nil,err
	}
	// Update Struct 更新结构体
	c.SetUrl(resp.Request.URL.String())
	c.SetStatusCode(resp.StatusCode)

	return resp.Body,nil
}

// Post 方法
/*
*  传入参数：theUrl 类型string，params 类型map key和value都是string
*  返回参数：io.Reader,错误
*/
func (c *Requests)Post(theUrl string,params map[string]string) (io.Reader,error) {
	var tr *http.Transport
	if c.Proxy != "" {
		tr = &http.Transport{
			Proxy: func(request *http.Request) (*url.URL, error) {
				return url.Parse(c.Proxy)
			},
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
	}else {
		tr = &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		}
	}
	client := &http.Client{Transport:tr}
	var body io.Reader
	// 参数生成
	if params != nil{
		values := url.Values{}
		for k,v := range params{
			values.Set(k,v)
		}
		body = strings.NewReader(values.Encode())
	}else {
		body = nil
	}
	req,err := http.NewRequest("POST",theUrl,body)
	if err != nil{
		return nil,err
	}
	if c.UserAgent != ""{
		req.Header.Set("User-Agent",c.UserAgent)
	}else{
		req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	}
	if c.Headers != nil {
		for key,value := range c.Headers {
			req.Header.Set(key,value)
		}
	}

	resp,err := client.Do(req)
	if err != nil{
		return nil,err
	}

	// Update Struct 更新结构体
	c.SetUrl(resp.Request.URL.String())
	c.SetStatusCode(resp.StatusCode)
	return resp.Body,nil
}


// Setter
/*
* 设置结构体的Url
*
*/
func (c *Requests)SetUrl(theUrl string)  {
	c.Url = theUrl
}
/*  设置结构体的状态码
*   传入参数：状态码，类型int
*   返回参数：无
*/
func (c *Requests)SetStatusCode(statusCode int)  {
	c.StatusCode = statusCode
}

func (c *Requests) Download() {

}
