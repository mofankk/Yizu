package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"yizu/api"
	"yizu/conf"
	"yizu/modules"
)

var Auth map[string]bool

func noNeedAuth() {
	Auth = make(map[string]bool)
	Auth["POST/register"] = true
	Auth["POST/login"] = true
}

func Run() {

	noNeedAuth()
	//u := api.UserManager{}
	hi := api.HiGin{}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 设置图片上传大小限制
	router.MaxMultipartMemory = 8 << 20  // 8 MiB

	// 设置中间件
	router.Use(SessionCheck())

	//router.Use(AuthRequired())


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

func SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if needAuth(c) {
			cookie, err := c.Cookie("session.id")
			if err != nil || cookie == ""{
				log.Debugf("未携带请求Session %v", err)
				c.JSON(http.StatusNetworkAuthenticationRequired, modules.SessionErr())
				c.Abort()
			}
		}
		c.Next()
	}
}

func needAuth(c *gin.Context) bool {
	url := c.Request.URL.Path
	method := c.Request.Method

	if Auth[method + url] {
		return false
	}
	return true
}