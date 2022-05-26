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

// 联合查询结构
type FavoriteVideo struct {
	Id            int64  `gorm:"column:video_id; primary_key;"`
	AuthorId      int64  `gorm:"column:author_id;"`
	PlayUrl       string `gorm:"column:play_url;"`
	CoverUrl      string `gorm:"column:cover_url;"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64  `gorm:"column:comment_count;"`
	PublishTime   int64  `gorm:"column:publish_time;"`
	Author        User   `gorm:"foreignKey:AuthorId;"` // 注意设置外键 AuthorId
	// Author        *User
}

func (Favorite) TableName() string {
	return "favorites"
}

func (FavoriteVideo) TableName() string {
	return "videos"
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

func GetFavoriteList(uid int64) ([]FavoriteVideo, error) {
	var videos []FavoriteVideo
	db := GetDB()
	err := db.Preload("Author").
		Joins("left join favorites on videos.video_id = favorites.video_id").
		Where("favorites.user_id = ?", uid).Find(&videos).Error
	if err == gorm.ErrRecordNotFound {
		return []FavoriteVideo{}, nil
	} else if err != nil {
		return nil, err
	}
	return videos, nil
}
