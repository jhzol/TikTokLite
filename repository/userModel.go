package repository

import (
	"TikTokLite/log"
	"TikTokLite/util"
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type User struct {
	// gorm.Model
	Id       int64  `gorm:"column:user_id; primary_key;"`
	Name     string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Follow   int64  `gorm:"column:follow_count"`
	Follower int64  `gorm:"column:follower_count"`
	Token    string `gorm:"column:token"`
}

func (User) TableName() string {
	return "users"
}

//检查该用户名是否已经存在
func UserNameIsExist(userName string) error {
	db := GetDB()
	user := User{}
	err := db.Where("user_name = ?", userName).Find(&user).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("username exist")
	}
	return nil
}

//创建用户
func InsertUser(userName, password string) (*User, error) {
	db := GetDB()
	var builder strings.Builder
	builder.WriteString(userName)
	builder.WriteString(util.GetCurrentTimeForString())
	token := builder.String()
	user := User{
		Name:     userName,
		Password: password,
		Follow:   0,
		Follower: 0,
		Token:    token,
	}
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Infof("regist user:%+v", user)
	return &user, nil
}

//获取用户信息
func GetUserInfo(u interface{}) (*User, error) {
	db := GetDB()
	user := User{}
	var err error
	switch u.(type) {
	case int64:
		err = db.Where("user_id = ?", u).Find(&user).Error
	case string:
		err = db.Where("user_name = ?", u).Find(&user).Error
	default:
		err = errors.New("")
	}
	if err != nil {
		return nil, errors.New("user error")
	}
	return &user, nil
}
