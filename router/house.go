package router

import (
	"github.com/gin-gonic/gin"
	"yizu/api"
)

func houseRouter(r *gin.Engine) {

	h := api.HouseManager{}

	// 房子
	g := r.Group("/house")
	g.Handle("GET", "/list", h.List)                // 获取房源列表
	g.Handle("POST", "/add", h.Modify)              // 新建或修改房屋信息
	g.Handle("GET", "/detail", h.GetHouse)          // 查看房子详情
	g.Handle("GET", "/history", h.ScanHistory)       // 查看浏览历史
	g.Handle("DELETE", "/history", h.DelScanHistory) // 删除浏览历史
}
