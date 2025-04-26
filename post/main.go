package main

import (
	"context"
	"post/api/client"
	"post/api/handler/post"
	postRepo "post/infrastructure/repository/post"
	postService "post/usecase/post"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://postdb:27017"))
	if err != nil {
		panic(err)
	}
	mongoDB := mongoClient.Database("post")

	postRepo := postRepo.NewRepository(mongoDB)
	postService := postService.NewService(postRepo)

	userClient := client.NewUserClient("http://user:8080")

	post.MakeHandler(app, postService, userClient)

	return app
}
