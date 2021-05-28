package modules

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

func Failure() ResInfo {
	return ResInfo{
		Code: 1,
		Msg:  "操作失败",
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

func InsertErr() ResInfo {
	return ResInfo{
		Code: 1,
		Msg:  "创建失败",
	}
}

func UpdateErr() ResInfo {
	return ResInfo{
		Code: 1,
		Msg:  "修改失败",
	}
}

type ResultInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

func QuerySuccess() ResultInfo {
	x := ResultInfo{
		Code: 0,
		Msg:  "查询成功",
		Data: nil,
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

func NoRecord() ResInfo {
	return ResInfo{
		1,
		"没有找到该记录",
	}
}

func CookieErr() ResInfo {
	return ResInfo{
		1,
		"请携带Cookie请求",
	}
}

func SessionErr() ResInfo {
	return ResInfo{
		1,
		"没有Session, 请重新登陆",
	}
}
