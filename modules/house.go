package modules

import "github.com/google/uuid"

// House 房子表
type House struct {
	Id         uuid.UUID `gorm:"primaryKey"`
	Name       string    `json:"name"`	// 房屋名称
	CreateTime string    `json:"create_time"`
	UpdateTime string    `json:"update_time"`
	DeleteTime string    `json:"-"`
	CreateUser uuid.UUID `json:"-"`		// 创建人

	Province string `json:"province"`	// 省
	City     string `json:"city"` 		// 市
	District string `json:"district"`	// 区

	Rooms    int 	`json:"rooms"`		// 房屋内总数量
}

// HouseDetail 房子详情
type HouseDetail struct {
	Id      uuid.UUID `gorm:"primaryKey"`
	HouseId uuid.UUID
}
