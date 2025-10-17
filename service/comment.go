package service

import (
	"golang_task4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建评论
func CreateComment(c *gin.Context) {
	comment := models.Comment{}
	err := c.ShouldBind(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to bind post",
		})
		return
	} else {
		models.CreateComment(&comment)
		c.JSON(http.StatusOK, gin.H{
			"message": "Create Success",
		})
	}
}

// 查询评论集合
func GetCommentList(c *gin.Context) {
	var comment models.Comment
	data := make([]*models.Comment, 10)
	data = models.GetCommentList(comment.PostID)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
