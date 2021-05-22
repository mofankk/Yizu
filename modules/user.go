package modules

import "github.com/google/uuid"

// User 用户表
type User struct {
	Id       uuid.UUID `gorm:"primaryKey"`
	NickName string
	Phone    string
	Role     int
}