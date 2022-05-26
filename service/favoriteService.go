package service

import (
	"TikTokLite/log"
	message "TikTokLite/proto/pkg"
	"TikTokLite/repository"
)

func FavoriteAction(token string, vid int64, action int8) error {
	userInfo, _ := CheckCurrentUser(token)
	uid := userInfo.Id
	if action == 1 {
		log.Infof("like action uid:%v,vid:%v", uid, vid)
		err := repository.LikeAction(uid, vid)
		if err != nil {
			return err
		}
	} else {
		log.Infof("unlike action uid:%v,vid:%v", uid, vid)
		err := repository.UnLikeAction(uid, vid)
		if err != nil {
			return err
		}
	}
	return nil
}

func FavoriteList(token string, uid int64) (*message.DouyinFavoriteListResponse, error) {
	favList, err := repository.GetFavoriteList(uid)
	if err != nil {
		return nil, err
	}
	// log.Infof("user:%v, followList:%+v", uid, favList)

	favListResponse := message.DouyinFavoriteListResponse{
		VideoList: make([]*message.Video, len(favList)),
	}

	myFavList, _ := tokenFavList(token)
	log.Infof("my fav list: %+v", myFavList)

	for i, v := range favList {

		// u, err := repository.GetUserInfo(v.AuthorId)
		// if err != nil {
		// 	return nil, err
		// }
		user := messageUserInfo(&v.Author)

		video := &message.Video{
			Id:            v.Id,
			Author:        user,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
		}

		if _, ok := myFavList[video.Id]; ok {
			video.IsFavorite = true
		}

		favListResponse.VideoList[i] = video
	}

	return &favListResponse, nil
}

func tokenFavList(token string) (map[int64]struct{}, error) {
	m := make(map[int64]struct{})
	user, err := CheckCurrentUser(token)
	if err != nil {
		return m, nil
	}
	list, err := repository.GetFavoriteList(user.Id)
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		m[v.Id] = struct{}{}
	}
	return m, nil
}
