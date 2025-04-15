package util

import "github.com/gin-gonic/gin"

func GetUsername(ctx *gin.Context) (string, bool) {
	username, ok := ctx.Get("username")
	return username.(string), ok
}

func MustGetUsername(ctx *gin.Context) string {
	return ctx.MustGet("username").(string)
}

func RegisterAuthorizedUser(ctx *gin.Context) bool {
	username := ctx.Request.Header.Get("X-Username")
	ctx.Set("username", username)
	return username == ""
}
