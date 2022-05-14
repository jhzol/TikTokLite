package service

import (
	"TikTokLite/proto/pkg"
	"TikTokLite/repository"
)

func UserRegister(userName, passWord string) (*message.DouyinUserRegisterResponse, error) {
	registResponse := &message.DouyinUserRegisterResponse{}
	err := repository.UserNameIsExist(userName)
	if err != nil {
		return registResponse, err
	}
	user, err := repository.InsertUser(userName, passWord)
	if err != nil {
		return registResponse, err
	}
	registResponse.UserId = user.Id
	registResponse.Token = user.Token
	return registResponse, nil
}
