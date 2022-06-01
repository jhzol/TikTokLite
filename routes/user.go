package routes

import (
	"TikTokLite/controller"
	"TikTokLite/util"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.POST("/login/", controller.UserLogin)
		user.GET("/", util.AuthMiddleware(), controller.GetUserInfo)
		//user.GET("/", controller.GetUserInfo)
		user.POST("/register/", controller.UserRegister)
	}

}
