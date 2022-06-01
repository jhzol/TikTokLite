package controller

import (
	"TikTokLite/common"
	"TikTokLite/response"
	"TikTokLite/service"
	"TikTokLite/util"
	"strconv"

	// "encoding/json"
	"github.com/gin-gonic/gin"
)

//视频流
func Feed(ctx *gin.Context) {
	var userId int64
	currentTime, err := strconv.ParseInt(ctx.Query("latest_time"), 10, 64)
	if err != nil {
		currentTime = util.GetCurrentTime()
	}
	token := ctx.Query("token")
	userId, err = common.VerifyToken(token)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	feedList, err := service.GetFeedList(currentTime, userId)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
	}
	response.Success(ctx, "success", feedList)
}
