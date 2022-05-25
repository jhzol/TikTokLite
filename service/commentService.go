package service

import (
	"TikTokLite/log"
	"TikTokLite/proto/pkg"
	"TikTokLite/repository"
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

func CommentAction(commentId, videoId int64, token, comment_text, actionType string) (*message.DouyinCommentActionResponse, error) {
	userinfo, _ := CheckCurrentUser(token)
	userId := userinfo.Id

	if actionType == "1" {
		//commentInfo, err := repository.CommentAdd(userId, videoId, comment_text)
		commentInfo, err := repository.CommentAdd(userId, videoId, comment_text)

		if err != nil {
			return nil, err
		}
		user, _ := repository.GetUserInfo(userId)
		commentResponse := &message.DouyinCommentActionResponse{
			Comment: &message.Comment{
				Id:         commentInfo.CommentId,
				User:       messageUserInfo(user),
				Content:    comment_text,
				CreateDate: commentInfo.Time,
			},
		}

		return commentResponse, nil
	} else {
		err := repository.CommentDelete(commentId)
		if err != nil {
			return nil, err
		}
		commentResponse := &message.DouyinCommentActionResponse{
			Comment: nil,
		}
		return commentResponse, nil
	}

}

// //下面这段就是根据用户name(user_name)找到用户ID(user_id)
// func GetUserData(user_name string) (*User, error) {
// 	info, err := repository.GetUserInfo(user_name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user := searchUserData(info)
// 	return &user, nil
// }
// func searchUserData(info *repository.User) User {
// 	return User{
// 		Id:       info.Id,
// 		Name:     info.Name,
// 		Follow:   info.Follow,
// 		Follower: info.Follower,
// 		Token:    info.Token,
// 	}
// }

//用户评论
func CommentList(token string, videoId int64) (*message.DouyinCommentListResponse, error) {
	comments, err := repository.CommentList(videoId)
	if err != nil {
		return nil, err
	}
	log.Infof("comments:%v\n", comments)

	list := &message.DouyinCommentListResponse{
		CommentList: make([]*message.Comment, len(comments)),
	}

	for i, comment := range comments {
		//为了找到video_id所对应的user_id，在通过user_id找到user_name.传递给前端
		userID := comment.UserId
		user, _ := repository.GetUserInfo(userID)
		userinfo := messageUserInfo(user)

		v := &message.Comment{
			Id:         comment.CommentId,
			User:       userinfo,
			Content:    comment.Comment,
			CreateDate: comment.Time,
		}
		list.CommentList[i] = v
	}

	return list, nil

}
