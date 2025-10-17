package models

import (
	"golang_task4/utils"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content" form:"content"`
	UserID  uint   `json:"userID" form:"userID"`
	User    User   `json:"user" form:"user"`
	PostID  uint   `json:"postID" form:"postID"`
	Post    Post   `json:"post" form:"post"`
}

func (Comment) TableName() string {
	return "comments"
}

// 创建评论
func CreateComment(comment *Comment) *gorm.DB {
	return utils.DB.Create(*comment)
}

// 查询某文章的所有评论
func GetCommentList(postId uint) []*Comment {
	commentList := make([]*Comment, 10)
	utils.DB.Where("post_id = ?", postId).Find(&commentList)
	return commentList
}
