package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	if err == nil {
		return true
	}

	return false
}

func GetCacheInfo(c *gin.Context) (*modules.CacheInfo, bool){
	cookie, err := c.Cookie("session.id")
	if err != nil {
		log.Errorf("解析Cookie信息失败: %v", err)
		 return nil, false
	}

	redis := yizuutil.GetRedis()
	cache, err := redis.Get(redis.Context(), cookie).Result()
	if err != nil {
		log.Errorf("从Redis中获取Cookie失败: %v", err)
		return nil, false
	}
	var info modules.CacheInfo
	err = json.Unmarshal([]byte(cache), &info)
	if err != nil {
		log.Errorf("反序列化Session信息失败 %v", err)
		return nil, false
	}
	return &info, true
}
