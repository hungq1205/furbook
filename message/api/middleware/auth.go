package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MustAuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("X-Username")
		if username == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("username", username)
		c.Next()
	}
}

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("X-Username")
		c.Set("username", username)
		c.Next()
	}
}
