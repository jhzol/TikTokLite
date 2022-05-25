package routes

import (
	"TikTokLite/controller"
	"github.com/gin-gonic/gin"
)

func FavoriteRoutes(r *gin.RouterGroup) {
	favorite := r.Group("favorite")
	{
		favorite.POST("/action/", controller.FavoriteAction)
		favorite.GET("/list/", controller.GetFavoriteList)
	}

}
