package main

import (
	"github.com/gin-gonic/gin"
	_ "yizu/modules"
	"yizu/router"
)

func main() {
	//modules.SyncDB() // 注释掉这里是想测试下包多次引入是否多次执行init()
	gin.SetMode(gin.ReleaseMode)
	router.Run()
}