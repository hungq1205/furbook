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
		groupGroup.GET("/:groupId", func(ctx *gin.Context) {
			getGroup(ctx, groupService, messageService)
		})

		groupGroup.GET("/:groupId/members", func(ctx *gin.Context) {
			getMembersOfGroup(ctx, groupService, userClient)
		})

		authGroup := groupGroup.Group("", middleware.MustAuthMiddleware())

		authGroup.GET("", func(ctx *gin.Context) {
			getGroupsOfUser(ctx, groupService, messageService)
		})

		authGroup.POST("", func(ctx *gin.Context) {
			createGroup(ctx, groupService, messageService)
		})

		authGroup.DELETE("", func(ctx *gin.Context) {
			deleteGroup(ctx, groupService)
		})

		authGroup.PUT("", func(ctx *gin.Context) {
			updateGroup(ctx, groupService, messageService)
		})

		authGroup.POST("/:groupId/members", func(ctx *gin.Context) {
			addMemberToGroup(ctx, groupService, messageService)
		})

		authGroup.DELETE("/:groupId/members", func(ctx *gin.Context) {
			removeMemberToGroup(ctx, groupService, messageService)
		})
	}
}
