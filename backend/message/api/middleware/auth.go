package middleware

import (
	"message/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MustAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("X-Username")
		if username == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		util.RegisterUsername(c, username)
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("X-Username")
		util.RegisterUsername(c, username)
		c.Next()
	}
}
