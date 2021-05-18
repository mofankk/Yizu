package test

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
	"yizu/modules"
)

func TestDBConn(t *testing.T) {
	db, err := modules.GetDB()
	if err != nil {
		log.Error("数据库连接失败")
	}
	var ctx context.Context
	fmt.Println("hello,postgresql")
	db.Logger.Info(ctx, "hello")
}

