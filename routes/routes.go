package routes

import (
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
		douyin.GET("/feed/", controller.Feed)
	}

	return r
}
