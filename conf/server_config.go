package conf

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type redisConfig struct {
	Size     int
	Address  string
	Password string
	DB       string
}

type postgresConfig struct {
	Username string
	Password string
	Address  string
	Port     string
	DBName   string
}

type Config struct {
	RdConfig     redisConfig
	PgConfig     postgresConfig
	HouseImgUrl  string
	HouseImgPath string
	AvatarUrl    string
}

var config *Config
var once sync.Once

func ServerConfig() *Config{

	once.Do(func() {
		config = &Config{}
		r := redisConfig{
			Size:     30,
			Address:  "152.136.114.51:923",
			Password: "Cx330$2021.@",
			DB:       "0",
		}
		p := postgresConfig{
			Username: "baitong",
			Password: "Cx330$2021.@",
			Address:  "152.136.114.51",
			Port:     "2237",
			DBName:   "yizu",
		}
		config.RdConfig = r
		config.PgConfig = p
		config.HouseImgUrl = filepath.Join("..", "h_image")
		config.HouseImgPath = filepath.Join("..", "h_image_path")
		config.AvatarUrl = filepath.Join("..", "avatar")

		//file, err := os.Open("server_config.json")
		//if err != nil {
		//	log.Errorf("载入配置文件出错: %v", err)
		//}
		//defer file.Close()
		//
		//jstr, _ := io.ReadAll(file)
		//err = json.Unmarshal(jstr, config)
		//if err != nil {
		//	log.Errorf("JSON解析失败: %v", err)
		//}
	})

	return config
}

func (*Config) SaveToFile() {
	file, _ := os.Create("server_config.json")
	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	enc.Encode(config)
}

func init() {
	cfg := ServerConfig()
	cfg.SaveToFile()
}
