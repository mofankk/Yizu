package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"yizu/api"
	"yizu/conf"
)

func Run() {

	//u := api.UserManager{}
	s := api.ScanHistory{}
	hi := api.HiGin{}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 设置图片上传大小限制
	router.MaxMultipartMemory = 8 << 20  // 8 MiB

	//router.Use(AuthRequired())
	//{
	//	router.Handle("GET", "/house/list", h.List)
	//}


	// 登陆注册
	//router.Handle("POST", "/login", u.Login)
	//router.Handle("DELETE", "/logout", u.Logout)
	//router.Handle("DELETE", "/logoff", u.Logoff) //用户注销

	// 浏览历史
	router.Handle("GET", "/scan/list", s.List)
	router.Handle("DELETE", "/scan/del", s.Delete)

	router.Handle("GET", "higin", hi.Hello)

	houseRouter(router)

	log.Info("Yizu启动成功, 服务端口为: ", conf.ServerConfig().Port)
	router.Run(":" + conf.ServerConfig().Port)
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		x := c.Request.Method
		if x != "" {
			fmt.Println(x)
		} else {
			fmt.Println("认证失败")
		}
	}
}
