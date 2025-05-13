package middleware

import (
	"github.com/gin-gonic/gin"
	"message/util"
	"net/http"
)

func MustAuthorizeMiddleware() gin.HandlerFunc {
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

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("X-Username")
		util.RegisterUsername(c, username)
		c.Next()
	}
}
