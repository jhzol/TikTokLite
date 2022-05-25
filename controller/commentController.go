package controller

import (
	"TikTokLite/log"
	"TikTokLite/response"
	"TikTokLite/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//发布评论
func CommentAction(ctx *gin.Context) {
	var err error
	token := ctx.Query("token")
	//user_id := ctx.Query("user_id")
	arr := strings.Split(token, "*")
	user_name := arr[0]
	//这里需要根据user_name来得到use_id才能进行插入数据，另写一个方法实现
	user, err := service.GetUserData(user_name)
	video_id := ctx.Query("video_id")
	action_type := ctx.Query("action_type")
	comment_text := ctx.Query("comment_text")
	comment_id := ctx.Query("comment_id")
	actionType, _ := strconv.Atoi(action_type)
	commentId, _ := strconv.Atoi(comment_id)
	userIdTemp := strconv.FormatInt(user.Id, 10)
	userId, _ := strconv.Atoi(userIdTemp)
	//userId, _ := strconv.Atoi(user_id)
	videoId, _ := strconv.Atoi(video_id)

	commentResponse, err := service.CommentAction(actionType, commentId, userId, videoId, comment_text)
	if err != nil {
		log.Infof("comment error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", commentResponse)
}

//获取评论列表
func GetCommentList(ctx *gin.Context) {
	var err error
	video_id := ctx.Query("video_id")
	token := ctx.Query("token")
	videoId, _ := strconv.ParseInt(video_id, 10, 64)

	listResponse, err := service.CommentList(token, videoId)
	if err != nil {
		log.Infof("list error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", listResponse)
}
