package controller

import (
	"TikTokLite/log"
	"TikTokLite/proto/pkg"
	"TikTokLite/response"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context) {

}

func UserRegister(ctx *gin.Context) {
	var err error
	uesrName := ctx.Query("username")
	password := ctx.Query("password")
	registResponse := &message.DouyinUserRegisterResponse{}
	if len(uesrName) > 32 || len(password) > 32 {
		response.Fail(ctx, "username or password invalid", registResponse)
		return
	}
	registResponse, err = service.UserRegister(uesrName, password)
	if err != nil {
		log.Infof("registe error : %s", err)
		response.Fail(ctx, err.Error(), registResponse)
		return
	}
	response.Success(ctx, "success", registResponse)

}

func GetUserInfo(ctx *gin.Context) {
	// var user message.User

}
