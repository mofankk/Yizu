package modules

import (
	"github.com/google/uuid"
)

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

// 在线签约模块

// CacheInfo 缓存信息
type CacheInfo struct {
	HouseProvince string // 省
	HouseCity     string // 市
	HouseDistrict string // 区
	HouseStreet   string // 街道
}
