package modules

// User 用户表
type User struct {
	Id            string  `json:"id" gorm:"primaryKey"`
	NickName      string  `json:"nick_name"`                 // 用户昵称
	Username      string  `json:"username" gorm:"unique"`    // 登陆系统用户账号 不允许重复
	Password      string  `json:"password"`                  // 用户密码
	Phone         string  `json:"phone" gorm:"unique"`       // 用户手机号
	Role          int     `json:"role"`                      // 用户角色	1-管理员 2-普通浏览用户 3-房东 4-房东认证后的用户
	Age           int     `json:"age"`                       // 年龄
	Job           string  `json:"job"`                       // 职业/行业
	Birthday      string  `json:"birthday"`                  // 生日
	SelfIntroduce string  `json:"self_introduce"`            // 自我介绍
	Score         float32 `json:"score" gorm:"numeric(3,1)"` // 评价得分
	Picture       string  `json:"picture"`                   // 用户头像

	Name       string `json:"-"`           // 真实姓名
	IdCardNo   string `json:"-"`           // 身份证号
	Status     int    `json:"-"`           // 账号状态 1：正常, 2：注销, 3：封号
	ForbidTime string `json:"forbid_time"` // 账号解封时间
}

func (*User) TableName() string {
	return "user"
}

// 用户注册需要得信息
type RegistInfo struct {
	NickName string `json:"nick_name"`           // 用户昵称
	Picture  string `json:"picture"`             // 用户头像
	Phone    string `json:"phone" gorm:"unique"` // 用户手机号
	Role     int    `json:"role"`                // 用户角色	1-管理员 2-普通浏览用户 3-房东 4-房东认证后的用户
	Age      int    `json:"age"`                 // 年龄
	Birthday string `json:"birthday"`            // 生日
	AuthCode string `json:"auth_code"`           // 短信验证码
}
