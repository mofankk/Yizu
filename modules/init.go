package modules

import (
	log "github.com/sirupsen/logrus"
	"time"
	"yizu/util"
)

// 与数据库同步表结构
func SyncDB() {
	for {
		db, err := yizuutil.GetDB()
		if err != nil {
			log.Errorf("PostgreSQL连接失败: %v", err)
		} else {
			log.Info("PostgreSQL连接成功")
			db.AutoMigrate(
				&User{},
				&House{},
				&HouseDetail{},
				&HouseHistory{},
				&UserComment{},
				&HouseComment{},
				)
			log.Info("PostgreSQL初始化成功")
			break
		}
		time.Sleep(10 * time.Second)
	}
}
