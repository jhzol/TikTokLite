package repository

import (
	"TikTokLite/util"
	"github.com/jinzhu/gorm"
)

type Video struct {
	Id            int64  `gorm:"column:video_id; primary_key;"`
	AuthorId      int64  `gorm:"column:author_id;"`
	PlayUrl       string `gorm:"column:play_url;"`
	CoverUrl      string `gorm:"column:cover_url;"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64  `gorm:"column:comment_count;"`
	PublishTime   int64  `gorm:"column:publish_time;"`
	Author        *User
}

func (Video) TableName() string {
	return "videos"
}

func InsertVideo(authorid int64, playurl, coverurl string) error {
	video := Video{
		AuthorId:      authorid,
		PlayUrl:       playurl,
		CoverUrl:      coverurl,
		FavoriteCount: 0,
		CommentCount:  0,
		PublishTime:   util.GetCurrentTime(),
	}
	db := GetDB()
	err := db.Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}

func GetVideoList(AuthorId int64) ([]Video, error) {
	var videos []Video
	db := GetDB()
	err := db.Where("author_id = ?", AuthorId).Find(&videos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return videos, err
	}
	return videos, nil
}

func GetVideoListByFeed(currentTime int64) ([]Video, error) {
	var videos []Video
	db := GetDB()
	err := db.Where("publish_time < ?", currentTime).Limit(30).Find(&videos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return videos, err
	}
	return videos, nil
}
