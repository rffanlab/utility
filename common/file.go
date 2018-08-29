package common

import (
	"os"
	"io"
	"io/ioutil"
	"github.com/astaxie/beego/logs"
	"bufio"
	"bytes"
	"math"
	"crypto/md5"
	"encoding/hex"
	"strings"
	"fmt"
	"path"
)

const FILECHUNK  = 8192
// 方法：复制文件
/*
*  传入参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func CopyFile(srcPath, dstPath string) (written int64, err error) {
	src,err := os.Open(srcPath)
	if err != nil{
		return
	}
	defer src.Close()
	dst,err := os.OpenFile(dstPath,os.O_WRONLY|os.O_CREATE,0644)
	if err != nil{
		return
	}
	defer dst.Close()
	return io.Copy(dst,src)
}


// 方法：移动文件
/*
*  传入参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func MoveFile(srcPath, dstPath string) (bool, error) {
	_,err := CopyFile(srcPath,dstPath)
	if err != nil{
		return false,err
	}
	os.Remove(srcPath)
	return true,nil
}

// 方法：读取文本文档
/*
*  传入参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func ReadTextFile(srcPath string) string {
	buf,err := ioutil.ReadFile(srcPath)
	if err != nil{
		logs.Error(err)
	}
	return string(buf)
}

// 方法：逐行读取文档
/*
*  传入参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func ReadLines(path string) (lines []string, err error) {
	var (
		file *os.File
		part [] byte
		prefix bool
	)

	if file, err = os.Open(path); err != nil {
		return
	}

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte,1024))

	for {
		if part, prefix, err = reader.ReadLine();err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines,buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func AppendLine(line string,filePath string) error  {
	if strings.HasPrefix(filePath,"./") || strings.HasPrefix(filePath,".\\") {
		dir := GetCurrentDirectory()
		filePath = path.Join(dir,filePath)
	}
	if !Exist(filePath) {
		fmt.Errorf("文件不存在，创建文件中")
		Touch(filePath)
	}
	fp,err := os.OpenFile(filePath,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		return err
	}
	buf := []byte(line)
	fp.Write(buf)
	fp.Close()
	return nil
}



// 方法：大文档md5
/*
*  传入参数：
*  @Param:theFilePath Type:string Comment:传入文档路径
*  返回参数：
*  @Param:themd5 Type:string Comment:md5返回值
*  @Param:err Type:error Comment:错误
*/
func LargeFilemd5(theFilepath string) (themd5 string,err error) {
	file,err := os.Open(theFilepath)
	if err != nil{
		return
	}
	defer file.Close()
	info,err := file.Stat()
	if err != nil{
		return
	}
	filesize := info.Size()
	blocks := uint64(math.Ceil(float64(filesize)/float64(FILECHUNK)))
	hash := md5.New()
	for i := uint64(0);i <blocks;i++{
		blocksize := int(math.Min(FILECHUNK,float64(filesize-int64(i*FILECHUNK))))
		buf := make([]byte,blocksize)
		file.Read(buf)
		io.WriteString(hash,string(buf))
	}
	themd5 = hex.EncodeToString(hash.Sum(nil))
	return
}

// 方法：小文档md5
/*
*  传入参数：
*  @Param:theFilePath Type:string Comment:传入文档路径
*  返回参数：
*  @Param:themd5 Type:string Comment:md5返回值
*  @Param:err Type:error Comment:错误
*/
func Filemd5(theFilePath string) (themd5 string, err error) {
	file,err := os.Open(theFilePath)
	if err != nil{
		return
	}
	hash := md5.New()
	io.Copy(hash,file)
	themd5 = hex.EncodeToString(hash.Sum(nil))
	return
}

// 方法：小文档sha1
/*
*  传入参数：
*  @Param:theFilePath Type:string Comment:传入文档路径
*  返回参数：
*  @Param:themd5 Type:string Comment:md5返回值
*  @Param:err Type:error Comment:错误
*/
func FileSha1(theFilePath string) (thesha1 string, err error) {

	return
}

// 方法：大文档md5
/*
*  传入参数：
*  @Param:theFilePath Type:string Comment:传入文档路径
*  返回参数：
*  @Param:themd5 Type:string Comment:md5返回值
*  @Param:err Type:error Comment:错误
*/
func LargeFileSha1(theFilePath string) (thesha1 string, err error) {

	return
}


// 方法：创建空文档
/*
*  传入参数：
*  @Param:theFilePath Type:string Comment:要新建的文档的路径
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func Touch(theFilePath string) (stat bool,err error) {
	f ,  err:= os.Create(theFilePath)
	defer f.Close()
	if err != nil{
		return
	}
	stat = true
	return
}

func GetFileSuffix(filepath string) (suffix,body string) {
	// 由于split出来的不可能小于1 所以就不判断小1的情况了。
	strs := strings.Split(filepath,".")
	if len(strs) == 1 {
		return "",filepath
	}else {
		tmpstr := ""
		for i := 0;i <len(strs) -1 ;i ++ {
			tmpstr = tmpstr + strs[i]
		}
		return strs[1],tmpstr
	}
}

// 文件是否存在
func Exist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fmt.Println(err)
		return false;
	}
	return true
	
}




