package post

import (
	"post/api/client"
	"post/api/middleware"
	"post/usecase/post"

	"github.com/gin-gonic/gin"
)

func MakeHandler(app *gin.Engine, postService *post.Service, userClient client.UserClient) {
	postGroup := app.Group("/api/post")
	{
		postGroup.GET("/:postID", func(c *gin.Context) {
			GetPost(c, postService, userClient)
		})

		postGroup.GET("/lost", func(c *gin.Context) {
			GetNearLostPosts(c, postService, userClient)
		})

		postGroup.GET("/ofUser/:username/participated", func(c *gin.Context) {
			GetParticipatedPostsOfUser(c, postService, userClient)
		})

		postGroup.GET("/ofUser/:username", func(c *gin.Context) {
			GetPostsOfUser(c, postService, userClient)
		})

		postGroup.POST("/ofUsers", func(c *gin.Context) {
			GetPostsOfUsers(c, postService, userClient)
		})

		postGroup.GET("/:postID/comments", func(c *gin.Context) {
			GetComments(c, postService, userClient)
		})

		authGroup := postGroup.Group("", middleware.MustAuthorizeMiddleware())

		authGroup.POST("/blog", func(c *gin.Context) {
			CreateBlogPost(c, postService)
		})

		authGroup.POST("/lost", func(c *gin.Context) {
			CreateLostPetPost(c, postService, userClient)
		})

		authGroup.PATCH("/:postID/content", func(c *gin.Context) {
			PatchContentPost(c, postService)
		})

		authGroup.PATCH("/:postID/lostFoundStatus", func(c *gin.Context) {
			PatchLostFoundStatus(c, postService)
		})

		authGroup.DELETE("", func(c *gin.Context) {
			DeletePost(c, postService)
		})

		authGroup.POST("/:postID/comments", func(c *gin.Context) {
			CreateComment(c, postService)
		})

		authGroup.POST("/:postID/interactions", func(c *gin.Context) {
			UpsertInteraction(c, postService)
		})

		authGroup.DELETE("/:postID/interactions", func(c *gin.Context) {
			DeleteInteraction(c, postService)
		})

		authGroup.POST("/:postID/participation", func(c *gin.Context) {
			Participate(c, postService)
		})

		authGroup.DELETE("/:postID/participation", func(c *gin.Context) {
			Unparticipate(c, postService)
		})
	}
}
