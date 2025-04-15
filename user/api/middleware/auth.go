package middleware

import (
	"net/http"
	"user-service/util"

	"github.com/gin-gonic/gin"
)

func MustAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAuth := util.RegisterAuthorizedUser(c)
		if !isAuth {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		util.RegisterAuthorizedUser(c)
		c.Next()
	}
}
