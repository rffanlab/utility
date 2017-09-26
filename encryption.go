package utility

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"github.com/pkg/errors"
	"time"
)
/******************************************
*             加密方法包                  *
*                                         *
*******************************************/

// 方法：MD5加密（返回32位的md5加密字符串）
/*
*  传入参数：
*  @Param:theStr Type:string
*  @Param: Type:
*  @Param: Type:
*  返回参数：
*  @Param:string Type:string
*  @Param: Type:
*/
func Md5Encryption(theStr string) string {
	h := md5.New()
	h.Write([]byte(theStr))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 方法：生成随机字符串
/*
*  传入参数：
*  @Param:length Type:int
*  @Param: Type:
*  @Param: Type:
*  返回参数：
*  @Param: Type:string
*  @Param: Type:
*/
func RandStr(length int) string {
	// 已经被注释的随机字符串方法是伪随机字符串方法
	//var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	//b := make([]rune,length)
	//for i := range b {
	//	b[i] = letters[rand.Intn(len(letters))]
	//}
	//return string(b)
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano()+ int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// 方法：随机验证码
/*
*  传入参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func RandVerifyCode(length int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()+ int64(rand.Intn(100))))
	for i := 0 ;i<length;i++{
		result = append(result,bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 方法：盐值加密(盐值长度为固定4个字符串)
/*
*  传入参数：
*  @Param:strToEncrypt Type:string
*  @Param: Type:
*  @Param: Type:
*  返回参数：
*  @Param: Type:string
*  @Param: Type:
*/
func EncryptStrWithSalt(strToEncrypt string) string {
	salt := RandStr(4)
	logStr := strToEncrypt + salt
	encryptedStr := Md5Encryption(logStr)
	return encryptedStr+salt
}

// 方法：传入的字符串与盐值加密后的字符串的对比
/*
*  传入参数：
*  @Param:strToCompare Type:string
*  @Param:encryptStr Type:string
*  @Param: Type:
*  返回参数：
*  @Param:bool Type:bool
*  @Param:error Type:error
*/
func CompareStrToSaltEncryptedStr(strToCompare, encryptedStr string) (bool,error) {
	if len(encryptedStr) != 36 {
		return false,errors.New("Not A Vaild String To Compare")
	}
	salt := string([]rune(encryptedStr)[32:36])
	newEncrypt := Md5Encryption(strToCompare+salt)
	if newEncrypt+salt == encryptedStr{
		return true,nil
	}else {
		return false,nil
	}
}

