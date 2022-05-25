package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Relation struct {
	// gorm.Model
	Id       int64 `gorm:"column:relation_id; primary_key;"`
	Follow   int64 `gorm:"column:follow_id"`
	Follower int64 `gorm:"column:follower_id"`
}

func (Relation) TableName() string {
	return "relations"
}

func FollowAction(userId, toUserId int64) error {
	db := GetDB()
	relation := Relation{
		Follow:   userId,
		Follower: toUserId,
	}
	err := db.Where("follow_id = ? and follower_id = ?", userId, toUserId).Find(&Relation{}).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("you have followed this user")
	}
	err = db.Create(&relation).Error
	if err != nil {
		return err
	}
	return nil
}

func UnFollowAction(userId, toUserId int64) error {
	db := GetDB()
	err := db.Where("follow_id = ? and follower_id = ?", userId, toUserId).Delete(&Relation{}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetFollowList(userId int64, usertype string) ([]User, error) {
	db := GetDB()
	list := []User{}
	joinArg := "follower"
	if usertype == "follower" {
		joinArg = "follow"
	}
	err := db.Joins("left join relations on users.user_id = relations."+joinArg+"_id").
		Where("relations."+usertype+"_id = ?", userId).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		return []User{}, nil
	} else if err != nil {
		return nil, err
	}
	return list, nil
}
