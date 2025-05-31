package noti

import (
	"noti/api/client"
	"noti/api/middleware"
	noti "noti/usecase/noti"

	"github.com/gin-gonic/gin"
)

func MakeHandler(app *gin.Engine, notiService noti.UseCase, wsClient client.WsClient) {
	notiGroup := app.Group("/api/noti")
	{
		authGroup := notiGroup.Group("", middleware.MustAuthMiddleware())

		authGroup.GET("/:id", func(c *gin.Context) {
			GetNoti(c, notiService)
		})

		authGroup.GET("", func(c *gin.Context) {
			GetNotisOfUser(c, notiService)
		})

		authGroup.POST("", func(c *gin.Context) {
			CreateNoti(c, notiService, wsClient)
		})

		authGroup.POST("/createMultiple", func(c *gin.Context) {
			CreateNotiToUsers(c, notiService, wsClient)
		})

		authGroup.PATCH("/:id", func(c *gin.Context) {
			UpdateNoti(c, notiService)
		})

		authGroup.DELETE("/:id", func(c *gin.Context) {
			DeleteNoti(c, notiService)
		})
	}
}
