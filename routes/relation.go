package routes

import (
	"TikTokLite/controller"
	"github.com/gin-gonic/gin"
)

func RelationRoutes(r *gin.RouterGroup) {
	relation := r.Group("relation")
	{
		relation.POST("/action", controller.RelationAction)
		relation.GET("/follow/list", controller.GetFollowList)
		relation.GET("/follower/list", controller.GetFollowerList)
	}
}
