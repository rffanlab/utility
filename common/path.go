package common

import (
	"path/filepath"
	"os"
	"strings"
	"github.com/astaxie/beego/logs"
)


// 方法：获取当前程序运行的路径(例如：/home/gopath/src/ClientManagement/ClientManagement)
/*
*  传入参数：
*  @Param: Type:
*  @Param: Type:
*  @Param: Type:
*  返回参数：
*  @Param: Type:string
*  @Param: Type:
*/
func GetCurrentDirectory() string {
	dir,err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil{
		logs.Error(err)
	}
	return strings.Replace(dir,"\\","/",-1)
}
// 方法：检查路径存在
/*
*  传入参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func PathExistance(path string) (bool,bool, error) {
	stat,err := os.Stat(path)
	if err == nil{
		return true,stat.IsDir(),nil
	}else {
		if os.IsNotExist(err){
			return false,false,nil
		}
		return false,stat.IsDir(),err
	}
}

// 方法：自动创建目录
/*
*  传入参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func AutoCreateFolder(path string) (bool) {
	stat,_,err := PathExistance(path)
	if err != nil{
		return false
	}
	if stat {
		return true
	}
	cErr := os.MkdirAll(path,655)
	if cErr != nil{
		return false
	}else {
		return true
	}
}

// 方法：判断是否是相对路径
/*
*  传入参数：
*  @Param:thePathToJudge Type:string Comment:需要判断的路径
*  返回参数：
*  @Param:is Type:bool Comment:判断是否为相对路径
*  @Param:err Type:error Comment:错误返回
*/
func IsRelativePath(thePathToJudge string) (is bool) {
	if strings.HasPrefix(thePathToJudge,"./") {
		is = true
	}else if strings.HasPrefix(thePathToJudge,"/") {
		is = false
	}else {
		if string(thePathToJudge[1]) == ":" {
			is = false
		}else {
			is = true
		}
	}
	return
}


