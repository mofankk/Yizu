package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"yizu/modules"
	yizuutil "yizu/util"
)

type UserManager struct {

}

// Info 获取用户信息（资料页）
func (*UserManager) Info(c *gin.Context) {
	cookie, err := c.Cookie("session.id")
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, modules.SessionErr())
		return
	}

	redis := yizuutil.GetRedis()

	cache, err := redis.Get(redis.Context(), cookie).Result()
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, modules.SessionErr())
		return
	}
	var info modules.CacheInfo
	json.Unmarshal([]byte(cache), &info)

	db, err := yizuutil.GetDB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, modules.SysErr())
		return
	}
	var user modules.User
	err = db.Where(&modules.User{Id: info.UserId}).First(&user).Error
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, modules.SysErr())
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update 用户信息修改
// 注意: 修改手机号需要先进行短信验证码验证
func (*UserManager) Update(c *gin.Context) {
	var user modules.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
		return
	}

	db, err := yizuutil.GetDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.SysErr())
		return
	}

	err = db.Model(&modules.User{}).Where(&modules.User{Id: user.Id}).Updates(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.SysErr())
		return
	}

	c.JSON(http.StatusOK, modules.Success())
}

// 管理员相关

// List 用户列表
func (*UserManager) List(c *gin.Context) {
	db, err := yizuutil.GetDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.SysErr())
		return
	}

	var list []modules.User
	err = db.Find(&list).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.SysErr())
		return
	}

	res := modules.ResultInfo{}
	res.Data = list
	c.JSON(http.StatusOK, res)
}