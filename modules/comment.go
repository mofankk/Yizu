package modules

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// UserComment 对用户的评价
type UserComment struct {
	Id         string `gorm:"primaryKey"`
	UserId     string `json:"user_id"`
	Text       string `json:"text"`  // 评价信息
	Score      int    `json:"score"` // 评价得分
	CreateTime string `json:"create_time"`
	user       User   `gorm:"foreignKey:UserId"`
}

// HouseComment 对房子的评价
type HouseComment struct {
	Id               string `json:"id" gorm:"primaryKey"`
	HouseId          string `json:"house_id"`
	Kitchen          string `json:"kitchen"`           // 对厨房的评价
	KitchenScore     int    `json:"kitchen_score"`     // 评价得分
	Toilet           string `json:"toilet"`            // 对厕所的评价
	ToiletScore      int    `json:"toilet_score"`      // 评价得分
	LivingRoom       string `json:"living_room"`       // 对客厅的评价
	LivingRoomScore  int    `json:"living_room_score"` // 评价得分
	BedRoom          string `json:"bed_room"`          // 对卧室的评价
	BedRoomScore     int    `json:"bed_room_score"`    // 评价得分
	Surrounding      string `json:"surrounding"`       // 对周围环境的评价
	SurroundingScore int    `json:"surrounding_score"` // 评价得分
	Score            int    `json:"score"`             // 总体得分
	CreateTime       string `json:"create_time"`       // 创建时间
	house            House  `gorm:"foreignKey:HouseId"`
}

func (*UserComment) TableName() string {
	return "user_comment"
}
func (*HouseComment) TableName() string {
	return "house_comment"
}

// BeforeCreate 钩子函数,用于生成UUID
func (h *HouseComment) BeforeCreate(tx *gorm.DB) (err error) {
	h.Id = uuid.New().String()
	nt := time.Now().Format("2006-01-02 15:04:05")
	h.CreateTime = nt
	return
}

// BeforeCreate 钩子函数,用于生成UUID
func (u *UserComment) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	nt := time.Now().Format("2006-01-02 15:04:05")
	u.CreateTime = nt
	return
}
