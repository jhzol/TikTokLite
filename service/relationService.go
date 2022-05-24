package service

import (
	"TikTokLite/log"
	"TikTokLite/proto/pkg"
	"TikTokLite/repository"
	"errors"
)

func RelationAction(toUserId int64, token, action string) error {
	userInfo, _ := CheckCurrentUser(token)
	uid := userInfo.Id
	if uid == toUserId {
		return errors.New("you can't follow yourself")
	}
	if action == "1" {
		log.Infof("follow action id:%v,toid:%v", uid, toUserId)
		err := repository.FollowAction(uid, toUserId)
		if err != nil {
			return err
		}
	} else {
		log.Infof("unfollow action id:%v,toid:%v", uid, toUserId)
		err := repository.UnFollowAction(uid, toUserId)
		if err != nil {
			return err
		}
	}
	return nil
}

func RelationFollowList(userId int64, token string) (*message.DouyinRelationFollowListResponse, error) {
	followList, err := repository.GetFollowList(userId, "follow")
	if err != nil {
		return nil, err
	}
	log.Infof("user:%v, followList:%+v", userId, followList)
	list, err := tokenFollowList(token)
	if err != nil {
		return nil, err
	}
	followListResponse := message.DouyinRelationFollowListResponse{
		UserList: make([]*message.User, len(followList)),
	}
	for i, u := range followList {
		follow := messageUserInfo(&u)
		if _, ok := list[follow.Id]; ok {
			follow.IsFollow = true
		}
		followListResponse.UserList[i] = follow
	}

	return &followListResponse, nil
}

func RelationFollowerList(userId int64, token string) (*message.DouyinRelationFollowerListResponse, error) {
	followList, err := repository.GetFollowList(userId, "follower")
	if err != nil {
		return nil, err
	}
	log.Infof("user:%v, followerList:%+v", userId, followList)
	list, err := tokenFollowList(token)
	if err != nil {
		return nil, err
	}
	followListResponse := message.DouyinRelationFollowerListResponse{
		UserList: make([]*message.User, len(followList)),
	}
	for i, u := range followList {
		follow := messageUserInfo(&u)
		if _, ok := list[follow.Id]; ok {
			follow.IsFollow = true
		}
		followListResponse.UserList[i] = follow
	}

	return &followListResponse, nil
}

func tokenFollowList(token string) (map[int64]struct{}, error) {
	m := make(map[int64]struct{})
	user, err := CheckCurrentUser(token)
	if err != nil {
		return m, nil
	}
	list, err := repository.GetFollowList(user.Id, "follow")
	if err != nil {
		return nil, err
	}
	for _, u := range list {
		m[u.Id] = struct{}{}
	}
	return m, nil
}
