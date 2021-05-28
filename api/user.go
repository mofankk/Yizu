package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yizu/modules"
	yizuutil "yizu/util"
)

type UserManager struct {

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