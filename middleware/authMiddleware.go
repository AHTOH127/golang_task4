package middleware

import (
	"golang_task4/service"

	"github.com/gin-gonic/gin"
)

// 通过Jwt进行鉴权
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		service.Login(c)
		c.Next()
	}
}
