package routes

import (
	"TikTokLite/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.POST("/login/", controller.UserLogin)
		user.GET("/", controller.GetUserInfo)
		user.POST("/register/", controller.UserRegister)
	}

}
