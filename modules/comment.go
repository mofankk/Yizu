package modules

import (
	"github.com/google/uuid"
	"time"
)

// Comment 评论表
type Comment struct {
	Id         uuid.UUID `grom:"primaryKey"`
	HouseId    uuid.UUID
	UserId     uuid.UUID
	Text       string
	CreateTime time.Time
}