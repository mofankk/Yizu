package main

import (
	"github.com/gin-gonic/gin"
	_ "yizu/modules"
	"yizu/router"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router.Run()
}