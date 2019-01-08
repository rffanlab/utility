package message

type BaseResp struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func Success() (baseResp BaseResp) {
	baseResp.Code = 200
	baseResp.Status = true
	baseResp.Msg = OPERATION_SUCCESS
	return
}

func Fail() (baseResp BaseResp) {
	baseResp.Code = -1
	baseResp.Status = false
	baseResp.Msg = OPERATION_FAILED
	return baseResp
}
