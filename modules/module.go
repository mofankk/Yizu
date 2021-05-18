package modules

import (
	"github.com/google/uuid"
	"time"
)

// House 房子表
type House struct {
	Id         uuid.UUID `gorm:"primaryKey"`
	Name       string    `json:"name"`
	CreateTime string    `json:"create_time"`
	UpdateTime string    `json:"update_time"`
	DeleteTime string    `json:"-" gorm:"default:'null'"`
	CreateUser uuid.UUID `json:"-" gorm:"default:'null'"`

	Province string `json:"province" gorm:"default:'null'"`
	City     string `json:"city" gorm:"default:'null'"`
	District string `json:"district" gorm:"default:'null'"`
}

// HouseDetail 房子详情
type HouseDetail struct {
	Id      uuid.UUID `gorm:"primaryKey"`
	HouseId uuid.UUID
}

// User 用户表
type User struct {
	Id       uuid.UUID `gorm:"primaryKey"`
	NickName string
	Phone    string
	Role     int
}

// Comment 评论表
type Comment struct {
	Id         uuid.UUID `grom:"primaryKey"`
	HouseId    uuid.UUID
	UserId     uuid.UUID
	Text       string
	CreateTime time.Time
}

// 历史表

// Role 角色表
type Role struct {
	Id   uuid.UUID `grom:"primaryKey"`
	Type int       // 1-管理员 2-普通浏览用户 3-房东 4-房东认证后的用户
	Name string
}

// 权限表

// 用户动态表（发表的动态）

// 聊天功能 咨询
