package modules

import (
	log "github.com/sirupsen/logrus"
	"time"
	"yizu/util"
)

func init() {
	for {
		db, err := yizuutil.GetDB()
		if err != nil {
			log.Errorf("数据库连接失败： %v", err)
		} else {
			log.Info("数据库连接成功")

			//更新数据库表
			if e := db.AutoMigrate(House{}); e != nil {
				log.Errorf("更新表'house'失败: %v", e)
			}

			break
		}
		time.Sleep(10 * time.Second)
	}
}
