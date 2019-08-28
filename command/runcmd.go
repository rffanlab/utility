package command

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os/exec"
)

func RunCmd(theCmd string, args ...string) (out,errout string,err error) {
	cmd := exec.Command(theCmd, args...)
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	stderr,err := cmd.StderrPipe()
	if err != nil {
		logs.Error(err)
		return
	}
	// 保证关闭输出流
	defer stdout.Close()
	defer stderr.Close()
	// 运行命令
	if err = cmd.Start(); err != nil {
		logs.Error(err)
		return
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		logs.Error(err)
		return
	}
	out = string(opBytes)
	opBytes,err = ioutil.ReadAll(stderr)
	if err != nil {
		logs.Error(err)
		return
	}
	errout = string(opBytes)
	return
}