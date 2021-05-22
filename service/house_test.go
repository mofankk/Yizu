package service

import (
	log "github.com/sirupsen/logrus"
	"testing"
	"yizu/modules"
	yizuutil "yizu/util"
)

func TestHouseService_QueryHouseList(t *testing.T) {
	condition := modules.HouseQueryCondition{
		Province:   "山东",
		City:       "潍坊",
		District:   "安丘",
		Street:     "青云学府",
		StartTime:  "2020-11-11",
		EndTime:    "2021-05-22",
		RentUp:     100.7,
		RentLow:    50,
		UpdateTime: "",
		Page:       2,
		PageSize:   30,
	}
	h := HouseService{}
	db, err := yizuutil.GetDB()
	if err != nil {
		log.Error("数据库连接失败")
	}
	h.QueryHouseList(&condition, db)
}
