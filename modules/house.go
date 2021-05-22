package modules

// House 房子表
type House struct {
	Id         string `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"` // 房屋名称
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	DeleteTime string `json:"delete_time"`
	CreateUser string `json:"create_user"` // 创建人

	Province string `json:"province"` // 省
	City     string `json:"city"`     // 市
	District string `json:"district"` // 区
	Street   string `json:"street"`   // 街道

	Rooms int `json:"rooms"` // 卧室数量

	Rent     float32 `json:"rent" gorm:"type:numeric(5,2)"`    // 每月租金
	PayCycle string  `json:"pay_cycle"`                        // 支付周期(月付、季付、年付)
	Deposit  float32 `json:"deposit" gorm:"type:numeric(5,2)"` // 押金

	Address string `json:"address"` // 房屋详细地址
	ImgUrl	string `json:"img_url"` // 房屋列表图片

	Score 		  float32  `json:"score" gorm:"numeric(3,1)"` // 评价得分
}

// HouseDetail 房子详情
type HouseDetail struct {
	Id       string `json:"id" gorm:"primaryKey"`
	HouseId  string `json:"house_id"`
	Name     string `json:"name"`     // 区域名称
	Describe string `json:"describe"` // 描述
	ImgPath string `json:"img_path"`  // 图片路径, 因为可能有多个,这里只存储图片的目录路径
}
