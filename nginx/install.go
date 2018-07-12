package nginx

import "utility/downloader"

func (c *Nginx) InstallNginx() {
	d := downloader.Downloader{
		TargetUrl:c.GetDownloadLink(),
	}
	d.LiteDown()




}
