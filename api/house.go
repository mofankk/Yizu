package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"yizu/conf"
	"yizu/modules"
	"yizu/service"
	"yizu/util"
)

// HouseManager 房子管理模块
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
	res.Data = info
	c.JSON(http.StatusOK, res)
}

// Modify 新建或者更新房子信息
func (*HouseManager) Modify(c *gin.Context) {

	// 房子封面图
	file, err := c.FormFile("house_img")
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
		return
	}
	err = c.SaveUploadedFile(file, conf.ServerConfig().HouseImgUrl + "/" + file.Filename) // 需要将文件名入库
	if err != nil {
		c.JSON(http.StatusInternalServerError, modules.Failure())
		log.Errorf("图片保存失败: %v", err)
	} else {
		c.JSON(http.StatusOK, modules.Success())
	}
	// 解析房子其他信息
	info := modules.House{}
	req, _ := ioutil.ReadAll(c.Request.Body)
	err = json.Unmarshal(req, &info)
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
		return
	}

	db, e := yizuutil.GetDB()
	if e != nil {
		c.JSON(http.StatusInternalServerError, modules.SysErr())
		return
	}
	// 采用事务进行提交
	tx := db.Begin()
	flag := false
	if info.Id == "" {
		info.Id = uuid.New().String()
		err = db.Create(&info).Error
		if err != nil {
			flag = true
		}
	} else {
		err = db.Omit("create_time", "delete_time", "img_url").Updates(&info).Error
		if err != nil {
			flag = true
		}
	}
	err = tx.Where(modules.House{Id: info.Id}).Update("house_img", file.Filename).Error
	if err != nil {
		flag = true
	}
	if flag {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, modules.Failure())
	} else {
		tx.Commit()
		c.JSON(http.StatusOK, modules.Success())
	}
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

// GetHouse 查看房子详情
// 对应浏览房子详情页操作
// 要在浏览历史记录中更新
func (*HouseManager) GetHouse(c *gin.Context) {
	houseId := c.Query("id")
	// 参数检查
	if houseId == "" {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
		return
	}

	db, err := yizuutil.GetDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.SysErr())
		return
	}
	var house modules.House
	err = db.Where("id = ?", houseId).Find(&house).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.SysErr())
		return
	}
	res := modules.ResultInfo{}
	res.Data = house
	c.JSON(http.StatusOK, res)
}

// ScanHistory 获取浏览历史
func (*HouseManager) ScanHistory(c *gin.Context) {
	cacheInfo, ok := service.GetCacheInfo(c)
	if !ok {
		c.JSON(http.StatusBadRequest, modules.SessionErr())
		return
	}
	userId := cacheInfo.UserId

	db, err := yizuutil.GetDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.SysErr())
		return
	}
	var list []modules.HouseHistory
	err = db.Where("user_id = ?", userId).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.Failure())
		return
	}
	res := modules.ResultInfo{}
	res.Data = list
	c.JSON(http.StatusOK, res)
}

// DelScanHistory 删除浏览历史
func (*HouseManager) DelScanHistory(c *gin.Context) {
	houseId := c.Query("id")
	// 参数检查
	if houseId == "" {
		c.JSON(http.StatusBadRequest, modules.ArgErr())
		return
	}

	db, err := yizuutil.GetDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.SysErr())
		return
	}

	// 从列表查出来的houseId不存在在删除的时候查找失败的情况
	err = db.Model(&modules.HouseHistory{}).Delete("id = ?", houseId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, modules.Failure())
		return
	}
	c.JSON(http.StatusOK, modules.Success())
}