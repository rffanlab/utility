package common

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const IPAPIURL = "http://ip-api.com/json"

type IPInfo struct {
	As          string  `json:"as"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Isp         string  `json:"isp"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Org         string  `json:"org"`
	Query       string  `json:"query"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	Status      string  `json:"status"`
	Timezone    string  `json:"timezone"`
	Zip         string  `json:"zip"`
}

type Failed struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Query   string `json:"query"`
}

func IsIPv4(ip string) bool {
	stat := false
	patten := "^((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)$"
	reg := regexp.MustCompile(patten)
	match := reg.FindAllString(ip, -1)
	if len(match) > 0 {
		stat = true
	}
	return stat
}

func GetIPinfo(ipaddr string) (IPInfo, error) {
	var failure Failed
	var success IPInfo
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	realUrl := IPAPIURL + "/" + ipaddr
	req, err := http.NewRequest("GET", realUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	var respStr string
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	respStr = string(respByte)
	fmt.Println(respStr)
	if strings.Contains(respStr, "fail") {
		json.Unmarshal(respByte, &failure)
		return success, errors.New(failure.Message)
	} else {
		json.Unmarshal(respByte, &success)
		return success, nil
	}
}
