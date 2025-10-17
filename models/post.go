package models

import (
	"golang_task4/utils"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title" form:"title"`
	Content string `gorm:"not null" json:"content" form:"content"`
	UserID  uint   `json:"userID" form:"userID"`
	User    User   `json:"user" form:"user"`
}

func (Post) TableName() string {
	return "posts"
}

func CreatePost(post *Post) *gorm.DB {
	return utils.DB.Create(&post)
}

func GetPostInfo(title string) Post {
	var p Post
	utils.DB.Where("title = ?", title).First(&p)
	return p
}

func UpdatePost(post *Post) *gorm.DB {
	return utils.DB.Model(&post).Updates(Post{
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
		User:    post.User,
	})
}

func DeletePost(post *Post) *gorm.DB {
	return utils.DB.Delete(&post)
}
