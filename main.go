package main

import (
	"TikTokLite/common"
	"TikTokLite/log"
	"TikTokLite/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	Init()
	defer common.CloseDataBase()
	defer common.CloseReBase()
	defer log.Sync()
	r := gin.Default()
	r = routes.SetRoute(r)
	r.Run()
}

func Init() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("read config error:%+v", err)
		panic(err)
	}
	log.InitLog()
	common.InitDatabase()
	common.RedisInit()
	videoPath := viper.GetString("videofile")
	picPath := viper.GetString("picfile")
	os.Mkdir(videoPath, os.ModePerm)
	os.Mkdir(picPath, os.ModePerm)
}
