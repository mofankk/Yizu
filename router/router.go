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
	hi := api.HiGin{}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 设置图片上传大小限制
	router.MaxMultipartMemory = 8 << 20  // 8 MiB

	//router.Use(AuthRequired())
	//{
	//	router.Handle("GET", "/house/list", h.List)
	//}


	router.Handle("GET", "higin", hi.Hello)

	// 路由注册
	sessionRouter(router)
	houseRouter(router)
	userRouter(router)
	commentRouter(router)


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
