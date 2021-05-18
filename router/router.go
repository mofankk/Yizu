package router

import (
	"github.com/gin-gonic/gin"
	"yizu/handler"
)

func Run() {

	u := handler.UserManager{}
	h := handler.HouseManager{}
	s := handler.ScanHistory{}
	hi := handler.HiGin{}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// 登陆注册
	router.Handle("POST", "/login", u.Login)
	router.Handle("DELETE", "/logout", u.Logout)
	router.Handle("DELETE", "/logoff", u.Logoff) //用户注销

	// 房子
	router.Handle("GET", "/house/list", h.List)
	router.Handle("POST", "/house/add", h.Modify)
	router.Handle("DELETE", "/house/del", h.Delete)

	// 浏览历史
	router.Handle("GET", "/scan/list", s.List)
	router.Handle("DELETE", "/scan/del", s.Delete)

	router.Handle("GET", "higin", hi.Hello)

	router.Run(":2017")

}
