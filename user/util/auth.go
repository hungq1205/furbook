package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserId(ctx *gin.Context) (uint, bool) {
	username, ok := ctx.Get("userId")
	return username.(uint), ok
}

func MustGetUserId(ctx *gin.Context) uint {
	return ctx.MustGet("userId").(uint)
}

func RegisterAuthorizedUser(ctx *gin.Context) (bool, error) {
	userId := ctx.Request.Header.Get("X-USER-ID")
	userIdUint, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		return false, err
	}

	ctx.Set("userId", uint(userIdUint))
	return userId == "", nil
}
