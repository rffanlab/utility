package command

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func RunCmd(theCmd string, args ...string) (out,errout string,err error) {
	cmd := exec.Command(theCmd, args...)
	// 获取输出对象，可以从该对象中读取输出结果
	workingDir,err := GetWorkingDir(theCmd)
	if err != nil {
		logs.Error(err)
		return
	}
	cmd.Dir = workingDir
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

func GetWorkingDir(cmdPath string) (working string ,err error) {
	fi,err := os.Stat(cmdPath)
	if err != nil{
		logs.Error(err)
		return
	}
	if fi.IsDir() {
		working = cmdPath
	}else {
		cmdPath = strings.Replace(cmdPath,"\\","/",-1)
		pos := strings.LastIndex(cmdPath,"/")
		if pos>0 {
			working = cmdPath[0:pos]
		}else {
			working = cmdPath
		}

	}


	return
}