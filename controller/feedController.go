package controller

import (
	"TikTokLite/response"
	"TikTokLite/service"
	"TikTokLite/util"
	"strconv"
	// "encoding/json"
	"github.com/gin-gonic/gin"
)

//视频流
func Feed(ctx *gin.Context) {
	currentTime, err := strconv.ParseInt(ctx.Query("latest_time"), 10, 64)
	if err != nil {
		currentTime = util.GetCurrentTime()
	}
	feedList, err := service.GetFeedList(currentTime)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
	}
	response.Success(ctx, "success", feedList)
}
