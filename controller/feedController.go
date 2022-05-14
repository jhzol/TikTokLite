package controller

import (
	"TikTokLite/proto/pkg"
	"TikTokLite/response"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

//点赞视频
func Feed(ctx *gin.Context) {
	//some test
	feedResponse := &message.DouyinFeedResponse{
		NextTime: time.Now().UnixNano(),
	}
	author := message.User{
		Id:       1,
		Name:     "test",
		IsFollow: false,
	}
	video := message.Video{
		Id:            1,
		Author:        &author,
		PlayUrl:       "http://112.74.109.70:9000/video/01.mp4",
		CoverUrl:      "https://github.com/jhzol/test/blob/master/image/image-20220507170717302.png?raw=true",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
	feedResponse.VideoList = append(feedResponse.VideoList, &video)
	// data, _ := json.Marshal(feedResponse)
	response.Fail(ctx, "feed error", feedResponse)
}
