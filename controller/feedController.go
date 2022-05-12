package controller

import (
	"TikTokLite/log"
	"TikTokLite/proto/pkg"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

//点赞视频
func Feed(ctx *gin.Context) {
	feedResponse := message.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		NextTime:   time.Now().UnixNano(),
	}
	author := message.User{
		Id:       1,
		Name:     "test",
		IsFollow: false,
	}
	video := message.Video{
		Id:            1,
		Author:        &author,
		PlayUrl:       "http://112.74.109.70:36647/api/v1/buckets/video/objects/download?preview=true&prefix=MDQubXA0&version_id=null",
		CoverUrl:      "https://github.com/jhzol/test/blob/master/image/image-20220507170717302.png?raw=true",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
	feedResponse.VideoList = append(feedResponse.VideoList, &video)
	// data, _ := json.Marshal(feedResponse)
	log.Infof("message : [%+v]", feedResponse)
	ctx.JSON(200, feedResponse)
}
