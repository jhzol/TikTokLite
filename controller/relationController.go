package controller

import (
	"TikTokLite/response"
	"TikTokLite/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

//关注操作
func RelationAction(ctx *gin.Context) {
	token := ctx.Query("token")
	toUserId := ctx.Query("to_user_id")
	touid, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	action := ctx.Query("action_type")
	err = service.RelationAction(touid, token, action)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", nil)
}

//获取关注列表
func GetFollowList(ctx *gin.Context) {
	token := ctx.Query("token")
	UserId := ctx.Query("user_id")
	uid, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	followList, err := service.RelationFollowList(uid, token)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", followList)
}

//获取关注者列表
func GetFollowerList(ctx *gin.Context) {
	token := ctx.Query("token")
	UserId := ctx.Query("user_id")
	uid, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	followerList, err := service.RelationFollowerList(uid, token)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", followerList)
}
