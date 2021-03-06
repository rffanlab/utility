package python

import (
	"utility/request"
	"github.com/PuerkitoBio/goquery"
		"strings"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type Python struct {
	InstallPath string
	Version string
}

const PYTHON_DOWNLOAD_LINKE  = "https://www.python.org/downloads/"

//https://www.python.org/ftp/python/3.7.0/Python-3.7.0.tgz

func (c *Python)GetVersionList() (versions []string,err error) {
	r := request.Requests{}
	respobody,err := r.Get(PYTHON_DOWNLOAD_LINKE,nil)
	if err != nil{
		return nil,err
	}
	document,err := goquery.NewDocumentFromReader(respobody)
	if err!=nil {
		return nil,err
	}
	document.Find("ol.list-row-container.menu").Each(func(i int, selection *goquery.Selection) {
		selection.Find("li").Each(func(i int, selection *goquery.Selection) {
			selection.Find("span.release-number").Each(func(i int, selection *goquery.Selection) {
				ver := selection.Find("a").Eq(0).Text()
				strs := strings.Split(ver," ")
				if len(strs)>0 {
					str := strings.TrimSpace(strs[1])
					if strings.HasSuffix(str,".0") {
						newStr := ""
						theStrs := strings.Split(str,".")
						// 开始判断版本
						bigVersion,err := strconv.Atoi(theStrs[0])
						if err != nil{
							return
						}
						secondVersion,err := strconv.Atoi(theStrs[1])
						if err != nil {
							return
						}
						if bigVersion<3  {
							for i:=0;i<(len(theStrs)-1) ;i++  {
								if i == 0 {
									newStr = theStrs[i]
								}else  {
									newStr = newStr+"."+theStrs[i]
								}
							}
							str = newStr
						}
						if  bigVersion>=3 && secondVersion<=2{
							for i:=0;i<(len(theStrs)-1) ;i++  {
								if i == 0 {
									newStr = theStrs[i]
								}else  {
									newStr = newStr+"."+theStrs[i]
								}
							}
							str = newStr
						}
					}
					versions = append(versions, str)
				}
			})
		})
		selection.Find("")
	})
	for _,value := range versions{
		fmt.Println(value)
	}
	return
}

func (c *Python)MakeDownloadLinkByVer(ver string) (downloadLink string,err error) {
	fmt.Println(ver)
	if ver == "" {
		return "",errors.New("请传入一个正确版本");
	}
	return "https://www.python.org/ftp/python/"+ver+"/Python-"+ver+".tgz",nil
}

func (c *Python) GetLatestVersionOfPython3()  {

}
