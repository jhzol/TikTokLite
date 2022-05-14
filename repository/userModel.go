package repository

import (
	"TikTokLite/log"
	"TikTokLite/util"
	"errors"
	"github.com/jinzhu/gorm"
)

type User struct {
	// gorm.Model
	Id       int64  `gorm:"column:user_id; primary_key;"`
	Name     string `gorm:"column:user_name"`
	PassWord string `gorm:"column:password"`
	Follow   int64  `gorm:"column:follow_count"`
	Follower int64  `gorm:"column:follower_count"`
	Token    string `gorm:"column:token"`
}

func (User) TableName() string {
	return "User"
}
func UserNameIsExist(userName string) error {
	db := GetDB()
	user := &User{}
	err := db.Where("user_name = ?", userName).Find(user).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("username exist")
	}
	return nil
}

func InsertUser(userName, passWord string) (*User, error) {
	db := GetDB()
	token := util.GetCurrentTimeForString()
	user := User{
		Name:     userName,
		PassWord: passWord,
		Token:    token,
	}
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Infof("regist user:%+v", user)
	return &user, nil
}
