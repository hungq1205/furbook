package main

import (
	"post/api/handler/post"
	postRepo "post/infrastructure/repository/post"
	postService "post/usecase/post"

	"github.com/gin-gonic/gin"
)

func main() {
	app := makeHandler()
	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}

func makeHandler() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	postRepo := postRepo.NewRepository()
	postService := postService.NewPostService(postRepo)

	post.MakeHandler(app, postService)

	return app
}
