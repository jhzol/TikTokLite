package routes

import (
	"TikTokLite/common"
	"TikTokLite/controller"

	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine) *gin.Engine {
	douyin := r.Group("/douyin")
	{
		UserRoutes(douyin)
		PublishRoutes(douyin)
		CommentRoutes(douyin)
		FavoriteRoutes(douyin)
		RelationRoutes(douyin)
		douyin.GET("/feed/", common.AuthMiddleware(), controller.Feed)
	}

	return r
}
