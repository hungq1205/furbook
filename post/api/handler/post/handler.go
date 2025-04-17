package post

import (
	"post/api/middleware"
	"post/usecase/post"

	"github.com/gin-gonic/gin"
)

func MakeHandler(app *gin.Engine, postService *post.Service) {
	postGroup := app.Group("/api/post")
	{
		postGroup.GET("/:postID", func(c *gin.Context) {
			GetPost(c, postService)
		})

		postGroup.GET("/ofuser/:username", func(c *gin.Context) {
			GetPostsOfUser(c, postService)
		})

		postGroup.GET("/ofUsers", func(c *gin.Context) {
			GetPostsOfUsers(c, postService)
		})

		postGroup.GET("/:postID/comments", func(c *gin.Context) {
			GetComments(c, postService)
		})

		authGroup := postGroup.Group("", middleware.MustAuthorizeMiddleware())

		authGroup.POST("/blog", func(c *gin.Context) {
			CreateBlogPost(c, postService)
		})

		authGroup.POST("/lost", func(c *gin.Context) {
			CreateBlogPost(c, postService)
		})

		authGroup.PATCH("/:postID/content", func(c *gin.Context) {
			PatchContentPost(c, postService)
		})

		authGroup.PATCH("/:postID/found", func(c *gin.Context) {
			PatchFoundPost(c, postService)
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
	}
}
