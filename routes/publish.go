package routes

import (
	"TikTokLite/controller"

	"github.com/gin-gonic/gin"
)

func PublishRoutes(r *gin.RouterGroup) {
	publish := r.Group("publish")
	{
		publish.POST("/action/", controller.PublishAction)
		publish.GET("/list/", controller.GetPublishList)
	}
}
