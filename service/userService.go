package service

import (
	"TikTokLite/log"
	"TikTokLite/proto/pkg"
	"TikTokLite/repository"
	"errors"
	"strconv"
	"sync"
)

var (
	currentUser sync.Map
)

func UserRegister(userName, password string) (*message.DouyinUserRegisterResponse, error) {
	err := repository.UserNameIsExist(userName)
	if err != nil {
		return nil, err
	}
	info, err := repository.InsertUser(userName, password)
	if err != nil {
		return nil, err
	}
	user := messageUserInfo(info)
	registResponse := &message.DouyinUserRegisterResponse{
		UserId: info.Id,
		Token:  info.Token,
	}
	currentUser.Store(info.Token, user)
	return registResponse, nil
}

func UserLogin(userName, password string) (*message.DouyinUserLoginResponse, error) {
	info, err := repository.GetUserInfo(userName)
	if err != nil {
		return nil, err
	}
	if password != info.Password {
		return nil, errors.New("password error")
	}
	loginResponse := &message.DouyinUserLoginResponse{
		UserId: info.Id,
		Token:  info.Token,
	}
	user := messageUserInfo(info)
	currentUser.Store(info.Token, user)
	log.Infof("login user:%+v", user)
	return loginResponse, nil
}

//获取登录用户的信息
func UserInfo(userID, token string) (*message.DouyinUserResponse, error) {
	info, err := CheckCurrentUser(token)
	if err != nil {
		return nil, err
	}
	if strconv.FormatInt(info.Id, 10) != userID {
		return nil, errors.New("token error")
	}
	return &message.DouyinUserResponse{User: info}, nil
}

func CheckCurrentUser(token string) (*message.User, error) {
	user, ok := currentUser.Load(token)
	if !ok {
		return nil, errors.New("please login your account or user doesn't exist")
	}
	return user.(*message.User), nil
}

func messageUserInfo(info *repository.User) *message.User {
	return &message.User{
		Id:              info.Id,
		Name:            info.Name,
		FollowCount:     info.Follow,
		FollowerCount:   info.Follower,
		IsFollow:        false,
		Avatar:          info.Avatar,
		BackgroundImage: info.BackgroundImage,
		Signature:       info.Signature,
	}
}
