package main

import (
	"fmt"
	"user/api/client"
	"user/api/user"
	friendRepo "user/infrastructure/repository/friend"
	userRepo "user/infrastructure/repository/user"
	friendService "user/usecase/friend"
	userService "user/usecase/user"

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
	app.Use(gin.Recovery())
	app.Use(gin.Logger())

	dsn := "host=userdb user=postgres password=root dbname=user port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(fmt.Sprintf("failed to connect database %v", err.Error()))
	}

	userRepo := userRepo.NewRepository(db)
	friendRepo := friendRepo.NewRepository(db)

	friendService := friendService.NewService(friendRepo)
	userService := userService.NewService(userRepo)
	groupClient := client.NewGroupClient("http://message:8080")

	user.MakeHandler(app, userService, friendService, groupClient)

	return app
}
