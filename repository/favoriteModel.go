package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Favorite struct {
	Id      int64 `gorm:"column:favorite_id; primary_key;"`
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

func (Favorite) TableName() string {
	return "favorites"
}

func LikeAction(uid, vid int64) error {
	db := GetDB()
	favorite := Favorite{
		UserId:  uid,
		VideoId: vid,
	}
	err := db.Where("user_id = ? and video_id = ?", uid, vid).Find(&Favorite{}).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("you have liked this video")
	}
	err = db.Create(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func UnLikeAction(uid, vid int64) error {
	db := GetDB()
	err := db.Where("user_id = ? and video_id = ?", uid, vid).Delete(&Favorite{}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetFavoriteList(uid int64) ([]Video, error) {
	var videos []Video
	db := GetDB()
	err := db.Joins("left join favorites on videos.video_id = favorites.video_id").
		Where("favorites.user_id = ?", uid).Find(&videos).Error
	if err == gorm.ErrRecordNotFound {
		return []Video{}, nil
	} else if err != nil {
		return nil, err
	}
	return videos, nil
}
