package modules

// Contract 签约模块
type Contract struct {
	Id         string `gorm:"primaryKey"`
	UserAId    string `json:"-"`
	UserBId    string `json:"-"`
	UserAName  string `json:"user_a_name"`
	UserBName  string `json:"user_b_name"`
	CreateTime string `json:"create_time"`
	StartTime  string `json:"start_time"` // 开始时间
	EndTime    string `json:"end_time"`   // 结束时间
	CardAId    string `json:"card_a_id"`  // 身份证号
	CardBId    string `json:"card_b_id"`  // 身份证号
	HouseId    string `json:"house_id"`   // 房屋ID
	RoomDesc   string `json:"room_desc"`  // 具体的房屋
}
