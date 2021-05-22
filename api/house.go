package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"yizu/modules"
	"yizu/util"
)

// 评价管理

type HouseManager struct {

}

func (*HouseManager) List(c *gin.Context) {

}

func (*HouseManager) Delete(c *gin.Context) {

}

func (*HouseManager) Modify(c *gin.Context) {

	info := modules.House{}
	req, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(req, &info)
	if err != nil {
		c.Writer.Write(modules.ArgErr())
		return
	}

	db, e := yizuutil.GetDB()
	if e != nil {
		c.Writer.Write(modules.SysErr())
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

	c.Writer.Write(modules.Success())
	return
}