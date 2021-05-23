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
	g.Handle("DELETE", "/del", h.Delete)
	g.Handle("POST", "/upimg", h.UploadImg)

}