package group

import (
	"message/api/client"
	"message/api/middleware"
	"message/usecase/group"
	"message/usecase/message"

	"github.com/gin-gonic/gin"
)

func MakeHandler(app *gin.Engine, groupService group.UseCase, messageService message.UseCase, userClient client.UserClient) {
	groupGroup := app.Group("/api/group")
	{
		authGroup := groupGroup.Group("", middleware.MustAuthMiddleware())

		authGroup.GET("/:groupId", func(ctx *gin.Context) {
			getGroup(ctx, groupService, messageService, userClient)
		})

		authGroup.GET("/:groupId/members", func(ctx *gin.Context) {
			getMembersOfGroup(ctx, groupService, userClient)
		})

		authGroup.GET("/direct/:username", func(ctx *gin.Context) {
			getDirectGroup(ctx, groupService, messageService, userClient)
		})

		authGroup.GET("", func(ctx *gin.Context) {
			getGroupsOfUser(ctx, groupService, messageService, userClient)
		})

		authGroup.POST("", func(ctx *gin.Context) {
			createGroup(ctx, groupService, messageService, userClient)
		})

		authGroup.DELETE("", func(ctx *gin.Context) {
			deleteGroup(ctx, groupService)
		})

		authGroup.PUT("", func(ctx *gin.Context) {
			updateGroup(ctx, groupService, messageService, userClient)
		})

		authGroup.POST("/:groupId/members", func(ctx *gin.Context) {
			addMemberToGroup(ctx, groupService, messageService, userClient)
		})

		authGroup.DELETE("/:groupId/members", func(ctx *gin.Context) {
			removeMemberToGroup(ctx, groupService, messageService, userClient)
		})
	}
}
