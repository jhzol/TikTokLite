package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UserLogin(ctx *gin.Context) {

}

func UserRegister(ctx *gin.Context) {

}

func GetUserInfo(ctx *gin.Context) {
	// var user message.User
	userId := ctx.Query("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
