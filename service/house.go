package service

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"yizu/modules"
)

type HouseService struct {

}

// QueryHouseList 按照条件查询房源列表
func (*HouseService) QueryHouseList(h *modules.HouseQueryArg, db *gorm.DB) []modules.House {

	// 拼接查询条件
	db = db.Where("1 = 1")
	if h.Province != "" {
		db = db.Where(modules.House{Province: h.Province})
	}
	if h.City != "" {
		db = db.Where(modules.House{City: h.City})
	}
	if h.District != "" {
		db = db.Where(modules.House{District: h.District})
	}
	if h.Street != "" {
		db = db.Where(modules.House{Street: h.Street})
	}
	if h.RentUp != 0 {
		db = db.Where("rent <= ?", h.RentUp)
	}
	if h.RentLow != 0 {
		db = db.Where("rent >= ?", h.RentLow)
	}
	if h.StartTime != "" {
		db = db.Where("create_time >= ?", h.StartTime)
	}
	if h.EndTime != "" {
		db = db.Where("create_time <= ?", h.EndTime)
	}

	// 分页查询
	if h.Page > 0 && h.PageSize > 0 {
		db = db.Offset((h.Page - 1) * h.PageSize).Limit(h.PageSize)
	} else { // 默认查询 1 页，每页 30 条
		db = db.Offset(0).Limit(30)
	}

	// 记录排序，默认最新更新的在前面
	if h.UpdateTime == "asc" {
		db = db.Order("update_time ASC")
	} else {
		db = db.Order("update_time DESC")
	}

	var list []modules.House
	err := db.Find(&list).Error
	if err != nil {
		log.Errorf("数据库查询失败: %v", err)
	}
	return list
}
