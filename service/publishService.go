package service

import (
	"TikTokLite/minioStore"
	"TikTokLite/proto/pkg"
	"TikTokLite/repository"
	"strconv"
)

func PublishVideo(token, saveFile string) (*message.DouyinPublishActionResponse, error) {
	user, _ := CheckCurrentUser(token)
	userId := user.Id
	client := minioStore.NewMinioClient()
	videourl, err := client.UploadFile("video", saveFile, strconv.FormatInt(userId, 10))
	if err != nil {
		return nil, err
	}
	err = repository.InsertVideo(userId, videourl, "https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/7909abe413ec4a1e82032d2beb810157~tplv-k3u1fbpfcp-zoom-in-crop-mark:1304:0:0:0.awebp?")
	if err != nil {
		return nil, err
	}
	return &message.DouyinPublishActionResponse{}, nil
}

func PublishList(token string, userId int64) (*message.DouyinPublishListResponse, error) {
	u, err := repository.GetUserInfo(userId)
	if err != nil {
		return nil, err
	}
	user := messageUserInfo(u)
	videos, err := repository.GetVideoList(user.Id)
	if err != nil {
		return nil, err
	}
	list := &message.DouyinPublishListResponse{
		VideoList: make([]*message.Video, len(videos)),
	}
	followlist, _ := tokenFollowList(token)
	follow := false
	if _, ok := followlist[user.Id]; ok {
		follow = true
	}
	user.IsFollow = follow
	for i, video := range videos {
		v := &message.Video{
			Id:            video.Id,
			Author:        user,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		}
		list.VideoList[i] = v
	}
	return list, nil
}
