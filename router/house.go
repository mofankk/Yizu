package router

import (
	"github.com/gin-gonic/gin"
	"yizu/api"
)

func houseRouter(r *gin.Engine) {

	h := api.HouseManager{}

	// 房子
	g := r.Group("/house")
	g.Handle("GET", "/list", h.List)
	g.Handle("POST", "/add", h.Modify)
	g.Handle("POST", "/upimg", h.UploadImg)
	g.Handle("GET", "/detail", h.GetHouse) 	// 查看房子详情
	g.Handle("GET", "/detail", h.ScanHistory) 	// 查看房子详情
	g.Handle("DELETE", "/detail", h.DelScanHistory) 	// 查看房子详情
}