package downloader

import (
	"os"
	"io"
	"net/http"
	"path"
)

type Downloader struct {
	TargetUrl string
	SavePath string
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
	f, err := os.Create(save)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}





