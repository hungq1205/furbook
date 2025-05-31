package main

import (
	"fmt"
	"noti/api/client"
	"noti/api/noti"
	notiRepo "noti/infrastructure/repository/noti"
	notiService "noti/usecase/noti"

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

	dsn := "host=notidb user=postgres password=root dbname=noti port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(fmt.Sprintf("failed to connect database %v", err.Error()))
	}

	wsClient := client.NewWsClient("http://gateway:8080")

	notiRepo := notiRepo.NewRepository(db)
	notiService := notiService.NewService(notiRepo)
	noti.MakeHandler(app, notiService, wsClient)

	return app
}
