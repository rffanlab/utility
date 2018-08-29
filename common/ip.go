package common

import (
	"net/http"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/json"
	"github.com/pkg/errors"
)

const IPAPIURL  = "http://ip-api.com/json"

type IPInfo struct {
	As string `json:"as"`
	City string `json:"city"`
	Country string `json:"country"`
	CountryCode string `json:"countryCode"`
	Isp string `json:"isp"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Org string `json:"org"`
	Query string `json:"query"`
	Region string `json:"region"`
	RegionName string `json:"regionName"`
	Status string `json:"status"`
	Timezone string `json:"timezone"`
	Zip string `json:"zip"`
}

type Failed struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Query string `json:"query"`
}


func GetIPinfo(ipaddr string) (IPInfo,error) {
	var failure Failed
	var success IPInfo
	tr := &http.Transport{
		TLSClientConfig:&tls.Config{
			InsecureSkipVerify:true,
		},
	}
	client := &http.Client{Transport:tr}
	realUrl := IPAPIURL + "/"+ipaddr
	req,err := http.NewRequest("GET",realUrl,nil)
	if err != nil{
		fmt.Println(err)
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")
	resp,err := client.Do(req)
	if err != nil{
		fmt.Println(err)
	}
	var respStr string
	respByte,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
	}
	respStr = string(respByte)
	fmt.Println(respStr)
	if strings.Contains(respStr,"fail") {
		json.Unmarshal(respByte,&failure)
		return success,errors.New(failure.Message)
	}else {
		json.Unmarshal(respByte,&success)
		return success,nil
	}
}