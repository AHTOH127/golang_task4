package service

import (
	"golang_task4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建文章
func CreateArticle(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	models.CreatePost(&post)
	c.JSON(http.StatusCreated, gin.H{"message": "create success"})
}

// 查询文章详情
func GetArticleInfo(c *gin.Context) {
	post := models.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	data := models.GetPostInfo(post.Title)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// 修改文章
func UpdateArticle(c *gin.Context) {
	post := models.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	models.UpdatePost(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	post := models.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	models.DeletePost(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})
}
