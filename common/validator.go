package common

import (
	"reflect"
)


// 结构体：
/*
*  传入参数：
*  @Param:Status Type:string Comment:
*  @Param:ErrMsg Type:string Comment:
*  @Param: Type:
*  返回参数：
*  @Param: Type:
*  @Param: Type:
*/

type Validator struct {
	Status bool
	ErrMsg string
}

// 方法：正整数
/*
*  传入参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func (c *Validator) IDMustBePositiveInteger(id interface{})  {
	if reflect.TypeOf(id).Name() == "int" && id.(int) >0 {
		c.Status = true
	}else {
		c.Status = false
		c.ErrMsg = "ID必须是正整数"
	}
}

func (c *Validator) CheckNumOnlyHasTwoDigit(num float64) {

}


