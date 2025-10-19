package router

import (
	"golang_task4/middleware"
	"golang_task4/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 全局鉴权中间件
	r.Use(middleware.Auth())

	common := r.Group("/common")
	{
		// 注册
		common.POST("/register", service.Register)

		// 登录
		common.POST("/login", service.Login)
	}

	article := r.Group("/article")
	{
		// 创建文章
		article.POST("/createArticle", service.CreateArticle)

		// 获取文章详情
		article.POST("/getArticle", service.GetArticleInfo)

		// 修改文章
		article.POST("/updateArticle", service.UpdateArticle)

		// 删除文章
		article.POST("/deleteArticle", service.DeleteArticle)
	}

	comment := r.Group("/comment")
	{
		// 创建评论
		comment.POST("/createComment", service.CreateComment)

		// 查询某文章下的所有评论
		comment.POST("/getCommentList", service.GetCommentList)
	}

	return r
}
