package service

import (
	"encoding/json"
	"time"
	"yizu/modules"
	yizuutil "yizu/util"
)

func LoginSuccess(key string, user *modules.User) bool {

	// 生成Redis的key,也是session.id
	cacheInfo := modules.CacheInfo{}
	cacheInfo.RoleType = user.Role
	cacheInfo.UserId = user.Id
	cacheInfo.Username = user.Username
	cacheInfo.Password = user.Password

	b, _ := json.Marshal(cacheInfo)
	rdb := yizuutil.GetRedis()
	ctx := rdb.Context()
	err := rdb.SetEX(ctx, key, b, 43200 * time.Minute).Err()
	if err != nil {
		return true
	}

	return false
}
