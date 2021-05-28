package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"yizu/modules"
	"yizu/service"
	yizuutil "yizu/util"
)

// 评价管理

type CommentManager struct {

}

// List 获取评论列表
// 用户浏览到房子，查看完详情后，可以向房东请求展示评论详情
func (*CommentManager) List(c *gin.Context) {

}

// Update 更新评论
func (*CommentManager) Update(c *gin.Context) {

}

// Delete 删除评论-只有管理员才能做
func (*CommentManager) Delete(c *gin.Context) {

}

// Add 新增评论
func (*CommentManager) Add(c *gin.Context) {

}

// CommentForHouse 对房子进行评价
func (*CommentManager) CommentForHouse(c *gin.Context) {
	var hc modules.HouseComment
	err := c.ShouldBind(&hc)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
		return
	}
	// TODO 考虑使用遍历的方式实现
	//t := reflect.TypeOf(hc)
	//v := reflect.ValueOf(hc)
	//for i := 0; i < t.NumField(); i++ {
	//	x := v.Field(i).Interface();
	//
	//	if reflect.TypeOf(x) == reflect.Type.NumIn() {
	//
	//	}
	//}

	score := service.CountHouseScore(&hc)

	db, err := yizuutil.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.SysErr())
		return
	}
	tx := db.Begin()
	err = tx.Create(&hc).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.Failure())
		tx.Rollback()
		return
	}
	err = tx.Model(&modules.House{}).Where("id = ?", hc.HouseId).Update("score", score).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.Failure())
		tx.Rollback()
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, modules.Success())
}

// CommentForUser 用户进行评价
func (*CommentManager) CommentForUser(c *gin.Context) {
	var uc modules.UserComment
	err := c.ShouldBind(&uc)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
		return
	}
	// TODO 考虑使用遍历的方式实现
	//t := reflect.TypeOf(hc)
	//v := reflect.ValueOf(hc)
	//for i := 0; i < t.NumField(); i++ {
	//	x := v.Field(i).Interface();
	//
	//	if reflect.TypeOf(x) == reflect.Type.NumIn() {
	//
	//	}
	//}

	db, err := yizuutil.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.SysErr())
		return
	}
	tx := db.Begin()
	var user modules.User
	err = tx.Where(&modules.User{Id: uc.UserId}).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, modules.NoRecord())
		tx.Rollback()
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, modules.Failure())
		tx.Rollback()
		return
	}
	err = tx.Create(&uc).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.Failure())
		tx.Rollback()
		return
	}
	user.Times += 1
	user.Score = (user.Score + float32(uc.Score)) / float32(user.Times)
	err = tx.Model(&modules.User{}).Where("id = ?", uc.UserId).
		Updates(&modules.User{Times: user.Times, Score: user.Score}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.Failure())
		tx.Rollback()
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, modules.Success())
}