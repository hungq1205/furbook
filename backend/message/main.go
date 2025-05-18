package main

import (
	"message/api/client"
	"message/api/handler/group"
	"message/api/handler/message"
	repository "message/infrastructure/repository"
	groupService "message/usecase/group"
	messageService "message/usecase/message"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	dsn := "host=messagedb user=postgres password=root dbname=message port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	messageRepo := repository.NewMessageRepository(db)
	groupRepo := repository.NewGroupRepository(db)
	groupUserRepo := repository.NewGroupUserRepository(db)

	messageService := messageService.NewService(messageRepo, groupUserRepo)
	groupService := groupService.NewService(groupRepo, groupUserRepo)

	userClient := client.NewUserClient("http://user:8080")

	message.MakeHandler(app, messageService, groupService)
	group.MakeHandler(app, groupService, messageService, userClient)

	return app
}
