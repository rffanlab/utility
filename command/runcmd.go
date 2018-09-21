package command

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os/exec"
)

func RunCmd(theCmd string, args ...string) (out string,err error) {
	cmd := exec.Command(theCmd, args...)
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logs.Error(err)
		return
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err = cmd.Start(); err != nil {
		logs.Error(err)
		return
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		logs.Error(err)
	}
	logs.Info(string(opBytes))
	out = string(opBytes)
	return
}