package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

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
func Md5(theStr string) string {
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
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano()+ int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}