package service

import (
	"golang_task4/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建文章
func CreateArticle(c *gin.Context) {
	post := models.Post{}
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to bind post",
		})
		return
	} else {
		models.CreatePost(&post)
		c.JSON(http.StatusOK, gin.H{
			"message": "Create Success",
		})
	}
}

// 查询文章详情
func GetArticleInfo(c *gin.Context) {
	post := models.Post{}
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to bind post",
		})
	} else {
		data := models.GetPostInfo(post.Title)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

// 修改文章
func UpdateArticle(c *gin.Context) {
	post := models.Post{}
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to bind post",
		})
	} else {
		models.UpdatePost(&post)
		c.JSON(http.StatusOK, gin.H{
			"message": "Update Success",
		})
	}
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	post := models.Post{}
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to bind post",
		})
	} else {
		models.DeletePost(&post)
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete Success",
		})
	}
}
