package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"yizu/modules"
	"yizu/service"
	"yizu/util"
)

// 房子管理模块

type HouseManager struct {

}

// List 获取房源列表
func (*HouseManager) List(c *gin.Context) {
	arg := &modules.HouseQueryArg{}
	if err := c.ShouldBind(arg); err != nil {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
		return
	}
	db, err := yizuutil.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.SysErr())
		return
	}
	s := service.HouseService{}
	info := s.QueryHouseList(arg, db)
	res := modules.QuerySuccess()
	res.Result = info
	c.JSON(http.StatusOK, res)
}

func (*HouseManager) Delete(c *gin.Context) {

}

func (*HouseManager) Modify(c *gin.Context) {

	info := modules.House{}
	req, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(req, &info)
	if err != nil {

		return
	}

	db, e := yizuutil.GetDB()
	if e != nil {
		return
	}

	if  info.Id == "" {
		err = db.Create(&info).Error
		if err != nil {
			c.Writer.Write(modules.InsertErr())
			return
		}
	} else {
		err = db.Omit("create_time", "delete_time").Updates(&info).Error
		if err != nil {
			c.Writer.Write(modules.UpdateErr())
			return
		}
	}

	return
}

// SetLocation 设置地理位置信息
func (*HouseManager) SetLocation(c *gin.Context) {
	type Acc struct {
		Province string `json:"province"` // 省
		City     string `json:"city"`     // 市
		District string `json:"district"` // 区
		Street   string `json:"street"`   // 街道
	}
	acc := &Acc{}
	if err := c.ShouldBind(acc); err != nil {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
		return
	}

}

// GetLocation 获取地理位置信息
// 在每次浏览房源的时候获取一下, 用于显示已配置过的信息
func (*HouseManager) GetLocation(c *gin.Context) {

}

