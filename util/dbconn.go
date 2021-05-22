package yizuutil

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/boj/redistore.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"yizu/conf"
)

var store *redistore.RediStore
var once sync.Once

func GetDB() (*gorm.DB, error) {

	//dsn := "host=152.136.114.51 user=baitong password=Cx330$2021.@ dbname=yizu port=2237 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=" + conf.ServerConfig().PgConfig.Address + " " +
		"user=" + conf.ServerConfig().PgConfig.Username + " " +
		"password=" + conf.ServerConfig().PgConfig.Password + " " +
		"dbname=" + conf.ServerConfig().PgConfig.DBName + " " +
		"port=" + conf.ServerConfig().PgConfig.Port + " " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("PostgreSQL连接失败: %v", err)
	}
	return db, err
}

func init() {
	once.Do(func() {
		var err error
		store, err = redistore.NewRediStoreWithDB(
			conf.ServerConfig().RdConfig.Size,
			"tcp",
			conf.ServerConfig().RdConfig.Address,
			conf.ServerConfig().RdConfig.Password,
			conf.ServerConfig().RdConfig.DB,
		)
		if err != nil {
			log.Errorf("Redis创建失败: %v", err)
		}
	})
}

func GetRedisStore() *redistore.RediStore {
	return store
}
