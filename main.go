package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"yizu/modules"
	"yizu/router"
	yizuutil "yizu/util"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	modules.SyncDB()
	gin.SetMode(gin.ReleaseMode)
	yizuutil.InitCasbin()
	router.Run()
}