package common

import "fmt"

/******************************************
*           整体的杂项方法包              *
*                                         *
*******************************************/
func Gavatar(email string) string {
	enEmail := Md5Encryption(email)
	return fmt.Sprintf("//cn.gravatar.com/avatar/%s?s=44&r=g",enEmail)
}


func MakeUserkey() string {
	return RandStr(4)+"-"+RandStr(5)+"-"+RandStr(5)+"-"+RandStr(4)
}