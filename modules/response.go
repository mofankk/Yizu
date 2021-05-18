package modules

import "encoding/json"

type resInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Success() []byte {
	x, _ := json.Marshal(resInfo{
		Code: 0,
		Msg:  "操作成功",
	})
	return x
}

func ArgErr() []byte {
	x, _ := json.Marshal(resInfo{
		Code: 1,
		Msg:  "参数错误",
	})
	return x
}

func SysErr() []byte {
	x, _ := json.Marshal(resInfo{
		Code: 1,
		Msg:  "系统出错，请稍后再试",
	})
	return x
}

func InsertErr() []byte {
	x, _ := json.Marshal(resInfo{
		Code: 1,
		Msg:  "创建失败",
	})
	return x
}

func UpdateErr() []byte {
	x, _ := json.Marshal(resInfo{
		Code: 1,
		Msg:  "修改失败",
	})
	return x
}


