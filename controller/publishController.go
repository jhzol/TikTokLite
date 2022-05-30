package controller

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"TikTokLite/service"
	"TikTokLite/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"path/filepath"
	"strconv"
)

//视频发布
func PublishAction(ctx *gin.Context) {
	// publishResponse := &message.DouyinPublishActionResponse{}
	token := ctx.PostForm("token")
	log.Infof("token:%s", token)
	data, err := ctx.FormFile("data")
	userId, err := util.VerifyToken(token)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	filename := filepath.Base(data.Filename)

	finalName := fmt.Sprintf("%s_%s", util.RandomString(), filename)
	videoPath := viper.GetString("videofile")
	saveFile := filepath.Join(videoPath, finalName)
	log.Debug(saveFile)
	if err := ctx.SaveUploadedFile(data, saveFile); err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	publish, err := service.PublishVideo(userId, saveFile)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", publish)

}

//获取视频列表
func GetPublishList(ctx *gin.Context) {
	token := ctx.Query("token")
	tokenUserId, err := util.VerifyToken(token)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	id := ctx.Query("user_id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
	}
	list, err := service.PublishList(tokenUserId, userId)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", list)
}
