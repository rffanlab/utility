package common

import (
	"github.com/pkg/errors"
	"net/http"
	"io"
	"crypto/tls"
	"net/url"
	"fmt"
	"encoding/json"
	"io/ioutil"

	"encoding/xml"
)

type Solusvm struct {
	UserAgent string
	URL string
	Appkey string
	AppHash string
}

type ResponseAction struct {
	Solusvm xml.Name `xml:"solusvm"`
	Status string `xml:"status"`
	Statusmsg string `xml:"statusmsg"`
}

type ResponseStatus struct {
	XMLName xml.Name `xml:"solusvm"`
	Status string `xml:"status"`
	Statusmsg string `xml:"statusmsg"`
	Vmstat string `xml:"vmstat"`
	Hostname string `xml:"hostname"`
	Ipaddress string `xml:"ipaddress"`
}

type ResponseBoot struct {
	Solusvm xml.Name `xml:"solusvm"`
	Status string `xml:"status"`
	Statusmsg string `xml:"statusmsg"`
}

type ResponseShutdown struct {
	Solusvm xml.Name `xml:"solusvm"`
	Status string `xml:"status"`
	Statusmsg string `xml:"statusmsg"`
}


// 检查是否更新了Solusvm的配置
func (c *Solusvm) CheckParams() (error){
	if c.URL == "" {
		return errors.New("请设置URL")
	}
	if c.Appkey == "" {
		return errors.New("请设置appkey")
	}
	if c.AppHash == "" {
		return errors.New("请设置AppHash")
	}
	return  nil
}

func (c *Solusvm)SetCPUrl()  {
	c.URL = c.URL+"/api/client/command.php"
}

func (c *Solusvm) Reboot() ResponseAction {
	var response ResponseAction
	params := make(map[string]string)
	params["key"] = c.Appkey
	params["hash"] = c.AppHash
	params["action"] = "reboot"
	body,err := c.Get(c.URL,params)
	if err != nil{
		fmt.Println(err)
	}
	iobody,err := ioutil.ReadAll(body)
	xmlstr := "<solusvm>"+string(iobody)+"</solusvm>"
	xml.Unmarshal([]byte(xmlstr),&response)
	fmt.Println(response)
	return response
}

func (c *Solusvm) Status() ResponseStatus {
	var response ResponseStatus
	params := make(map[string]string)
	params["key"] = c.Appkey
	params["hash"] = c.AppHash
	params["action"] = "status"
	body,err := c.Get(c.URL,params)
	if err != nil{
		fmt.Println(err)
	}
	iobody,err := ioutil.ReadAll(body)
	fmt.Println(string(iobody))
	xmlstr := "<solusvm>"+string(iobody)+"</solusvm>"
	xml.Unmarshal([]byte(xmlstr),&response)
	fmt.Println(response)
	return response
}

func (c *Solusvm) Boot() ResponseAction {
	var response ResponseAction
	params := make(map[string]string)
	params["key"] = c.Appkey
	params["hash"] = c.AppHash
	params["action"] = "boot"
	body,err := c.Get(c.URL,params)
	if err != nil{
		fmt.Println(err)
	}
	iobody,err := ioutil.ReadAll(body)
	xmlstr := "<solusvm>"+string(iobody)+"</solusvm>"
	xml.Unmarshal([]byte(xmlstr),&response)
	fmt.Println(response)
	return response
}

func (c *Solusvm) Shutdown() ResponseAction {
	var response ResponseAction
	params := make(map[string]string)
	params["key"] = c.Appkey
	params["hash"] = c.AppHash
	params["action"] = "shutdown"
	body,err := c.Get(c.URL,params)
	if err != nil{
		fmt.Println(err)
	}
	iobody,err := ioutil.ReadAll(body)
	xmlstr := "<solusvm>"+string(iobody)+"</solusvm>"
	xml.Unmarshal([]byte(xmlstr),&response)
	fmt.Println(response)
	return response
}

func (c *Solusvm) Info() {
	var response map[string]string
	params := make(map[string]string)
	params["key"] = c.Appkey
	params["hash"] = c.AppHash
	params["action"] = "info"
	body,err := c.Get(c.URL,params)
	if err != nil{
		fmt.Println(err)
	}
	iobody,err := ioutil.ReadAll(body)
	fmt.Println(string(iobody))
	json.Unmarshal(iobody,response)
	fmt.Println(response)
}


func (c *Solusvm)Get(theUrl string,params map[string]string) (io.Reader,error) {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
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
		req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")
	}
	resp,err := client.Do(req)
	if err != nil{
		return nil,err
	}
	return resp.Body,nil
}
