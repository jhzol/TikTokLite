package main

import (
	"TikTokLite/log"
	"TikTokLite/repository"
	"TikTokLite/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	Init()
	defer repository.CloseDataBase()
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
	repository.InitDatabase()
	videoPath := viper.GetString("videofile")
	picPath := viper.GetString("picfile")
	os.Mkdir(videoPath, os.ModePerm)
	os.Mkdir(picPath, os.ModePerm)
}
