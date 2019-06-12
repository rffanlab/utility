package downloader

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"utility/system"
)

type Downloader struct {
	TargetUrl string
	SavePath string
	Cookie string
}

func (c *Downloader) Wget() {

}

func (c *Downloader) LiteDown() {
	res, err := http.Get(c.TargetUrl)
	if err != nil {
		panic(err)
	}
	var save string
	if c.SavePath != "" {
		save = path.Join(c.SavePath,path.Base(c.TargetUrl))
	}else {
		save = path.Base(c.TargetUrl)
	}
	c.SavePath = save
	f, err := os.Create(save)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}

//
func (c *Downloader)FullDownlod() error {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	//req,err := http.NewRequest("HEAD",c.TargetUrl,nil)
	//if err != nil {
	//	return err
	//}
	//if c.Cookie != "" {
	//	req.Header.Set("Cookie",c.Cookie)
	//	req.Header.Set("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36")
	//}
	client:=&http.Client{Transport:tr}
	//resp,err := client.Do(req)
	//if err != nil {
	//	return  err
	//}
	//fmt.Println(resp.ContentLength)
	//acceptRange := resp.Header.Get("Accept-Ranges")
	acceptRange :=""
	var save string
	if c.SavePath != "" {
		save = path.Join(c.SavePath,path.Base(c.TargetUrl))
	}else {
		save = path.Base(c.TargetUrl)
	}
	c.SavePath = save
	if acceptRange != "" {
		 fmt.Println("可以开启多线程下载")
		 fmt.Printf("获取系统线程数%d",system.GetThreadNum())
		nreq,err := http.NewRequest("GET",c.TargetUrl,nil)
		if err != nil {
			return err
		}
		if c.Cookie!="" {
			fmt.Println(c.Cookie)
			nreq.Cookies()
			nreq.Header.Set("Cookie",c.Cookie)
			nreq.Header.Set("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36")
		}
		resp,err := client.Do(nreq)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		f,err := os.Create(save)
		if err != nil {
			return err
		}
		io.Copy(f,resp.Body)
	}else {
		fmt.Println("无法开启多线程下载，正在单线程下载中")
		nreq,err := http.NewRequest("GET",c.TargetUrl,nil)
		if err != nil {
			return err
		}
		if c.Cookie!="" {
			fmt.Println(c.Cookie)
			nreq.Header.Set("Cookie",c.Cookie)
		}
		resp,err := client.Do(nreq)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		f,err := os.Create(save)
		if err != nil {
			return err
		}
		io.Copy(f,resp.Body)
	}
	return nil
}




