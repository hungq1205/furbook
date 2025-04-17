package user

import (
	"user/api/middleware"
	"user/usecase/friend"
	"user/usecase/user"

	"github.com/gin-gonic/gin"
)

func MakeHandler(app *gin.Engine, userService user.UseCase, friendService friend.UseCase) {
	userGroup := app.Group("/api/user")
	{
		userGroup.Use(middleware.MustAuthMiddleware())

		userGroup.GET("/:username", func(ctx *gin.Context) {
			GetUser(ctx, userService, friendService)
		})

		userGroup.GET("", func(ctx *gin.Context) {
			GetUserList(ctx, userService, friendService)
		})

		userGroup.POST("", func(ctx *gin.Context) {
			CreateUser(ctx, userService, friendService)
		})

		userGroup.PUT("", func(ctx *gin.Context) {
			UpdateUser(ctx, userService, friendService)
		})

		userGroup.DELETE("", func(ctx *gin.Context) {
			DeleteUser(ctx, userService)
		})

		userGroup.GET("/:username/exists", func(ctx *gin.Context) {
			CheckUsernameExists(ctx, userService)
		})

		// friend

		userGroup.GET("/:username/friends/:friend/exists", func(ctx *gin.Context) {
			CheckFriendship(ctx, friendService)
		})

		userGroup.GET("/:username/friends", func(ctx *gin.Context) {
			GetFriendList(ctx, friendService)
		})

		userGroup.GET("/:username/friends/count", func(ctx *gin.Context) {
			CountFriends(ctx, friendService)
		})

		userGroup.DELETE("/:username/friends", func(ctx *gin.Context) {
			DeleteFriend(ctx, friendService)
		})

		// friend request

		userGroup.GET("/:username/friends/requests/:friend/exists", func(ctx *gin.Context) {
			CheckFriendRequest(ctx, friendService)
		})

		userGroup.GET("/:username/friends/requests", func(ctx *gin.Context) {
			GetFriendRequestList(ctx, friendService)
		})

		userGroup.GET("/:username/friends/requests/count", func(ctx *gin.Context) {
			CountFriendRequests(ctx, friendService)
		})

		userGroup.POST("/:username/friends/requests", func(ctx *gin.Context) {
			SendFriendRequest(ctx, friendService)
		})

		userGroup.DELETE("/:username/friends/requests", func(ctx *gin.Context) {
			DeleteFriendRequest(ctx, friendService)
		})
	}
}
