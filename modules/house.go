package modules

// House 房子表
type House struct {
	Id         string `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"` // 房屋名称
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	DeleteTime string `json:"delete_time"`
	CreateUser string `json:"create_user"` // 创建人,可以理解为房子目前的房东

	Province string `json:"province"` // 省
	City     string `json:"city"`     // 市
	District string `json:"district"` // 区
	Street   string `json:"street"`   // 街道

	Rooms int `json:"rooms"` // 卧室数量

	Rent     float32 `json:"rent" gorm:"type:numeric(5,2)"`    // 月租金
	PayCycle string  `json:"pay_cycle"`                        // 支付周期(月付、季付、年付)
	Deposit  float32 `json:"deposit" gorm:"type:numeric(5,2)"` // 押金

	Address string `json:"address"` // 房屋详细地址
	ImgUrl  string `json:"img_url"` // 房屋列表图片

	Score       float32 `json:"score" gorm:"numeric(3,1)"` // 评价得分
	Description string  `json:"description"`               // 房屋简介
}

// HouseDetail 房子详情
type HouseDetail struct {
	Id       string `json:"id" gorm:"primaryKey"`
	HouseId  string `json:"house_id"`
	Name     string `json:"name"`     // 区域名称
	Describe string `json:"describe"` // 描述
	ImgPath  string `json:"img_path"` // 图片路径, 因为可能有多个,这里只存储图片的目录路径
}

// HouseHistory 房子浏览历史
type HouseHistory struct {
	Id          int     `json:"-" gorm:"primaryKey;AUTO_INCREMENT"`
	UserId      string  `json:"-"`
	HouseId     string  `json:"house_id"`
	ImgUrl      string  `json:"img_url"`
	Name        string  `json:"name"`
	Rent        float32 `json:"rent" gorm:"type:numeric(5,2)"` // 每月租金
	PayCycle    string  `json:"pay_cycle"`                     // 支付周期(月付、季付、年付)
	Description string  `json:"description"`                   // 房屋简介
}

func (*House) TableName() string {
	return "house"
}
func (*HouseDetail) TableName() string {
	return "house_detail"
}
func (*HouseHistory) TableName() string {
	return "house_history"
}

// HouseQueryArg 房源列表查询条件
// 后期可以设置多个街道或者标志性建筑这种查询条件
type HouseQueryArg struct {
	Province string `json:"province"` // 省
	City     string `json:"city"`     // 市
	District string `json:"district"` // 区
	Street   string `json:"street"`   // 街道

	StartTime string `json:"start_time"` // 发布时间-开始
	EndTime   string `json:"end_time"`   // 发布时间-截止

	RentUp  float32 `json:"rent_up"`  // 月租金-最大值
	RentLow float32 `json:"rent_low"` // 月租金-最小值

	UpdateTime string `json:"update_time"` // 更新时间 用于对显示结果排序

	Page     int `json:"page"`      // 当前页码
	PageSize int `json:"page_size"` // 每页数量
}
