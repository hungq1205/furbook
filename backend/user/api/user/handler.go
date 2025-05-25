package user

import (
	"user/api/client"
	"user/api/middleware"
	"user/usecase/friend"
	"user/usecase/user"

	"github.com/gin-gonic/gin"
)

func MakeHandler(app *gin.Engine, userService user.UseCase, friendService friend.UseCase, groupClient client.GroupClient) {
	userGroup := app.Group("/api/user")
	{
		// Public routes
		userGroup.GET("/:username", func(c *gin.Context) {
			GetUser(c, userService, friendService)
		})

		userGroup.POST("/list", func(c *gin.Context) {
			GetUserList(c, userService, friendService)
		})

		userGroup.POST("", func(c *gin.Context) {
			CreateUser(c, userService, friendService)
		})

		// Authenticated routes
		authGroup := userGroup.Group("", middleware.MustAuthMiddleware())

		authGroup.PATCH("", func(c *gin.Context) {
			UpdateUser(c, userService, friendService)
		})

		authGroup.DELETE("", func(c *gin.Context) {
			DeleteUser(c, userService)
		})

		authGroup.GET("/friends", func(c *gin.Context) {
			GetFriendList(c, friendService, groupClient)
		})

		authGroup.GET("/check-friendship/:username", func(c *gin.Context) {
			CheckFriendship(c, friendService)
		})

		authGroup.DELETE("/friends", func(c *gin.Context) {
			DeleteFriend(c, friendService)
		})

		authGroup.GET("/friend-requests", func(c *gin.Context) {
			GetFriendRequestList(c, friendService)
		})

		authGroup.POST("/friend-requests", func(c *gin.Context) {
			SendFriendRequest(c, friendService)
		})

		authGroup.DELETE("/friend-requests/revoke", func(c *gin.Context) {
			DeleteFriendRequest(c, friendService)
		})

		authGroup.DELETE("/friend-requests/decline", func(c *gin.Context) {
			DeleteIncomingFriendRequest(c, friendService)
		})
	}
}
