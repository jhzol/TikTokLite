package service

import (
	"TikTokLite/proto/pkg"
	"TikTokLite/repository"
)

func GetFeedList(currentTime int64) (*message.DouyinFeedResponse, error) {
	videoList, err := repository.GetVideoListByFeed(currentTime)
	if err != nil {
		return nil, err
	}
	feed := &message.DouyinFeedResponse{
		VideoList: make([]*message.Video, len(videoList)),
	}
	user, _ := CheckCurrentUser("testyser202205161336")
	for i, video := range videoList {
		video := &message.Video{
			Id:            video.Id,
			Author:        user,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		}
		feed.VideoList[i] = video
	}
	return feed, nil
}
