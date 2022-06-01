package repository

import (
	"TikTokLite/common"
	"TikTokLite/log"
	"time"
)

type Comment struct {
	CommentId int64  `gorm:"column:comment_id; primary_key;"`
	UserId    int64  `gorm:"column:user_id"`
	VideoId   int64  `gorm:"column:video_id"`
	Comment   string `gorm:"column:comment"`
	Time      string `gorm:"column:time"`
}

func CommentAdd(userId, videoId int64, comment_text string) (*Comment, error) {
	db := common.GetDB()

	nowtime := time.Now().Format("01-02")
	comment := Comment{
		UserId:  userId,
		VideoId: videoId,
		Comment: comment_text,
		Time:    nowtime,
	}
	result := db.Create(&comment)

	if result.Error != nil {
		return nil, result.Error
	}
	log.Infof("comment:%+v", comment)
	return &comment, nil
}

func CommentDelete(comment_id int64) error {
	db := common.GetDB()
	commentTemp := Comment{}

	err := db.Where("comment_id = ?", comment_id).Take(&commentTemp).Error
	if err != nil {
		return err
	}
	db.Delete(&commentTemp)
	return nil

}

func CommentList(videoId int64) ([]Comment, error) {
	var comments []Comment
	db := common.GetDB()

	/* c := common.GetRE()
	values, _ := redis.Values(c.Do("lrange", videoId, "0", "-1"))
	for _,v := range values{

	} */

	err := db.Where("video_id = ?", videoId).Order("comment_id DESC").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil

}

// func GetVideoInfo(v interface{}) (*Video, error) {
// 	db := GetDB()
// 	video := Video{}
// 	err := db.Where("video_id = ?", v).Find(&video).Error
// 	if err != nil {
// 		return nil, errors.New("video error")
// 	}
// 	return &video, err
// }

// //根据user_id找到所有的用户信息
// func GetUser(v interface{}) (*User, error) {
// 	db := GetDB()
// 	var user User
// 	err := db.Where("user_id = ?", v).Find(&user).Error
// 	if err != nil {
// 		return nil, errors.New("user error")
// 	}
// 	return &user, nil

// }
