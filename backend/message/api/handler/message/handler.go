package message

import (
	"github.com/gin-gonic/gin"
	"message/api/middleware"
	"message/usecase/group"
	"message/usecase/message"
)

func MakeHandler(app *gin.Engine, messageService message.UseCase, groupService group.UseCase) {
	messageGroup := app.Group("/api/message")
	{
		messageGroup.Use(middleware.AuthorizeMiddleware())

		messageGroup.GET("/group/:groupID", func(ctx *gin.Context) {
			getGroupMessageList(ctx, messageService, groupService)
		})

		messageGroup.GET("/direct", func(ctx *gin.Context) {
			getDirectMessageList(ctx, messageService)
		})

		messageGroup.POST("/group/:groupID", func(ctx *gin.Context) {
			createGroupMessage(ctx, messageService, groupService)
		})

		messageGroup.POST("/direct", func(ctx *gin.Context) {
			createDirectMessage(ctx, messageService, groupService)
		})
	}
}
