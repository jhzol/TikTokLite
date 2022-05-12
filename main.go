package main

import (
	"TikTokLite/log"
	"TikTokLite/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	log.InitLog()
	defer log.Sync()
	r := gin.Default()
	r = routes.SetRoute(r)
	r.Run()
}
