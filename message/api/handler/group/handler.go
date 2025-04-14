package group

import (
	"github.com/gin-gonic/gin"
	"test/api/client"
	"test/usecase/group"
)

func MakeHandler(app *gin.Engine, groupService group.UseCase, userClient client.UserClient) {
	groupGroup := app.Group("/api/group")
	{
		groupGroup.GET("/:groupId", func(ctx *gin.Context) {
			getGroup(ctx, groupService)
		})

		groupGroup.GET("/:groupId/members", func(ctx *gin.Context) {
			getMembersOfGroup(ctx, groupService, userClient)
		})

		groupGroup.GET("", func(ctx *gin.Context) {
			getGroupsOfUser(ctx, groupService)
		})

		groupGroup.POST("", func(ctx *gin.Context) {
			createGroup(ctx, groupService)
		})

		groupGroup.DELETE("", func(ctx *gin.Context) {
			deleteGroup(ctx, groupService)
		})

		groupGroup.PUT("", func(ctx *gin.Context) {
			updateGroup(ctx, groupService)
		})

		groupGroup.POST("/:groupId/members", func(ctx *gin.Context) {
			addMemberToGroup(ctx, groupService)
		})

		groupGroup.DELETE("/:groupId/members", func(ctx *gin.Context) {
			removeMemberToGroup(ctx, groupService)
		})
	}
}
