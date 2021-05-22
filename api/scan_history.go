package api

import (
	"github.com/gin-gonic/gin"
)

// 评价管理

type ScanHistory struct {

}

func (*ScanHistory) List(c *gin.Context) {

}

func (*ScanHistory) Delete(c *gin.Context) {

}

// Sync 点用户点开房子详情时，自动同步到浏览历史中
func (*ScanHistory) Sync(c *gin.Context) {

}