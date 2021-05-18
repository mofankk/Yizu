package yizuutil

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {

	//dsn := "host=localhost user=mofan password=mofan2021.@ dbname=ruhe port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=152.136.114.51 user=baitong password=Cx330$2021.@ dbname=yizu port=2237 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("数据库连接失败： %v", err)
	}
	return db, err
}
