package common

const (
	SHUIJIN_API string = "1-api-red.xunlei.com"
)

type Xunlei struct {
	Username string
	Password string
	GiftOpen int    // 0 关闭，1开启，默认为0
	Cnum int        //  开启宝箱可以小号的原石，默认为0 可以设置为500 即小于500的消耗都会被开启
	TurnTable int   // 轮盘游戏的开启，默认为0 关闭状态，1为开启
}

// 方法：设置是否开启宝箱开关
/*
*  传入参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*  返回参数：
*  @Param: Type: Comment:
*  @Param: Type: Comment:
*/
func (c *Xunlei) TurnOnGiftOpen() (*Xunlei) {
	c.GiftOpen = 1
	return c
}

func (c *Xunlei) SetLimitForOpenGift(num int) (*Xunlei) {
	c.Cnum = num
	return c
}

func (c *Xunlei)TurnOnTrunTable() (*Xunlei) {
	c.TurnTable = 1
	return c
}

func (c *Xunlei) Login() {
	// TODO 迅雷的登陆方法
}

func (c *Xunlei) GiftBox() {
	// TODO 迅雷宝箱方法
}

func (c *Xunlei)FetchCrystal()  {
	// TODO 获取迅雷水晶的方法
}

func (c *Xunlei)TrunTable()  {
	// TODO 迅雷轮盘的使用
}