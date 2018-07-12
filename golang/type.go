package golang

import (
	"utility/request"
	)

type Golang struct {
	InstallPath string
	Version string



}

const GOLANG_OFFICIAL_SITE  = "https://golang.org"

func (c *Golang) GetLatestVersion() {
	downloaderUrl := GOLANG_OFFICIAL_SITE+"/dl/"
	r:= request.Requests{
	}
	r.Get(downloaderUrl,nil)

	}



