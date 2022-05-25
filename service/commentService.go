package service

import (
	message "TikTokLite/proto/pkg"
	"TikTokLite/repository"

	"fmt"
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

func CommentAction(actionType, commentId, userId int64, videoId int64, comment_text string) (*message.DouyinCommentActionResponse, error) {
	commentResponse := &message.DouyinCommentActionResponse{}
	if actionType == 1 {
		//commentInfo, err := repository.CommentAdd(userId, videoId, comment_text)
		_, err := repository.CommentAdd(userId, videoId, comment_text)

		if err != nil {
			//return commentResponse, err

			return nil, err
		}
		//commentResponse.StatusCode = 0 //评论成功
		//commentResponse.StatusMsg = "comment success!"
		//fmt.Printf("commentResponse:%v", commentResponse)  comment:{CommentId:6 UserId:29 VideoId:6 Comment:哈哈哈哈
		//	Time:2022-05-24 17:36:52}
		return commentResponse, nil
	} else if actionType == 2 {
		err := repository.CommentDelete(commentId)
		if err != nil {
			return commentResponse, err
		}
		//commentResponse.StatusCode = 0
		//commentResponse.StatusMsg = "comments deleted successfully!"
		return commentResponse, nil
	}

	//info, err := repository.InsertUser(userName, password)
	//if err != nil {
	//	return registResponse, err
	//}

	//commentResponse.StatusCode = 1 //评论失败
	//commentResponse.StatusMsg = "comment failure!"
	return commentResponse, nil
}

func messageCommentInfo(info *repository.Comment) *message.Comment {
	userInfo, err := repository.GetUserInfo(info.UserId)
	userTemp := messageUserInfo(userInfo)
	if err != nil {
		return nil
	}
	return &message.Comment{
		Id:         info.CommentId,
		User:       userTemp,
		Content:    info.Comment,
		CreateDate: info.Time,
	}
}

//下面这段就是根据用户name(user_name)找到用户ID(user_id)
func GetUserData(user_name string) (*User, error) {
	info, err := repository.GetUserInfo(user_name)
	if err != nil {
		return nil, err
	}

	user := searchUserData(info)
	return &user, nil
}
func searchUserData(info *repository.User) User {
	return User{
		Id:       info.Id,
		Name:     info.Name,
		Follow:   info.Follow,
		Follower: info.Follower,
		Token:    info.Token,
	}
}

//用户评论
func CommentList(token string, videoId int64) (*message.DouyinCommentListResponse, error) {

	fmt.Println("--------------------------------------------")

	v, err := repository.GetVideoInfo(videoId)
	if err != nil {
		return nil, err
	}

	video := messageVideoInfo(v)

	comments, err := repository.CommentList(video.Id)
	fmt.Printf("comments:%v\n", comments)

	if err != nil {
		return nil, err
	}

	list := &message.DouyinCommentListResponse{
		CommentList: make([]*message.Comment, len(comments)),
	}

	for i, comment := range comments {
		//为了找到video_id所对应的user_id，在通过user_id找到user_name.传递给前端
		userID := comment.UserId
		user, _ := repository.GetUser(userID)
		users := messageUserInfo(user)
		/* 	if err != nil {
			return nil, err
		}
		users := &user{
			Id:   user.Id,
			Name: user.Name,
		} */

		v := &message.Comment{
			Id:         comment.CommentId,
			User:       users,
			Content:    comment.Comment,
			CreateDate: comment.Time,
		}
		list.CommentList[i] = v
	}

	return list, nil

}

func messageVideoInfo(info *repository.Video) *message.Video {

	return &message.Video{
		Id:            info.Id,
		PlayUrl:       info.PlayUrl,
		CoverUrl:      info.CoverUrl,
		FavoriteCount: info.FavoriteCount,
		CommentCount:  info.CommentCount,
		IsFavorite:    false,
	}
}
