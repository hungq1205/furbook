package main

import (
	"user/api/user"
	friendRepo "user/infrastructure/repository/friend"
	userRepo "user/infrastructure/repository/user"
	friendService "user/usecase/friend"
	userService "user/usecase/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
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
	db, err := postgres.Open(dsn)
	if err != nil {
		panic(err)
	}

	userRepo := userRepo.NewRepository(db)
	friendRepo := friendRepo.NewRepository(db)

	friendService := friendService.NewService(friendRepo)
	userService := userService.NewService(userRepo)

	user.MakeHandler(app, userService, friendService)

	return app
}
