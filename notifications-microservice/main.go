package main

import (
	"notifications-microservice/src/config"
	"notifications-microservice/src/controllers"
	"notifications-microservice/src/db"
	"notifications-microservice/src/repositories"
	"notifications-microservice/src/router"
	"notifications-microservice/src/services"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	loadSecretsFunc := config.LoadSecrets
	connectDBFunc := db.ConnectDB
	e, port, errSecrets, errDB := initializeApp(loadSecretsFunc, connectDBFunc)

	if errSecrets != nil {
		panic(errSecrets)
	}
	if errDB != nil {
		panic(errDB)
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func initializeApp(
	loadSecretsFunc func() (string, string, error),
	connectDB func(string) (*mongo.Client, error),
) (*echo.Echo, string, error, error) {

	port, mongoURI, errSecrets := loadSecretsFunc()
	if errSecrets != nil {
		return nil, "", errSecrets, nil
	}

	client, errDB := connectDB(mongoURI)
	if errDB != nil {
		return nil, "", nil, errDB
	}

	db := client.Database("notificationmcsv")

	notificationRepo := repositories.NewNotificationRepository(db)
	notificationService := services.NewNotificationService(notificationRepo)
	notificationController := controllers.NewNotificationController(notificationService)

	e := echo.New()

	router := router.NewRouter(notificationController)

	router.NotificationRoutes(e)

	return e, port, nil, nil
}
