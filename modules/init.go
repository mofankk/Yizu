package modules

import (
	log "github.com/sirupsen/logrus"
	"time"
	"yizu/util"
)

// 与数据库同步表结构
func Init() {
	for {
		db, err := yizuutil.GetDB()
		if err != nil {
			log.Errorf("%s%v", DatabaseConnErr, err)
		} else {
			log.Info("数据库连接成功")
			db.AutoMigrate(
				&User{},
				&House{},
				&HouseDetail{},
				&HouseHistory{},
				&UserComment{},
				&HouseComment{},
				)

			break
		}
		time.Sleep(10 * time.Second)
	}
}
