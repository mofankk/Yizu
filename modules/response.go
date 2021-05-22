package modules

import "encoding/json"

type ResInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Success() ResInfo {
	return ResInfo{
		Code: 0,
		Msg:  "操作成功",
	}
}

func ArgErr() ResInfo {
	return ResInfo{
		Code: 1,
		Msg:  "参数错误",
	}
}

func SysErr() ResInfo {
	return ResInfo{
		Code: 1,
		Msg:  "系统出错，请稍后再试",
	}
}

func InsertErr() []byte {
	x, _ := json.Marshal(ResInfo{
		Code: 1,
		Msg:  "创建失败",
	})
	return x
}

func UpdateErr() []byte {
	x, _ := json.Marshal(ResInfo{
		Code: 1,
		Msg:  "修改失败",
	})
	return x
}

type ResultInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Result interface{} `json:"result"`
}

func QuerySuccess() ResultInfo {
	x := ResultInfo{
		Code: 1,
		Msg:  "修改失败",
		Result: nil,
	}
	return x
}

func LoginSuccess() ResInfo {
	return ResInfo{
		0,
		"登陆成功",
	}
}

func LoginFail() ResInfo {
	return ResInfo{
		1,
		"登陆失败",
	}
}
