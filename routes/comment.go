package routes

import (
	"TikTokLite/controller"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.RouterGroup) {
	comment := r.Group("comment")
	{
		comment.POST("/action/", controller.CommentAction)
		comment.GET("/list/", controller.GetCommentList)
	}

}
