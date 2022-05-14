package main

import (
	"TikTokLite/log"
	"TikTokLite/repository"
	"TikTokLite/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	Init()
	defer repository.CloseDataBase()
	log.InitLog()
	defer log.Sync()
	r := gin.Default()
	r = routes.SetRoute(r)
	r.Run()
}

func Init() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()
	err := viper.ReadInConfig()
	fmt.Println(viper.Get("mysql.host"))
	if err != nil {
		log.Errorf("read config error:%+v", err)
		panic(err)
	}
	repository.InitDatabase()
}
