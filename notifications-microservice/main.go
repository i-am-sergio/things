package main

import (
	"notifications-microservice/src/config"
	"notifications-microservice/src/db"
	"notifications-microservice/src/repositories"
	"notifications-microservice/src/router"
	"notifications-microservice/src/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e, port := initializeApp()
	e.Logger.Fatal(e.Start(":" + port))
}

func initializeApp() (*echo.Echo, string) {
	port, mongoURI, err := config.LoadSecrets()
	if err != nil {
		panic(err)
	}

	client := db.ConnectDB(mongoURI)
	db := client.Database("notificationmcsv")

	notificationRepo := repositories.NewNotificationRepository(db)
	notificationService := services.NewNotificationService(notificationRepo)

	e := echo.New()
	router.NotificationRoutes(e, notificationService)

	return e, port
}
