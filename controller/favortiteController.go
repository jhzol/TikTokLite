package controller

import (
	"TikTokLite/response"
	"github.com/gin-gonic/gin"
)

//点赞视频
func FavoriteAction(ctx *gin.Context) {

}

//获取点赞列表
func GetFavoriteList(ctx *gin.Context) {
	response.Success(ctx, "success", nil)
}
