package utility

import (
	"time"
	"fmt"
	"strings"
	"strconv"
)

// 防止忘记
// 时间包相关 内容http://blog.csdn.net/wangshubo1989/article/details/73543377
//
//月份 1,01,Jan,January
//
//日　 2,02,_2
//
//时　 3,03,15,PM,pm,AM,am
//
//分　 4,04
//
//秒　 5,05
//
//年　 06,2006
//
//周几 Mon,Monday
//
//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
//
//时区字母缩写 MST


// 方法：获取当前时间的时间戳
/*
*  传入参数：
*  返回参数：
*  @Param: Type:int64 Comment:时间戳
*  @Param: Type: Comment:
*/
func GetTimestamp() int64 {
	t := time.Now()
	return t.Unix()
}

// 方法：比较现在，和时间戳的差值
/*
*  传入参数：
*  @Param:timestamp Type:int64 Comment:时间戳
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type:bool    Comment: 状态表示是否在当前时间之前，true表示在当前时间之后,false 表示在当前时间之前
*  @Param: Type:float64 Comment:以秒为单位，时间戳，无符号
*  @Param: Type:error Comment:错误
*/
func CompareTimestampNow(timestamp int64) (stat bool,second float64,err error) {
	n := time.Now()
	tm := time.Unix(timestamp,0)
	seconds := n.Sub(tm)
	second = seconds.Seconds()
	stat = n.Before(tm)
	fmt.Println(n)
	fmt.Println(tm)
	return
}


//将特定日期转换为时间格式  2017-10-13 13:38
func TransferDateFromStringToTime(date string) (time.Time,error) {
	t,err := time.Parse("2006-01-02 15:04",date)
	if err != nil{
		return t,err
	}
	return t,nil
}

func TransferTimeToNormalFormat(date string) string {
	splitedStr := strings.Split(date," ")
	hour := splitedStr[1]
	fmt.Println(hour)




	return splitedStr[0]+" "+hour



}

// 当前时间 格式为yyyy-MM-dd HH:mm:ss
func TimeNowForSecond() string {
	return time.Now().Format("2006-01-02 15:04:05")
}




func TimeToTimestamp(date time.Time) int64 {
	return date.Unix()
}


func Year() string {
	return strconv.Itoa(time.Now().Year())
}

func Month() string {
	theMonth := int(time.Now().Month())
	var month string
	if theMonth <10 {
		month = fmt.Sprintf("0%d",theMonth)
	} else {
		month = fmt.Sprintf("%d",theMonth)
	}
	return month
}

func Day() string {
	theDay := time.Now().Day()
	var returnDay string
	if theDay<10 {
		returnDay = fmt.Sprintf("0%d",theDay)
	}else {
		returnDay = fmt.Sprintf("%d",theDay)
	}
	return returnDay
}

func Today() string {
	return fmt.Sprintf("%s_%s_%s",Year(),Month(),Day())
}


