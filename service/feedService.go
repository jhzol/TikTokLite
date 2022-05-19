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
	for i, video := range videoList {
		v := &message.Video{
			Id:            video.Id,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		}
		author, err := repository.GetUserInfo(video.AuthorId)
		if err != nil {
			return nil, err
		}
		v.Author = messageUserInfo(author)
		feed.VideoList[i] = v
	}
	nextTime := currentTime
	if len(videoList) != 0 {
		nextTime = videoList[len(videoList)-1].PublishTime
	}
	feed.NextTime = nextTime
	return feed, nil
}
