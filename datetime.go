package utility

import (
	"strconv"
	"time"
	"fmt"
)

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




