package message

import (
	"message/api/client"
	"message/api/middleware"
	"message/usecase/group"
	"message/usecase/message"

	"github.com/gin-gonic/gin"
)

func MakeHandler(app *gin.Engine, messageService message.UseCase, groupService group.UseCase, wsClient client.WsClient) {
	messageGroup := app.Group("/api/message")
	{
		messageGroup.Use(middleware.MustAuthMiddleware())

		messageGroup.GET("/group/:groupID", func(ctx *gin.Context) {
			getGroupMessageList(ctx, messageService, groupService)
		})

		messageGroup.GET("/direct", func(ctx *gin.Context) {
			getDirectMessageList(ctx, messageService)
		})

		messageGroup.POST("/group/:groupID", func(ctx *gin.Context) {
			createGroupMessage(ctx, messageService, groupService, wsClient)
		})
	}
}
