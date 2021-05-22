package modules

// UserComment 评论表
type UserComment struct {
	Id         string `grom:"primaryKey"`
	UserId     string `json:"user_id"`
	Text       string `json:"text"`
	Score      int    `json:"score"`
	CreateTime string `json:"create_time"`
	User 		User  `json:"-" gorm:"foreignKey:UserId"`
}

type HouseComment struct {
	Id      string `json:"id" gorm:"primaryKey"`
	Kitchen string `json:"kitchen"` // 对厨房的评价
}
