package nginx

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"utility/request"
)

func (c *Nginx) GetLatestVersion() string {
	nginxVersion := "nginx-1.14.0"
	r := request.Requests{}
	body, err := r.Get(NGINX_OFFICIAL_SITE, nil)
	if err != nil {
		fmt.Println(err)
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		fmt.Println(err)
	}
	content := doc.Find("#content")
	if content != nil {
		stable := content.Find("table").Eq(1)
		if stable != nil {
			nginxVersion = stable.Find("td").Eq(1).Find("a").Eq(0).Text()
		}
	}
	fmt.Println(nginxVersion)
	return nginxVersion
}

func (c *Nginx) GetDownloadLink() string {
	return "https://nginx.org/download/" + c.GetLatestVersion() + ".tar.gz"
}
