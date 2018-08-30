package main

import (
	"utility/downloader"
	"utility/system"
)

func main() {
	d := downloader.Downloader{}
	d.SavePath = "D:/"
	d.TargetUrl = "http://download.oracle.com/otn-pub/java/jdk/10.0.2+13/19aef61b38124481863b1413dce1855f/jdk-10.0.2_linux-x64_bin.tar.gz"
	d.Cookie = " oraclelicense=accept-securebackup-cookie"
	d.FullDownlod()
	system.DetectOSType()
}
