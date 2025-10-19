package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 通过Jwt进行鉴权
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header 中取toke
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Authorization header is empty",
			})
		}
		c.Next()
	}
}
