package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"yizu/modules"
	yizuutil "yizu/util"
)

// Create 创建合同
// 自动后台绑定用户ID
func Create(c *gin.Context) {
	var con modules.Contract
	err := c.ShouldBind(&con)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, modules.ArgErr())
		return
	}
	db, err := yizuutil.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.SysErr())
		return
	}
	redis := yizuutil.GetRedis()
	cookie, err := c.Cookie("session.id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.SysErr())
		return
	}
	ca, err := redis.Get(redis.Context(), cookie).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.Failure())
		return
	}
	var info modules.CacheInfo
	json.Unmarshal([]byte(ca), &info)

	if info.RoleType == 2 {
		con.UserAId = info.UserId

	} else if info.RoleType == 3 {
		con.UserBId = info.UserId
	}
	err = db.Create(&con).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.Failure())
		return
	}
	c.JSON(http.StatusOK, modules.Success())
}
