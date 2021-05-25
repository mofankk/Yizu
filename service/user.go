package service

import (
	"encoding/json"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"yizu/modules"
	yizuutil "yizu/util"
)

func RegisterUser(user *modules.RegistInfo) bool {
	b, _ := json.Marshal(user)
	var userInfo modules.User
	err := json.Unmarshal(b, &userInfo)
	if err != nil {
		log.Errorf("注册信息解析失败")
		return false
	}
	userInfo.Id = uuid.New().String()
	db, err := yizuutil.GetDB()
	if err != nil {
		log.Errorf("数据库创建失败")
		return false
	}
	err = db.Create(userInfo).Error
	if err != nil {
		log.Errorf("创建用户信息失败: %v", err)
		return false
	}
	return true
}
