package middleware

import (
	"net/http"
	"user/util"

	"github.com/gin-gonic/gin"
)

func MustAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAuth, err := util.RegisterAuthorizedUser(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
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
		_, err := util.RegisterAuthorizedUser(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
