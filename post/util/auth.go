package util

import "github.com/gin-gonic/gin"

func RegisterUsername(c *gin.Context, username string) {
	c.Set("username", username)
}

func MustGetUsername(c *gin.Context) string {
	return c.MustGet("username").(string)
}

func TryGetUsername(c *gin.Context) (string, bool) {
	username, ok := c.Get("username")
	return username.(string), ok
}
