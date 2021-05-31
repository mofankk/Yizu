package yizuutil

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	log "github.com/sirupsen/logrus"
)

var e *casbin.Enforcer

func InitCasbin() {
	db, err := GetDB()
	if err != nil {
		log.Errorf("数据库连接失败: %v", err)
	}
	a, _ := gormadapter.NewAdapterByDB(db)
	e, _ = casbin.NewEnforcer("conf/model.conf", a)
}

func GetCasbin() *casbin.Enforcer {
	return e
}