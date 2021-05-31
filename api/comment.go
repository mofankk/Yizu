package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"gorm.io/gorm"
	"net/http"
	"yizu/modules"
	"yizu/service"
	yizuutil "yizu/util"
)

// 评价管理

type CommentManager struct {

}

// Delete 删除评论-只有管理员才能做
func (*CommentManager) Delete(c *gin.Context) {

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
	var house modules.House
	if db.Model(&modules.House{}).Where("id = ?", hc.HouseId).First(&house); house.Id == "" {
		c.JSON(http.StatusOK, modules.NoRecord())
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
// 每次更新都会对用户信息中的分数进行重新计算，基于用户表中的被打分次数
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
	// gorm.ErrRecordNotFound 这个问题要告知前端
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

// ListForHouse 获取房子的评论列表
// 这里需要进行鉴权
func (*CommentManager) ListForHouse(c *gin.Context) {
	cookie, err := c.Cookie("session.id")
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, modules.SessionErr())
		return
	}
	houseId := c.PostForm("house_id")

	redis := yizuutil.GetRedis()
	cacheInfo, err := redis.Get(redis.Context(), cookie).Result()
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, modules.SessionErr())
		return
	}
	var cache modules.CacheInfo
	json.Unmarshal([]byte(cacheInfo), &cache)

	ok := service.CheckScanPerm(cache.UserId, houseId)
	if !ok {
		c.JSON(http.StatusBadRequest, modules.AuthFail())
		log.Debugf("%s 没有查看 %s 的权限", cache.UserId, houseId)
		return
	}

	db, err := yizuutil.GetDB()
	if err  != nil {
		c.JSON(http.StatusBadRequest, modules.SysErr())
		return
	}
	var comments []modules.HouseComment
	// TODO 如果没有查到记录，这个Err会不会不为空
	err = db.Where(&modules.HouseComment{HouseId: houseId}).Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, modules.SysErr())
		return
	}
	res := modules.QuerySuccess()
	res.Data = comments
	c.JSON(http.StatusOK, res)
}

// ListForUser 获取用户的评论列表
// 这里需要进行鉴权
func (*CommentManager) ListForUser(c *gin.Context) {
	cookie, err := c.Cookie("session.id")
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, modules.SessionErr())
		return
	}
	userId := c.PostForm("user_id")

	redis := yizuutil.GetRedis()
	cacheInfo, err := redis.Get(redis.Context(), cookie).Result()
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, modules.SessionErr())
		return
	}
	var cache modules.CacheInfo
	json.Unmarshal([]byte(cacheInfo), &cache)

	ok := service.CheckScanPerm(cache.UserId, userId)
	if !ok {
		c.JSON(http.StatusBadRequest, modules.AuthFail())
		log.Debugf("%s 没有查看 %s 的权限", cache.UserId, userId)
		return
	}

	db, err := yizuutil.GetDB()
	if err  != nil {
		c.JSON(http.StatusBadRequest, modules.SysErr())
		return
	}
	var comments []modules.UserComment
	// 如果没有查到记录，这个Err会不会不为空, 答案是不会
	err = db.Where(&modules.UserComment{UserId: userId}).Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, modules.SysErr())
		return
	}
	res := modules.QuerySuccess()
	res.Data = comments
	c.JSON(http.StatusOK, res)
}

// 申请查看用户的评论列表

// 授权查看（可以与不可以）

// 申请查看房子的评论列表

// 查看房子的评论列表（房东自动允许)

// 获取需要授权的消息
